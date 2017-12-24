package api

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/tfrecordapi"
	"github.com/gokit/tenancykit/pkg/resources/userapi"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// UserAPI implements the http api for the user request-response api.
type UserAPI struct {
	userapi.UserHTTP
	Domain              string
	TwoFactorCodeLength int
	Backend             types.UserDBBackend
	TFBackend           tfrecordapi.TFRecordBackend
	TFDBBackend         types.TFRecordDBBackend
	IsNotFoundErrorFunc func(error) bool
}

// NewUserAPI returns a new instance of UserAPI.
func NewUserAPI(m metrics.Metrics, multitenant bool, domain string, tfCodeLen int, users types.UserDBBackend, tf types.TFRecordDBBackend) UserAPI {
	var api UserAPI
	api.Domain = domain
	api.Backend = users
	api.TFDBBackend = tf
	api.TwoFactorCodeLength = tfCodeLen
	api.IsNotFoundErrorFunc = db.IsNotFoundError
	api.TFBackend = backends.TFBackend{TFRecordDBBackend: tf}
	api.UserHTTP = userapi.New(m, backends.UserBackend{UserDBBackend: users, MultiTenant: multitenant})
	return api
}

// isNotFoundError returns true/false if the giving error matches
// an expected ErrNotFound error. It helps detach us
// from code smell and import pollution.
func (u UserAPI) isNotFoundError(err error) bool {
	if u.IsNotFoundErrorFunc == nil {
		return false
	}

	return u.IsNotFoundErrorFunc(err)
}

// RetrieveUser attempts to use the provided public id parameter to load
// the expected user account into the current request context.
//
// Method: GET
// Params: :public_id
//
// Failure:
//	if User does not exists: 404
//
func (u UserAPI) RetrieveUser(ctx *httputil.Context) error {
	id, ok := ctx.Bag().GetString("public_id")
	if !ok {
		err := errors.New("expected public_id param")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	// Retrieve user from db.
	user, err := u.Backend.Get(ctx.Context(), id)
	if err != nil {
		if !u.isNotFoundError(err) {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusNotFound,
		}
	}

	// Attempt to retrieve twofactor user record if user has one.
	if tf, err := u.TFDBBackend.GetByField(ctx.Context(), "user_id", user.PublicID); err == nil {
		ctx.Bag().Set(pkg.ContextKeyTFRecord, tf)
	}

	// set current user into context.
	ctx.Bag().Set(pkg.ContextKeyUser, user)
	return nil
}

// TwoFactorQRImage requests a user's twofactor png image which contains the necessary
// code needed to activate user code generation on user's device.
//
// Method: GET
// Params: :public_id
//
func (u UserAPI) TwoFactorQRImage(ctx *httputil.Context) error {
	if err := u.RetrieveUser(ctx); err != nil {
		return err
	}

	user, err := pkg.GetUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if !user.TwoFactorAuth {
		return httputil.HTTPError{
			Err:  errors.New("user has not enabled twofactor auth"),
			Code: http.StatusBadRequest,
		}
	}

	twofactor, err := pkg.GetTwoFactorRecord(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	qr, err := twofactor.QR()
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	return ctx.Stream(http.StatusOK, "image/png", bytes.NewBuffer(qr))
}

// DisableTwoFactor requests the disabling of twofactor authentication for a giving
// user, it expects to have the user's public id as a parameter.
//
// Method: GET
// Params: :public_id
//
func (u UserAPI) DisableTwoFactor(ctx *httputil.Context) error {
	if err := u.RetrieveUser(ctx); err != nil {
		return err
	}

	user, err := pkg.GetUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if !user.TwoFactorAuth {
		return nil
	}

	twofactor, err := pkg.GetTwoFactorRecord(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	// Delete user's twofactor record from db.
	if err := u.TFBackend.Delete(ctx.Context(), twofactor.PublicID); err != nil {
		if !u.isNotFoundError(err) {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusNotFound,
		}
	}

	user.TwoFactorAuth = false
	if err := u.Backend.Update(ctx.Context(), user.PublicID, user); err != nil {
		if !u.isNotFoundError(err) {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusNotFound,
		}
	}

	return nil
}

// EnableTwoFactor requests the enabling of twofactor authentication for a giving
// user, it expects to have the user's public id as a parameter.
//
// Method: GET
// Params: :public_id
//
func (u UserAPI) EnableTwoFactor(ctx *httputil.Context) error {
	if err := u.RetrieveUser(ctx); err != nil {
		return err
	}

	user, err := pkg.GetUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if user.TwoFactorAuth {
		return nil
	}

	user.TwoFactorAuth = true

	// Create new twofactor record for user.
	var newTF pkg.NewTF
	newTF.User = user
	newTF.Domain = u.Domain
	newTF.MaxLength = u.TwoFactorCodeLength

	ntwRecord, err := u.TFBackend.Create(ctx.Context(), newTF)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	ctx.Bag().Set(pkg.ContextKeyTFRecord, ntwRecord)
	return nil
}
