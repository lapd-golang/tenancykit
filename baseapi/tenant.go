package baseapi

import (
	"github.com/gokit/tenancykit/api/tenantapi"
	"github.com/gokit/tenancykit/backends"
	"github.com/gokit/tenancykit/db/types"
	"github.com/influx6/faux/metrics"
)

// TenantAPI implements necessary tenantapi.HTTP with appropriate
// types.TenantBackend.
type TenantAPI struct {
	tenantapi.TenantHTTP
}

// NewTenantAPI returns a new instance of TenantAPI.
func NewTenantAPI(m metrics.Metrics, tenant types.TenantDBBackend) TenantAPI {
	return TenantAPI{
		TenantHTTP: tenantapi.New(m, backends.TenantBackend{
			TenantDBBackend: tenant,
		}),
	}
}
