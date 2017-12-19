package baseapi

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/api/userapi"
	"github.com/gokit/tenancykit/db/backends"
	"github.com/gokit/tenancykit/db/types"
	"github.com/influx6/faux/httputil"
)

// UserAPI implements the http api for the user request-response api.
type UserAPI struct {
	userapi.UserHTTP
	TwoFactorCodeLength int
	Backend             types.UserDBBackend
	TenantBackend       types.TenantDBBackend
	TFBackend           backends.TFBackend
	IsNotFoundErrorFunc func(error) bool
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
		err := errors.New("Expected public_id param")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	var err error
	var currentUser tenancykit.CurrentUser

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
	if tf, err := u.TFBackend.GetByField(ctx, "user_id", currentUser.User.PublicID); err == nil {
		currentUser.TwoFactor = &tf
	}

	// set current user into context.
	ctx.Bag().Set(tenancykit.ContextKeyCurrentUser, currentUser)
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
	var currentUser tenancykit.CurrentUser

	if currentUser, err = tenancykit.GetCurrentUser(ctx); err != nil {
		if rerr := u.RetrieveUser(ctx); rerr != nil {
			return rerr
		}

		// Attempt to retreive current user once again.
		currentUser, err = tenancykit.GetCurrentUser(ctx)
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
	var currentUser tenancykit.CurrentUser

	if currentUser, err = tenancykit.GetCurrentUser(ctx); err != nil {
		if rerr := u.RetrieveUser(ctx); rerr != nil {
			return rerr
		}

		// Attempt to retreive current user once again.
		currentUser, err = tenancykit.GetCurrentUser(ctx)
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
	var currentUser tenancykit.CurrentUser

	if currentUser, err = tenancykit.GetCurrentUser(ctx); err != nil {
		if rerr := u.RetrieveUser(ctx); rerr != nil {
			return rerr
		}
		currentUser, err = tenancykit.GetCurrentUser(ctx)
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
	var newTF tenancykit.NewTF
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
	ctx.Bag().Set(tenancykit.ContextKeyCurrentUser, currentUser)

	return nil
}
