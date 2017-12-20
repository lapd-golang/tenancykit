package tenancykit

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
	TwoFactorCodeLength int
	Backend             types.UserDBBackend
	TenantBackend       types.TenantDBBackend
	TFBackend           tfrecordapi.TFRecordBackend
	TFDBBackend         types.TFRecordDBBackend
	IsNotFoundErrorFunc func(error) bool
}

// NewUserAPI returns a new instance of UserAPI.
func NewUserAPI(tfCodeLen int, m metrics.Metrics, users types.UserDBBackend, tenants types.TenantDBBackend, tf types.TFRecordDBBackend) UserAPI {
	var api UserAPI
	api.Backend = users
	api.TFDBBackend = tf
	api.TenantBackend = tenants
	api.TwoFactorCodeLength = tfCodeLen
	api.IsNotFoundErrorFunc = db.IsNotFoundError
	api.TFBackend = backends.TFBackend{TFRecordDBBackend: tf}
	api.UserHTTP = userapi.New(m, backends.UserBackend{UserDBBackend: users})
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

	var err error
	var currentUser pkg.CurrentUser

	// Retreive user from db.
	currentUser.User, err = u.Backend.Get(ctx, id)
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

	// Rerieve user's tenant.
	currentUser.Tenant, err = u.TenantBackend.Get(ctx, currentUser.User.TenantID)
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
	if tf, err := u.TFDBBackend.GetByField(ctx, "user_id", currentUser.User.PublicID); err == nil {
		currentUser.TwoFactor = &tf
	}

	// set current user into context.
	ctx.Bag().Set(pkg.ContextKeyCurrentUser, currentUser)
	return nil
}

// TwoFactorQRImage requests a user's twofactor png image which contains the necessary
// code needed to activate user code generation on user's device.
//
// Method: GET
// Params: :public_id
//
func (u UserAPI) TwoFactorQRImage(ctx *httputil.Context) error {
	var err error
	var currentUser pkg.CurrentUser

	if currentUser, err = pkg.GetCurrentUser(ctx); err != nil {
		if rerr := u.RetrieveUser(ctx); rerr != nil {
			return rerr
		}

		// Attempt to retreive current user once again.
		currentUser, err = pkg.GetCurrentUser(ctx)
		if err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}
	}

	if !currentUser.User.TwoFactorAuth {
		err = errors.New("User must first enable twofactor authentication")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	if currentUser.User.TwoFactorAuth && currentUser.TwoFactor == nil {
		err = errors.New("User TwoFactorRecord failed to be loaded")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	qr, err := currentUser.TwoFactor.QR()
	if err != nil {
		return err
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
	var err error
	var currentUser pkg.CurrentUser

	if currentUser, err = pkg.GetCurrentUser(ctx); err != nil {
		if rerr := u.RetrieveUser(ctx); rerr != nil {
			return rerr
		}

		// Attempt to retreive current user once again.
		currentUser, err = pkg.GetCurrentUser(ctx)
		if err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}
	}

	if !currentUser.User.TwoFactorAuth && currentUser.TwoFactor == nil {
		return nil
	}

	// Delete user's twofactor record from db.
	if err := u.TFBackend.Delete(ctx, currentUser.TwoFactor.PublicID); err != nil {
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

	currentUser.User.TwoFactorAuth = false
	if err := u.Backend.Update(ctx, currentUser.User.PublicID, currentUser.User); err != nil {
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
	var err error
	var currentUser pkg.CurrentUser

	if currentUser, err = pkg.GetCurrentUser(ctx); err != nil {
		if rerr := u.RetrieveUser(ctx); rerr != nil {
			return rerr
		}
		currentUser, err = pkg.GetCurrentUser(ctx)
		if err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}
	}

	domainURL := ctx.Request().URL
	domain := domainURL.Host

	if currentUser.User.TwoFactorAuth && currentUser.TwoFactor != nil {
		return nil
	}

	currentUser.User.TwoFactorAuth = true

	// Create new twofactor record for user.
	var newTF pkg.NewTF
	newTF.Domain = domain
	newTF.User = currentUser.User
	newTF.Tenant = currentUser.Tenant
	newTF.MaxLength = u.TwoFactorCodeLength

	ntwRecord, err := u.TFBackend.Create(ctx, newTF)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	currentUser.TwoFactor = &ntwRecord
	ctx.Bag().Set(pkg.ContextKeyCurrentUser, currentUser)

	return nil
}
