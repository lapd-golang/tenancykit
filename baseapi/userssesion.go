package baseapi

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/api/usersessionapi"
	"github.com/gokit/tenancykit/backends"
	"github.com/gokit/tenancykit/db"
	"github.com/gokit/tenancykit/db/types"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// UserSessionAPI augments the existing usersessionapi.TFRecordHTTP with other
// methods that expose more functionality.
type UserSessionAPI struct {
	usersessionapi.UserSessionHTTP
	Backend             usersessionapi.UserSessionBackend
	SessionBackend      types.UserSessionDBBackend
	UserBackend         types.UserDBBackend
	TenantBackend       types.TenantDBBackend
	TFBackend           types.TFRecordDBBackend
	TFSessionsBackend   types.TwoFactorSessionDBBackend
	IsNotFoundErrorFunc func(error) bool
}

// NewUserSessionAPI returns a new instance of UserSessionAPI.
func NewUserSessionAPI(m metrics.Metrics, users types.UserDBBackend, sessions types.UserSessionDBBackend, tfsessions types.TwoFactorSessionDBBackend, tenants types.TenantDBBackend, tf types.TFRecordDBBackend) UserSessionAPI {
	var api UserSessionAPI
	api.TFBackend = tf
	api.TenantBackend = tenants
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
		return nil
	}

	body := ctx.Request().Body
	defer body.Close()

	var credentials tenancykit.CreateUserSession
	if err := json.NewDecoder(body).Decode(&credentials); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	newSession, err := us.Backend.Create(ctx, credentials)
	if err != nil {
		if httperr, ok := err.(httputil.HTTPError); ok {
			return httperr
		}

		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	user, err := tenancykit.GetUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	var currentUser tenancykit.CurrentUser
	currentUser.User = user
	currentUser.Session = &newSession

	// Rerieve user's tenant.
	currentUser.Tenant, err = us.TenantBackend.Get(ctx, user.TenantID)
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

	// set current user into context.
	ctx.Bag().Set(tenancykit.ContextKeyCurrentUser, currentUser)

	return nil
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

	currentUser, err := tenancykit.GetCurrentUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusUnauthorized,
		}
	}

	if currentUser.Session == nil {
		cuSession, err := us.SessionBackend.GetByField(ctx, "user_id", currentUser.User.PublicID)
		if err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		currentUser.Session = &cuSession
	}

	if currentUser.TFSession == nil {
		tfSession, err := us.TFSessionsBackend.GetByField(ctx, "user_id", currentUser.User.PublicID)
		if err != nil {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		currentUser.TFSession = &tfSession
	}

	if err := us.Backend.Delete(ctx, currentUser.Session.PublicID); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if err = us.TFSessionsBackend.Delete(ctx, currentUser.TFSession.PublicID); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

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
	authorization, err := us.GetAuthorization(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	authtype, token, err := httputil.ParseAuthorization(authorization)
	if err != nil {
		return httputil.HTTPError{
			Err:  errors.New("Invalid Authorization header"),
			Code: http.StatusBadRequest,
		}
	}

	if authtype != "Bearer" {
		err := errors.New("Only `Bearer` Authorization supported")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	tokens, err := httputil.ParseTokens(token)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	if len(tokens) < 2 {
		err := errors.New("invalid token count, expected 2 token values")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	userID, userToken := tokens[0], tokens[1]

	// If we are already logged in and user's token is valid and session is not
	// expired, then go ahead.
	if lastCurrentUser, err := tenancykit.GetCurrentUser(ctx); err == nil {
		if lastCurrentUser.Session != nil && !lastCurrentUser.Session.Expired() {
			if lastCurrentUser.Session.ValidateToken(userToken) {
				return nil
			}
		}
	}

	user, err := us.UserBackend.Get(ctx, userID)
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

	session, err := us.SessionBackend.GetByField(ctx, "user_id", user.PublicID)
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
		if err := us.SessionBackend.Delete(ctx, session.PublicID); err != nil {
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

	if !session.ValidateToken(token) {
		return httputil.HTTPError{
			Err:  errors.New("Invalid user token from Authorization header"),
			Code: http.StatusUnauthorized,
		}
	}

	var currentUser tenancykit.CurrentUser
	currentUser.User = user
	currentUser.Session = &session

	// Rerieve user's tenant.
	currentUser.Tenant, err = us.TenantBackend.Get(ctx, user.TenantID)
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

	// Attempt to retrieve twofactor user record if user has one.
	if tf, err := us.TFBackend.GetByField(ctx, "user_id", currentUser.User.PublicID); err == nil {
		currentUser.TwoFactor = &tf
	}

	// set current user into context.
	ctx.Bag().Set(tenancykit.ContextKeyCurrentUser, currentUser)
	return nil
}

// GetAuthorization returns authorization value for giving request.
func (us UserSessionAPI) GetAuthorization(ctx *httputil.Context) (string, error) {
	if ctx.HasHeader("Authorization", "") {
		return ctx.GetHeader("Authorization"), nil
	}

	for _, cookie := range ctx.Cookies() {
		if strings.ToLower(cookie.Name) == "authorization" {
			val, err := base64.StdEncoding.DecodeString(cookie.Value)
			return string(val), err
		}
	}

	return "", errors.New("no valid authorization found")
}
