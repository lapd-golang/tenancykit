package api

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/activityapi"
	"github.com/influx6/faux/metrics"
)

// ActivityAPI defines a wrapper around the activity http interface.
type ActivityAPI struct {
	Backend types.ActivityDBBackend
	activityapi.ActivityHTTP
}

// NewActivity returns a new instance of the ActivityAPI.
func NewActivityAPI(m metrics.Metrics, backend types.ActivityDBBackend) ActivityAPI {
	return ActivityAPI{
		Backend:      backend,
		ActivityHTTP: activityapi.New(m, ActivityBackend{ActivityDBBackend: backend}),
	}
}

// ActivityBackend is a wrapper to implement userapi.Backend interface methods for
// types.UserBackend.
type ActivityBackend struct {
	types.ActivityDBBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (u ActivityBackend) Create(ctx context.Context, act pkg.Activity) (pkg.Activity, error) {
	return act, u.ActivityDBBackend.Create(ctx, act)
}
