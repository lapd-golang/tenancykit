package services

import (
	"errors"
	"net/http"

	"github.com/gokit/tenancykit/pkg/db"

	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/influx6/faux/httputil"
)

// AfterActivityConfirmation defines a function type which is ran once a activity
// is found to be do-able by a giving user's roles. It returns an error if further
// qualification is not found. e.g
// 1. Provides a useful means of blocking a user even if allowed to edit his post
// from editing others post even if he has a EditPost activity in roles.
type AfterActivityConfirmation func(*httputil.Context, pkg.User) error

// ActivityService implements a service for activity-based role authorization. It uses
// provided activty names to define giving permissions, which decouples roles and actions.
// Allow easily change to roles or permissions with minimal code changes and redeployment.
type ActivityService struct {
	set             *api.ActivitySet
	RolesBackend    api.RoleBackend
	ActivityBackend api.ActivityBackend
}

// New returns a new ActivityService instance.
func NewActivityService(roles types.RoleDBBackend, activities types.ActivityDBBackend) ActivityService {
	return ActivityService{
		set:             api.NewActivitySet(activities),
		RolesBackend:    api.RoleBackend{RoleDBBackend: roles},
		ActivityBackend: api.ActivityBackend{ActivityDBBackend: activities},
	}
}

// EnsureAuthorizedMiddleware returns a httputil.Middleware which wraps next handler and certifies that incoming user
// meets activity requirements before calling next handler to proceed. It provides a middleware guard for authorized actions.
func (as ActivityService) EnsureAuthorizedMiddleware(name pkg.ActivityName, afterFind AfterActivityConfirmation) httputil.Middleware {
	return func(next httputil.Handler) httputil.Handler {
		return httputil.OnNoError(func(ctx *httputil.Context) error {
			return as.EnsureAuthorized(ctx, name, afterFind)
		}, next)
	}
}

// EnsureAuthorized takes provided ActivateName, loading necessary activity data and validates authenticated user
// to be able to perform action/activity. It runs provided AfterActivityCOnfirmation function if provided.
func (as ActivityService) EnsureAuthorized(ctx *httputil.Context, name pkg.ActivityName, afterFind AfterActivityConfirmation) error {
	activity, err := as.set.GetActivity(ctx.Context(), name)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if currentUser, err := pkg.GetCurrentUser(ctx); err == nil {
		if len(currentUser.User.Roles) == 0 {
			return errors.New("User has no roles attached")
		}

		// if all roles have being loaded, then use this.
		if len(currentUser.User.Roles) == len(currentUser.Roles) {
			return as.checkRoles(ctx, currentUser.User, activity, currentUser.Roles, afterFind)
		}

		roles, err := as.RolesBackend.GetRoles(ctx.Context(), false, currentUser.User.Roles)
		if err != nil {
			if db.IsNotFoundError(err) {
				return httputil.HTTPError{
					Err:  err,
					Code: http.StatusBadRequest,
				}
			}
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		currentUser.Roles = roles
		ctx.Set(pkg.ContextKeyCurrentUser, currentUser)
		return as.checkRoles(ctx, currentUser.User, activity, currentUser.Roles, afterFind)
	}

	if user, err := pkg.GetUser(ctx); err == nil {
		if roles, err := pkg.GetUserRoles(ctx); err == nil {
			return as.checkRoles(ctx, user, activity, roles, afterFind)
		}

		// Load roles of current user and add into context.
		roles, err := as.RolesBackend.GetRoles(ctx.Context(), false, user.Roles)
		if err != nil {
			if db.IsNotFoundError(err) {
				return httputil.HTTPError{
					Err:  err,
					Code: http.StatusBadRequest,
				}
			}
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		ctx.Set(pkg.ContextKeyUserRoles, roles)
		return as.checkRoles(ctx, user, activity, roles, afterFind)
	}

	return httputil.HTTPError{
		Code: http.StatusBadRequest,
		Err:  errors.New("no user loaded from authenticated session"),
	}
}

// checkRoles validates giving activity with roles list and runs the after function if provided. It is a helper method.
func (as ActivityService) checkRoles(ctx *httputil.Context, user pkg.User, activity pkg.Activity, roles []pkg.Role, afterFind AfterActivityConfirmation) error {
	for _, role := range roles {
		if role.CanPerformAll(activity) {
			if afterFind == nil {
				return nil
			}

			// if after function returns error then deny user access here.
			if err := afterFind(ctx, user); err != nil {
				if herr, ok := err.(httputil.HTTPError); ok {
					return herr
				}
				return httputil.HTTPError{
					Err:  err,
					Code: http.StatusBadRequest,
				}
			}
		}
	}

	// if we got here, then no role with suitable activity right found.
	return httputil.HTTPError{
		Err:  errors.New("user does not have right to Activity"),
		Code: http.StatusBadRequest,
	}
}
