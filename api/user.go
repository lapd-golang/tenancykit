package api

import (
	"errors"
	"net/http"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/userapi"
	"github.com/influx6/faux/httputil"
	"github.com/influx6/faux/metrics"
)

// UserAPI implements the http api for the user request-response api.
type UserAPI struct {
	userapi.UserHTTP
	Backend             types.UserDBBackend
	IsNotFoundErrorFunc func(error) bool
}

// NewUserAPI returns a new instance of UserAPI.
func NewUserAPI(m metrics.Metrics, users types.UserDBBackend) UserAPI {
	var api UserAPI
	api.Backend = users
	api.IsNotFoundErrorFunc = db.IsNotFoundError
	api.UserHTTP = userapi.New(m, backends.UserBackend{UserDBBackend: users})
	return api
}

// NewMultiTenantUserAPI returns a new instance of UserAPI.
func NewMultiTenantUserAPI(m metrics.Metrics, users types.UserDBBackend, tenants types.TenantDBBackend) UserAPI {
	var api UserAPI
	api.Backend = users
	api.IsNotFoundErrorFunc = db.IsNotFoundError
	api.UserHTTP = userapi.New(m, backends.MultiUserBackend{UserDBBackend: users, Tenants: tenants})
	return api
}

// isNotFoundError returns true/false if the giving error matches
// an expected ErrNotFound error. It helps detach us
// from code smell and import pollution.
func (u UserAPI) isNotFoundError(err error) bool {
	if u.IsNotFoundErrorFunc == nil {
		return false
	}

	return u.IsNotFoundErrorFunc(err)
}

// RetrieveUser attempts to use the provided public id parameter to load
// the expected user account into the current request context.
//
// Method: GET
// Params: :public_id
//
// Failure:
//	if User does not exists: 404
//
func (u UserAPI) RetrieveUser(ctx *httputil.Context) error {
	id, ok := ctx.Bag().GetString("public_id")
	if !ok {
		err := errors.New("expected public_id param")
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	// Retrieve user from db.
	user, err := u.Backend.Get(ctx.Context(), id)
	if err != nil {
		if !u.isNotFoundError(err) {
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
	ctx.Bag().Set(pkg.ContextKeyUser, user)
	return nil
}
