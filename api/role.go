package api

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/roleapi"
	"github.com/influx6/faux/metrics"
)

// RoleAPI defines a wrapper around the activity http interface.
type RoleAPI struct {
	Backend types.RoleDBBackend
	roleapi.RoleHTTP
}

// NewRole returns a new instance of the RoleAPI.
func NewRoleAPI(m metrics.Metrics, backend types.RoleDBBackend) RoleAPI {
	return RoleAPI{
		Backend:  backend,
		RoleHTTP: roleapi.New(m, RoleBackend{RoleDBBackend: backend}),
	}
}

// RoleBackend is a wrapper to implement userapi.Backend interface methods for
// types.UserBackend.
type RoleBackend struct {
	types.RoleDBBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (u RoleBackend) Create(ctx context.Context, act pkg.Role) (pkg.Role, error) {
	return act, u.RoleDBBackend.Create(ctx, act)
}
