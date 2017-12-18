package baseapi

import (
	"errors"
	"net/http"

	"github.com/gokit/tenancykit"
	"github.com/influx6/faux/httputil"

	"github.com/gokit/tenancykit/api/twofactorsessionapi"
	"github.com/gokit/tenancykit/db/types"
)

// TwoFactorSessionAPI augments the existing twofactorsessionapi.TwoFactorSessionHTTP with other
// methods that expose more functionality.
type TwoFactorSessionAPI struct {
	twofactorsessionapi.TwoFactorSessionHTTP
	Backend types.TwoFactorSessionBackend
}

// ValidateUserToken processing incoming one-time request to validate users
// code, this is useful for times of requiring a one-time authentication of
// user request.
//
// Params:
//		:user_code => "434544" // provide user generate auth code
//
func (tfs TwoFactorSessionAPI) ValidateUserToken(ctx *httputil.Context) error {
	currentUser, err := tenancykit.GetCurrentUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	// if user does not require twofactor auth, then skip this.
	if !currentUser.User.TwoFactorAuth {
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

	if err := currentUser.TwoFactor.ValidateOTP(userCode); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

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
//		:till_logout => true	// indicate whether session is lives till user logs out.
//
func (tfs TwoFactorSessionAPI) NewSession(ctx *httputil.Context) error {
	currentUser, err := tenancykit.GetCurrentUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	// If user already has twofactor session locked in, then skip.
	if currentUser.TFSession != nil {
		return nil
	}

	// if user does not require twofactor auth, then skip this.
	if !currentUser.User.TwoFactorAuth {
		return nil
	}

	// If we failed to load users two factor record then return error.
	if currentUser.TwoFactor == nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	tillLogout, _ := ctx.Bag().GetBool("till_logout")
	userCode, ok := ctx.Bag().GetString("user_code")
	if !ok {
		return httputil.HTTPError{
			Err:  errors.New("user_code param required"),
			Code: http.StatusBadRequest,
		}
	}

	// If we have an existing two-factor session, then validate code sent is correct,
	// then update current tfrecord and pass on the session record to context.
	if tfSession, err := tfs.Backend.GetByField(ctx, "user_id", currentUser.User.PublicID); err == nil {
		if err := currentUser.TwoFactor.ValidateOTP(userCode); err != nil {
			tfs.Backend.Delete(ctx, tfSession.PublicID)

			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusBadRequest,
			}
		}

		currentUser.TFSession = &tfSession
		return nil
	}

	if err := currentUser.TwoFactor.ValidateOTP(userCode); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	newTFSession := tenancykit.NewTwoFactorSession(currentUser.User.PublicID, true)

	// if we are dealing with a one time session run, then dont bother saving, just
	// return and let context have the new session.
	if !tillLogout {
		currentUser.TFSession = &newTFSession
		ctx.Bag().Set(tenancykit.ContextKeyCurrentUser, currentUser)
		return nil
	}

	if err := tfs.Backend.Create(ctx, newTFSession); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	currentUser.TFSession = &newTFSession
	ctx.Bag().Set(tenancykit.ContextKeyCurrentUser, currentUser)

	return nil
}
