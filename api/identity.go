package api

import (
	"errors"
	"net/http"

	"github.com/gokit/tenancykit/pkg"

	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// IdentityAPI implements necessary tenantapi.HTTP with appropriate
// types.TenantBackend.
type IdentityAPI struct {
	userclaimjwt.IdentityHTTP
	Backend userclaimjwt.IdentityBackend
}

// NewIdentityAPI returns a new instance of IdentityAPI.
func NewIdentityAPI(m metrics.Metrics, id userclaimjwt.IdentityBackend) IdentityAPI {
	return IdentityAPI{
		Backend:      id,
		IdentityHTTP: userclaimjwt.New(m, id),
	}
}

// AuthenticateUser runs the IdentityBackend.Authorize method and loads user's
// data and roles into context. If authorization is wrong then an HTTPError is returned
// with appropriate status code and error response, but if successful, then status code
// and response is left to user and only context values for tenant, user and user roles
// are set.
func (id IdentityAPI) AuthenticateUser(ctx *httputil.Context) error {
	var authHeader string
	if ctx.HasHeader("Authorization", "") {
		authHeader = ctx.GetHeader("Authorization")
	}

	if authHeader == "" {
		return httputil.HTTPError{
			Err:  userclaimjwt.ErrNoJWTAuthorizationToken,
			Code: http.StatusBadRequest,
		}
	}

	authType, token, err := httputil.ParseAuthorization(authHeader)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	if authType != "Bearer" {
		err := errors.New("only 'Bearer' authorization allowed")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	claim, err := id.Backend.Get(ctx.Context(), token)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusUnauthorized,
		}
	}

	ctx.Set(pkg.ContextKeyUser, claim.Data.User)
	ctx.Set(pkg.ContextKeyTenant, claim.Data.Tenant)
	ctx.Set(pkg.ContextKeyUserRoles, claim.Data.Roles)
	return nil
}
