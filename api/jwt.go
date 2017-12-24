package api

import (
	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"
	"github.com/influx6/faux/metrics"
)

// IdentityAPI implements necessary tenantapi.HTTP with appropriate
// types.TenantBackend.
type IdentityAPI struct {
	userclaimjwt.IdentityHTTP
}

// NewIdentityAPI returns a new instance of IdentityAPI.
func NewIdentityAPI(m metrics.Metrics, id userclaimjwt.IdentityBackend) IdentityAPI {
	return IdentityAPI{
		IdentityHTTP: userclaimjwt.New(m, id),
	}
}
