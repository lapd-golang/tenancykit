package pkg

import (
	"errors"

	"github.com/influx6/faux/httputil"
)

// variables of context keys for the tenancypackage.
var (
	ContextKeyUser              = contextKey("user")
	ContextKeyUserRoles         = contextKey("user-roles")
	ContextKeyUserAuthorization = contextKey("user-authorization")
	ContextKeyTenant            = contextKey("tenant")
	ContextKeyCurrentUser       = contextKey("current-user")
	ContextKeyUserSession       = contextKey("user-session")
	ContextKeyTFRecord          = contextKey("user-two-factor-record")
	ContextKeyTFSession         = contextKey("user-two-factor-session")
)

// contextKey contains provided key value for creating context based key names.
type contextKey string

// String returns string content printable for contextkey.
func (c contextKey) String() string {
	return "Tenancykit context key: " + string(c)
}

// GetUserRoles retrieves currently stored UserSession struct in provided context.
func GetUserRoles(ctx *httputil.Context) ([]Role, error) {
	receivedVal, ok := ctx.Get(ContextKeyUserRoles)
	if !ok {
		return nil, errors.New("not found")
	}

	roles, ok := receivedVal.([]Role)
	if !ok {
		return nil, errors.New("received value is not users []Role")
	}

	return roles, nil
}

// GetUserSession retrieves currently stored UserSession struct in provided context.
func GetUserSession(ctx *httputil.Context) (UserSession, error) {
	receivedVal, ok := ctx.Get(ContextKeyUserSession)
	if !ok {
		return UserSession{}, errors.New("not found")
	}

	userSession, ok := receivedVal.(UserSession)
	if !ok {
		return UserSession{}, errors.New("received value is not UserSession")
	}

	return userSession, nil
}

// GetUserAuthorization retrieves currently stored userAuthorization struct in provided context.
func GetUserAuthorization(ctx *httputil.Context) (UserAuthorization, error) {
	receivedVal, ok := ctx.Get(ContextKeyUserAuthorization)
	if !ok {
		return UserAuthorization{}, errors.New("not found")
	}

	auth, ok := receivedVal.(UserAuthorization)
	if !ok {
		return UserAuthorization{}, errors.New("received value is not UserAuthorization")
	}

	return auth, nil
}

// GetCurrentUser retrieves currently stored CurrentUser struct in provided context.
func GetCurrentUser(ctx *httputil.Context) (CurrentUser, error) {
	receivedVal, ok := ctx.Get(ContextKeyCurrentUser)
	if !ok {
		return CurrentUser{}, errors.New("not found")
	}

	currentUser, ok := receivedVal.(CurrentUser)
	if !ok {
		return CurrentUser{}, errors.New("received value is not CurrentUser")
	}

	return currentUser, nil
}

// GetTwoFactorRecord retrieves currently stored user in provided context.
func GetTwoFactorRecord(ctx *httputil.Context) (TFRecord, error) {
	receivedVal, ok := ctx.Get(ContextKeyTFRecord)
	if !ok {
		return TFRecord{}, errors.New("not found")
	}

	tf, ok := receivedVal.(TFRecord)
	if !ok {
		return TFRecord{}, errors.New("received value is not TFRecord")
	}

	return tf, nil
}

// GetUser retrieves currently stored user in provided context.
func GetUser(ctx *httputil.Context) (User, error) {
	receivedVal, ok := ctx.Get(ContextKeyUser)
	if !ok {
		return User{}, errors.New("not found")
	}

	user, ok := receivedVal.(User)
	if !ok {
		return User{}, errors.New("received value is not User")
	}

	return user, nil
}

// GetTenant retrieves currently stored tenant in provided context.
func GetTenant(ctx *httputil.Context) (Tenant, error) {
	receivedVal, ok := ctx.Get(ContextKeyTenant)
	if !ok {
		return Tenant{}, errors.New("not found")
	}

	tenant, ok := receivedVal.(Tenant)
	if !ok {
		return Tenant{}, errors.New("received value is not Tenant")
	}

	return tenant, nil
}
