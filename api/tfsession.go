package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gokit/tenancykit/pkg/db"
	"github.com/gokit/tokens"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/twofactorsessionapi"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// TwoFactorSessionAPI augments the existing twofactorsessionapi.TwoFactorSessionHTTP with other
// methods that expose more functionality.
type TwoFactorSessionAPI struct {
	twofactorsessionapi.TwoFactorSessionHTTP
	Backend             types.TwoFactorSessionDBBackend
	TokenSet            tokens.TokenSet
	IsNotFoundErrorFunc func(error) bool
}

// NewTwoFactorSessionAPI returns a new instance of TwoFactorSessionAPI.
func NewTwoFactorSessionAPI(m metrics.Metrics, tokenset tokens.TokenSet, tfsession types.TwoFactorSessionDBBackend) TwoFactorSessionAPI {
	var api TwoFactorSessionAPI
	api.Backend = tfsession
	api.TokenSet = tokenset
	api.IsNotFoundErrorFunc = db.IsNotFoundError
	api.TwoFactorSessionHTTP = twofactorsessionapi.New(m, TwoFactorSessionBackend{
		TwoFactorSessionDBBackend: tfsession,
	})

	return api
}

// ValidateUserToken processing incoming one-time request to validate users
// code, this is useful for times of requiring a one-time authentication of
// user request.
//
// Params:
//		:user_code => "434544" // provide user generate auth code
//
func (tfs TwoFactorSessionAPI) ValidateUserToken(ctx *httputil.Context) error {
	currentUser, err := pkg.GetCurrentUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	// if user does not require twofactor auth, then skip this.
	if !currentUser.User.TwoFactorAuth {
		ctx.Status(http.StatusNoContent)
		return nil
	}

	// If we failed to load users two factor record then return error.
	if currentUser.TwoFactor == nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	userCode, ok := ctx.Bag().GetString("user_code")
	if !ok {
		return httputil.HTTPError{
			Err:  errors.New("user_code param required"),
			Code: http.StatusBadRequest,
		}
	}

	used, err := tfs.TokenSet.Has(ctx.Context(), currentUser.User.PublicID, userCode)
	if err != nil && !tfs.isNotFoundError(err) {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if used {
		return httputil.HTTPError{
			Err:  errors.New("twofactor user code already used"),
			Code: http.StatusBadRequest,
		}
	}

	if err := currentUser.TwoFactor.ValidateOTP(userCode); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	if _, err := tfs.TokenSet.Add(ctx.Context(), currentUser.User.PublicID, userCode); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	ctx.Status(http.StatusNoContent)

	return nil
}

// NewSession processing incoming request for validating user's login request
// using provided user code from two-factor app or otp.
// NewSession expects the user has already logged in through normal sessions
// and the appropriate `CurrentUser` data has being set in the incoming request
// context.
//
// Params:
//		:user_code => "434544" // provide user generate auth code
//
func (tfs TwoFactorSessionAPI) NewSession(ctx *httputil.Context) error {
	currentUser, err := pkg.GetCurrentUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	// If user already has twofactor session locked in, then skip.
	if currentUser.TFSession != nil {
		ctx.Status(http.StatusNoContent)
		return nil
	}

	// if user does not require twofactor auth, then skip this.
	if !currentUser.User.TwoFactorAuth {
		ctx.Status(http.StatusNoContent)
		return nil
	}

	// If we failed to load users two factor record then return error.
	if currentUser.TwoFactor == nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	userCode, ok := ctx.Bag().GetString("user_code")
	if !ok {
		return httputil.HTTPError{
			Err:  errors.New("user_code param required"),
			Code: http.StatusBadRequest,
		}
	}

	used, err := tfs.TokenSet.Has(ctx.Context(), currentUser.User.PublicID, userCode)
	if err != nil && !tfs.isNotFoundError(err) {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if used {
		return httputil.HTTPError{
			Err:  errors.New("twofactor user code already used"),
			Code: http.StatusBadRequest,
		}
	}

	// If we have an existing two-factor session, then validate code sent is correct,
	// then update current tfrecord and pass on the session record to context.
	if tfSession, err := tfs.Backend.GetByField(ctx.Context(), "user_id", currentUser.User.PublicID); err == nil {
		if err := currentUser.TwoFactor.ValidateOTP(userCode); err != nil {
			tfs.Backend.Delete(ctx.Context(), tfSession.PublicID)

			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusBadRequest,
			}
		}

		currentUser.TFSession = &tfSession
		ctx.Bag().Set(pkg.ContextKeyCurrentUser, currentUser)

		ctx.Status(http.StatusNoContent)
		return nil
	}

	if err := currentUser.TwoFactor.ValidateOTP(userCode); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	if _, err := tfs.TokenSet.Add(ctx.Context(), currentUser.User.PublicID, userCode); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	newTFSession := pkg.NewTwoFactorSession(currentUser.User.PublicID, true)

	if err := tfs.Backend.Create(ctx.Context(), newTFSession); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	currentUser.TFSession = &newTFSession
	ctx.Bag().Set(pkg.ContextKeyCurrentUser, currentUser)
	ctx.Status(http.StatusNoContent)

	return nil
}

// isNotFoundError returns true/false if the giving error matches
// an expected ErrNotFound error. It helps detach us
// from code smell and import pollution.
func (tfs TwoFactorSessionAPI) isNotFoundError(err error) bool {
	if tfs.IsNotFoundErrorFunc == nil {
		return false
	}

	return tfs.IsNotFoundErrorFunc(err)
}

// TwoFactorSessionBackend implements the api.twofactorsessionapi.Backend interface and wraps a types.TFRecordBackend.
type TwoFactorSessionBackend struct {
	types.TwoFactorSessionDBBackend
}

// Create attempts to create a new TFRecord in the db using the provided pkg.NewTF struct.
func (tf TwoFactorSessionBackend) Create(ctx context.Context, ntf pkg.TwoFactorSession) (pkg.TwoFactorSession, error) {
	return ntf, tf.TwoFactorSessionDBBackend.Create(ctx, ntf)
}
