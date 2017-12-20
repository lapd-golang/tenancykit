package tenancykit

import (
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/tenantapi"
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
