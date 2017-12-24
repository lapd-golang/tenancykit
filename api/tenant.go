package api

import (
	"errors"
	"net/http"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/tenantapi"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// TenantAPI implements necessary tenantapi.HTTP with appropriate
// types.TenantBackend.
type TenantAPI struct {
	tenantapi.TenantHTTP
	Backend             types.TenantDBBackend
	IsNotFoundErrorFunc func(error) bool
}

// NewTenantAPI returns a new instance of TenantAPI.
func NewTenantAPI(m metrics.Metrics, tenant types.TenantDBBackend) TenantAPI {
	return TenantAPI{
		IsNotFoundErrorFunc: db.IsNotFoundError,
		TenantHTTP: tenantapi.New(m, backends.TenantBackend{
			TenantDBBackend: tenant,
		}),
	}
}

// isNotFoundError returns true/false if the giving error matches
// an expected ErrNotFound error. It helps detach us
// from code smell and import pollution.
func (t TenantAPI) isNotFoundError(err error) bool {
	if t.IsNotFoundErrorFunc == nil {
		return false
	}

	return t.IsNotFoundErrorFunc(err)
}

// RetrieveTenant attempts to use the provided public id parameter to load
// the expected tenant account into the current request context.
//
// Method: GET
// Params: :public_id
//
// Failure:
//	if Tenant does not exists: 404
//
func (t TenantAPI) RetrieveTenant(ctx *httputil.Context) error {
	id, ok := ctx.Bag().GetString("public_id")
	if !ok {
		err := errors.New("expected public_id param")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	// Retrieve user from db.
	tenant, err := t.Backend.Get(ctx.Context(), id)
	if err != nil {
		if !t.isNotFoundError(err) {
			return httputil.HTTPError{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusNotFound,
		}
	}

	// set current user into context.
	ctx.Bag().Set(pkg.ContextKeyTenant, tenant)
	return nil
}
