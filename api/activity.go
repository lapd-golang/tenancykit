package api

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/gokit/tenancykit/pkg/db"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/activityapi"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// ActivityAPI defines a wrapper around the activity http interface.
type ActivityAPI struct {
	Backend ActivityBackend
	activityapi.ActivityHTTP
}

// NewActivity returns a new instance of the ActivityAPI.
func NewActivityAPI(m metrics.Metrics, ab types.ActivityDBBackend) ActivityAPI {
	backend := ActivityBackend{ActivityDBBackend: ab}
	return ActivityAPI{
		Backend:      backend,
		ActivityHTTP: activityapi.New(m, backend),
	}
}

// GetActivityByName attempts to retrieve Role instance from db using role's name.
//
// Method: POST
// Request Param: :name
// Response: JSON
//
func (av ActivityAPI) GetActivityByName(ctx *httputil.Context) error {
	activityName, ok := ctx.GetString("name")
	if !ok {
		return httputil.HTTPError{
			Err:  errors.New("no name param/parameter provided"),
			Code: http.StatusBadRequest,
		}
	}

	act, err := av.Backend.GetActivity(ctx.Context(), pkg.ActivityName(activityName))
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	return ctx.JSON(http.StatusOK, act)
}

// ActivityBackend is a wrapper to implement userapi.Backend interface methods for
// types.UserBackend.
type ActivityBackend struct {
	types.ActivityDBBackend
}

// GetActivity returns giving activity with given name.
func (u ActivityBackend) GetActivity(ctx context.Context, name pkg.ActivityName) (pkg.Activity, error) {
	return u.ActivityDBBackend.GetByField(ctx, "name", string(name))
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (u ActivityBackend) Create(ctx context.Context, actName pkg.ActivityName) (pkg.Activity, error) {
	act := pkg.NewActivity(actName)
	return act, u.ActivityDBBackend.Create(ctx, act)
}

// ActivitySet defines a set for activity list where activist loaded would not be loaded twice.
type ActivitySet struct {
	backend types.ActivityDBBackend
	ml      sync.RWMutex
	sets    map[string]pkg.Activity
}

// NewActivitySet returns a new instance of ActivitySet.
func NewActivitySet(backend types.ActivityDBBackend) *ActivitySet {
	return &ActivitySet{
		backend: backend,
		sets:    map[string]pkg.Activity{},
	}
}

// GetActivity returns giving activity with given name.
func (as ActivitySet) GetActivity(ctx context.Context, name pkg.ActivityName) (pkg.Activity, error) {
	as.ml.RLock()
	if act, ok := as.sets[string(name)]; ok {
		as.ml.RUnlock()
		return act, nil
	}
	as.ml.RUnlock()

	as.ml.Lock()
	defer as.ml.Unlock()

	act, err := as.backend.GetByField(ctx, "name", string(name))
	if err != nil {
		// if it's not a not found error then return error
		if !db.IsNotFoundError(err) {
			return pkg.Activity{}, err
		}

		act := pkg.NewActivity(name)
		if err = as.backend.Create(ctx, act); err != nil {
			return pkg.Activity{}, err
		}

		as.sets[act.Name] = act
		return act, nil
	}

	as.sets[act.Name] = act
	return act, nil
}
