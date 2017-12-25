package services

import (
	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/influx6/faux/httputil"
	"gopkg.in/fatih/set.v0"
)

// ActivityService implements a service for activity-based role authorization. It uses
// provided activty names to define giving permissions, which decouples roles and actions.
// Allow easily change to roles or permissions with minimal code changes and redeployment.
type ActivityService struct {
	available       *set.Set
	RolesBackend    api.RoleBackend
	ActivityBackend api.ActivityBackend
}

// New returns a new ActivityService instance.
func NewActivityService(roles types.RoleDBBackend, activities types.ActivityDBBackend) ActivityService {
	return ActivityService{
		available:       set.New(),
		RolesBackend:    api.RoleBackend{RoleDBBackend: roles},
		ActivityBackend: api.ActivityBackend{ActivityDBBackend: activities},
	}
}

// AfterActivityConfirmation defines a function type which is ran once a activity
// is found to be do-able by a giving user's roles. It returns an error if further
// qualification is not found. e.g
// 1. Provides a useful means of blocking a user even if allowed to edit his post
// from editing others post even if he has a EditPost activity in roles.
type AfterActivityConfirmation func(*httputil.Context, pkg.User) error

// EnsureAuthorized takes provided ActivateName, loading necessary activity data and validates authenticated user
// to be able to perform action/activity. It runs provided AfterActivityCOnfirmation function if provided.
func (as ActivityService) EnsureAuthorized(ctx *httputil.Context, name pkg.ActivityName, afterFind AfterActivityConfirmation) error {
	user, err := pkg.GetCurrentUser(ctx)
	if err != nil {
		return httputil.HTTPError{}
	}
	return nil
}
