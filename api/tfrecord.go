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
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// TwoFactorAPI exposes the full wrapper over the http api with underline db
// for the database linked in.
type TwoFactorAPI struct {
	tfrecordapi.TFRecordHTTP
	Domain              string
	CodeLength          int
	Backend             backends.TFBackend
	UserBackend         types.UserDBBackend
	IsNotFoundErrorFunc func(error) bool
}

// NewTwoFactorAPI returns a new instance of TwoFactorAPI.
func NewTwoFactorAPI(m metrics.Metrics, codeln int, domain string, tdb types.TFRecordDBBackend, users types.UserDBBackend) TwoFactorAPI {
	backend := backends.TFBackend{TFRecordDBBackend: tdb}
	return TwoFactorAPI{
		Domain:              domain,
		CodeLength:          codeln,
		Backend:             backend,
		UserBackend:         users,
		IsNotFoundErrorFunc: db.IsNotFoundError,
		TFRecordHTTP:        tfrecordapi.New(m, backend),
	}
}

// isNotFoundError returns true/false if the giving error matches
// an expected ErrNotFound error. It helps detach us
// from code smell and import pollution.
func (t TwoFactorAPI) isNotFoundError(err error) bool {
	if t.IsNotFoundErrorFunc == nil {
		return false
	}

	return t.IsNotFoundErrorFunc(err)
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
func (t TwoFactorAPI) RetrieveUser(ctx *httputil.Context) error {
	id, ok := ctx.GetString("public_id")
	if !ok {
		err := errors.New("expected public_id param")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	// Retrieve user from db.
	user, err := t.UserBackend.Get(ctx.Context(), id)
	if err != nil {
		if !t.isNotFoundError(err) {
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
	if tf, err := t.Backend.GetByField(ctx.Context(), "user_id", user.PublicID); err == nil {
		ctx.Set(pkg.ContextKeyTFRecord, tf)
	}

	// set current user into context.
	ctx.Set(pkg.ContextKeyUser, user)
	return nil
}

// TwoFactorQRImage requests a user's twofactor png image which contains the necessary
// code needed to activate user code generation on user's device.
//
// Method: GET
// Params: :public_id
//
func (t TwoFactorAPI) TwoFactorQRImage(ctx *httputil.Context) error {
	if err := t.RetrieveUser(ctx); err != nil {
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
func (t TwoFactorAPI) DisableTwoFactor(ctx *httputil.Context) error {
	if err := t.RetrieveUser(ctx); err != nil {
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
	if err := t.Backend.Delete(ctx.Context(), twofactor.PublicID); err != nil {
		if !t.isNotFoundError(err) {
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
	if err := t.UserBackend.Update(ctx.Context(), user.PublicID, user); err != nil {
		if !t.isNotFoundError(err) {
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

	// set current user into context.
	ctx.Set(pkg.ContextKeyUser, user)
	ctx.Status(http.StatusNoContent)
	return nil
}

// EnableTwoFactor requests the enabling of twofactor authentication for a giving
// user, it expects to have the user's public id as a parameter.
//
// Method: GET
// Params: :public_id
//
func (t TwoFactorAPI) EnableTwoFactor(ctx *httputil.Context) error {
	if err := t.RetrieveUser(ctx); err != nil {
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
	newTF.Domain = t.Domain
	newTF.MaxLength = t.CodeLength

	ntwRecord, err := t.Backend.Create(ctx.Context(), newTF)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if err := t.UserBackend.Update(ctx.Context(), user.PublicID, user); err != nil {
		if !t.isNotFoundError(err) {
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

	ctx.Status(http.StatusNoContent)
	ctx.Set(pkg.ContextKeyUser, user)
	ctx.Set(pkg.ContextKeyTFRecord, ntwRecord)
	return nil
}
