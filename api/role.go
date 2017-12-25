package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gokit/tenancykit/pkg/db"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/roleapi"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// RoleAPI defines a wrapper around the activity http interface.
type RoleAPI struct {
	Backend RoleBackend
	roleapi.RoleHTTP
}

// NewRole returns a new instance of the RoleAPI.
func NewRoleAPI(m metrics.Metrics, rb types.RoleDBBackend) RoleAPI {
	backend := RoleBackend{RoleDBBackend: rb}
	return RoleAPI{
		Backend:  backend,
		RoleHTTP: roleapi.New(m, backend),
	}
}

// GetUserRoles retrieves the current existing user within the context either through
// the ContextKeyUser or ContextKeyCurrentUser and loads all roles associated with
// user into appropriate context field. It returns error if user is not found through
// either keys. It won't do anything if there exists, a []Role slice already in context.
func (r RoleAPI) GetUserRoles(ctx *httputil.Context) error {
	// Do roles already exists in context?
	if _, err := pkg.GetUserRoles(ctx); err == nil {
		return nil
	}

	// Attempt to get current user struct to load roles.
	if currentUser, err := pkg.GetCurrentUser(ctx); err == nil {
		// if user roles have already being loaded, then skip.
		if len(currentUser.User.Roles) == len(currentUser.Roles) {
			return nil
		}

		// User has no roles, so skip.
		if len(currentUser.User.Roles) == 0 {
			return nil
		}

		var roles []pkg.Role
		for _, id := range currentUser.User.Roles {
			role, err := r.Backend.Get(ctx.Context(), id)
			if err != nil {
				// if role does not exists, then skip.
				if db.IsNotFoundError(err) {
					continue
				}

				return httputil.HTTPError{
					Err:  err,
					Code: http.StatusInternalServerError,
				}
			}

			roles = append(roles, role)
		}

		ctx.Set(pkg.ContextKeyUserRoles, roles)

		currentUser.Roles = roles
		ctx.Set(pkg.ContextKeyCurrentUser, currentUser)
		return nil
	}

	// Attempt to get user struct to load roles.
	if user, err := pkg.GetUser(ctx); err == nil {

		// User has no roles, so skip.
		if len(user.Roles) == 0 {
			return nil
		}

		var roles []pkg.Role
		for _, id := range user.Roles {
			role, err := r.Backend.Get(ctx.Context(), id)
			if err != nil {
				// if role does not exists, then skip.
				if db.IsNotFoundError(err) {
					continue
				}

				return httputil.HTTPError{
					Err:  err,
					Code: http.StatusInternalServerError,
				}
			}

			roles = append(roles, role)
		}

		ctx.Set(pkg.ContextKeyUserRoles, roles)
		return nil
	}

	return httputil.HTTPError{
		Err:  errors.New("No user loaded in context"),
		Code: http.StatusInternalServerError,
	}
}

// GetRoleByName attempts to retrieve Role instance from db using role's name.
//
// Method: POST
// Request Param: :name
// Response: JSON
//
func (r RoleAPI) GetRoleByName(ctx *httputil.Context) error {
	roleName, ok := ctx.GetString("name")
	if !ok {
		return httputil.HTTPError{
			Err:  errors.New("no name param/parameter provided"),
			Code: http.StatusBadRequest,
		}
	}

	role, err := r.Backend.GetRole(ctx.Context(), roleName)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	return ctx.JSON(http.StatusOK, role)
}

// RoleBackend is a wrapper to implement userapi.Backend interface methods for
// types.UserBackend.
type RoleBackend struct {
	types.RoleDBBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (u RoleBackend) Create(ctx context.Context, roleName pkg.RoleName) (pkg.Role, error) {
	role := pkg.NewRole(roleName)
	return role, u.RoleDBBackend.Create(ctx, role)
}

// GetRole returns giving role with given name.
func (u RoleBackend) GetRole(ctx context.Context, name string) (pkg.Role, error) {
	return u.RoleDBBackend.GetByField(ctx, "name", name)
}
