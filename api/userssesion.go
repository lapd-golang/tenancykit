package api

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/usersessionapi"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// UserSessionAPI augments the existing usersessionapi.UserSessionHTTP with other
// methods that expose more functionality.
type UserSessionAPI struct {
	usersessionapi.UserSessionHTTP
	Backend             usersessionapi.UserSessionBackend
	SessionBackend      types.UserSessionDBBackend
	UserBackend         types.UserDBBackend
	TFBackend           types.TFRecordDBBackend
	TFSessionsBackend   types.TwoFactorSessionDBBackend
	IsNotFoundErrorFunc func(error) bool
}

// NewUserSessionAPI returns a new instance of UserSessionAPI.
func NewUserSessionAPI(m metrics.Metrics, users types.UserDBBackend, sessions types.UserSessionDBBackend, tfsessions types.TwoFactorSessionDBBackend, tf types.TFRecordDBBackend) UserSessionAPI {
	var api UserSessionAPI
	api.TFBackend = tf
	api.UserBackend = users
	api.SessionBackend = sessions
	api.IsNotFoundErrorFunc = db.IsNotFoundError
	api.TFSessionsBackend = tfsessions
	api.Backend = backends.UserSessionBackend{
		UserBackend:          users,
		UserSessionDBBackend: sessions,
	}

	api.UserSessionHTTP = usersessionapi.New(m, api.Backend)

	return api
}

// isNotFoundError returns true/false if the giving error matches
// an expected ErrNotFound error. It helps detach us
// from code smell and import pollution.
func (us UserSessionAPI) isNotFoundError(err error) bool {
	if us.IsNotFoundErrorFunc == nil {
		return false
	}

	return us.IsNotFoundErrorFunc(err)
}

// Login handles the processing of a log in request, creating a new session
// as needed for user after validation of credentials. It loads validated
// user into context after is successfully validates credentails.
func (us UserSessionAPI) Login(ctx *httputil.Context) error {
	if err := us.GetUser(ctx); err == nil {
		currentUser, err := pkg.GetCurrentUser(ctx)
		if err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		if currentUser.Session == nil {
			return httputil.HTTPError{
				Err:  errors.New("CurrentUser session can not be nil"),
				Code: http.StatusInternalServerError,
			}
		}

		return ctx.JSON(200, pkg.UserAuthorization{Token: currentUser.Session.UserSessionToken()})
	}

	body := ctx.Request().Body
	defer body.Close()

	var credentials pkg.CreateUserSession
	if err := json.NewDecoder(body).Decode(&credentials); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	newSession, err := us.Backend.Create(ctx.Context(), credentials)
	if err != nil {
		if httperr, ok := err.(httputil.HTTPError); ok {
			return httperr
		}

		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	user, err := us.UserBackend.Get(ctx.Context(), newSession.UserID)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	var currentUser pkg.CurrentUser
	currentUser.User = user
	currentUser.Session = &newSession

	if user.TwoFactorAuth {
		tfrecord, err := us.TFBackend.GetByField(ctx.Context(), "user_id", user.PublicID)
		if err != nil {
			if !us.isNotFoundError(err) {
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

		currentUser.TwoFactor = &tfrecord
	}

	// set current user into context.
	ctx.Bag().Set(pkg.ContextKeyCurrentUser, currentUser)
	ctx.Set(pkg.ContextKeyUserAuthorization, pkg.UserAuthorization{Token: newSession.UserSessionToken()})

	authValue := fmt.Sprintf("Bearer %s", newSession.UserSessionToken())

	//ctx.SetCookie(&http.Cookie{
	//	Name:    "session-token",
	//	Expires: newSession.Expires,
	//	Value:   base64.StdEncoding.EncodeToString(authValue),
	//	Path:    "/",
	//})

	ctx.SetHeader("Authorization", authValue)

	return ctx.JSON(200, pkg.UserAuthorization{Token: newSession.UserSessionToken()})
}

// Logout ends a existing users session, and deletes all existing twofactor
// sessions thereby removing user from logged in state. It uses the current user
// retrieved either from context or through provided Authorization header.
//
// Params: None
//
// Headers:
//		Authorization: Bearer <TOKEN>
//
//		where <TOKEN> = <USERID>:<SESSIONTOKEN>
//
func (us UserSessionAPI) Logout(ctx *httputil.Context) error {
	if err := us.GetUser(ctx); err != nil {
		return err
	}

	currentUser, err := pkg.GetCurrentUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusUnauthorized,
		}
	}

	if currentUser.Session == nil {
		cuSession, err := us.SessionBackend.GetByField(ctx.Context(), "user_id", currentUser.User.PublicID)
		if err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		currentUser.Session = &cuSession
	}

	if currentUser.TFSession == nil && currentUser.User.TwoFactorAuth {
		tfSession, err := us.TFSessionsBackend.GetByField(ctx.Context(), "user_id", currentUser.User.PublicID)
		if err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		currentUser.TFSession = &tfSession
	}

	if err := us.Backend.Delete(ctx.Context(), currentUser.Session.PublicID); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if currentUser.TFSession != nil {
		if err = us.TFSessionsBackend.Delete(ctx.Context(), currentUser.User.PublicID); err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

// Authenticate attempts to authorize giving Authorization header as
// a valid session, it returns 200 if the Authorization header
// is valid and is still unexpired.
//
// Params: None
//
// Headers:
//		Authorization: Bearer <TOKEN>
//
//		where <TOKEN> = <USERID>:<SESSIONTOKEN>
//
// Response:
//         Success: 200
//         Failure: 501/404
func (us UserSessionAPI) Authenticate(ctx *httputil.Context) error {
	if err := us.GetUser(ctx); err != nil {
		return err
	}

	ctx.Status(http.StatusOK)
	return nil
}

// GetUser attempts to retrieve user details from provided Authorization header
// if not is available, then request is rejected. Once user is assert to
// valdate, then users tenant and twofactor records are pulled also and
// set into context.
//
// Params: None
//
// Headers:
//		Authorization: Bearer <TOKEN>
//
//		where <TOKEN> = <USERID>:<SESSIONTOKEN>
//
func (us UserSessionAPI) GetUser(ctx *httputil.Context) error {
	userID, userToken, err := us.GetAuthorizationCredentials(ctx)
	if err != nil {
		return err
	}

	// If we are already logged in and user's token is valid and session is not
	// expired, then go ahead.
	if lastCurrentUser, err := pkg.GetCurrentUser(ctx); err == nil && lastCurrentUser.User.PublicID == userID {
		ctx.Bag().Set(pkg.ContextKeyUser, lastCurrentUser.User)
		if lastCurrentUser.Session != nil && !lastCurrentUser.Session.Expired() {
			if lastCurrentUser.Session.ValidateToken(userToken) {

				// Validate session is still valid in db.
				if _, err := us.SessionBackend.Get(ctx.Context(), lastCurrentUser.Session.PublicID); err != nil {
					return httputil.HTTPError{
						Err:  err,
						Code: http.StatusUnauthorized,
					}
				}

				return nil
			}
		}
	}

	user, err := us.UserBackend.Get(ctx.Context(), userID)
	if err != nil {
		if !us.isNotFoundError(err) {
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

	ctx.Bag().Set(pkg.ContextKeyUser, user)

	session, err := us.SessionBackend.GetByField(ctx.Context(), "user_id", user.PublicID)
	if err != nil {
		if !us.isNotFoundError(err) {
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

	if session.Expired() {
		if err := us.SessionBackend.Delete(ctx.Context(), session.PublicID); err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		return httputil.HTTPError{
			Err:  errors.New("Expired session"),
			Code: http.StatusUnauthorized,
		}
	}

	if !session.ValidateToken(userToken) {
		return httputil.HTTPError{
			Err:  errors.New("Invalid user token from Authorization header"),
			Code: http.StatusUnauthorized,
		}
	}

	var currentUser pkg.CurrentUser
	currentUser.User = user
	currentUser.Session = &session

	// Attempt to retrieve twofactor user record if user has one.
	if tf, err := us.TFBackend.GetByField(ctx.Context(), "user_id", currentUser.User.PublicID); err == nil {
		currentUser.TwoFactor = &tf
	}

	// set current user into context.
	ctx.Bag().Set(pkg.ContextKeyCurrentUser, currentUser)
	return nil
}

// GetAuthorization returns authorization value for giving request from either it's header or
// cookies if found, else returns an error.
func (us UserSessionAPI) GetAuthorization(ctx *httputil.Context) (string, error) {
	if ctx.HasHeader("Authorization", "") {
		return ctx.GetHeader("Authorization"), nil
	}

	for _, cookie := range ctx.Cookies() {
		if strings.ToLower(cookie.Name) == "session-token" {
			if val, err := base64.StdEncoding.DecodeString(cookie.Value); err == nil {
				return string(val), nil
			}
			return cookie.Value, nil
		}
	}

	return "", errors.New("no valid authorization found")
}

// GetAuthorizationCredentials returns the current Authentication User ID and Auth token else an error.
func (us UserSessionAPI) GetAuthorizationCredentials(ctx *httputil.Context) (string, string, error) {
	authorization, err := us.GetAuthorization(ctx)
	if err != nil {
		return "", "", httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	authtype, token, err := httputil.ParseAuthorization(authorization)
	if err != nil {
		return "", "", httputil.HTTPError{
			Err:  errors.New("Invalid Authorization header"),
			Code: http.StatusBadRequest,
		}
	}

	if authtype != "Bearer" {
		err := errors.New("Only `Bearer` Authorization supported")
		return "", "", httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	tokens, err := httputil.ParseTokens(token)
	if err != nil {
		return "", "", httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	if len(tokens) < 2 {
		err := errors.New("invalid token count, expected 2 token values")
		return "", "", httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	userID, userToken := tokens[0], tokens[1]
	return userID, userToken, nil
}
