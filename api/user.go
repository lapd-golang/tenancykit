package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gokit/tenancykit/pkg"
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
	api.UserHTTP = userapi.New(m, UserBackend{UserDBBackend: users})
	return api
}

// NewMultiTenantUserAPI returns a new instance of UserAPI.
func NewMultiTenantUserAPI(m metrics.Metrics, users types.UserDBBackend, tenants types.TenantDBBackend) UserAPI {
	var api UserAPI
	api.Backend = users
	api.IsNotFoundErrorFunc = db.IsNotFoundError
	api.UserHTTP = userapi.New(m, MultiUserBackend{UserDBBackend: users, Tenants: tenants})
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

func (u UserAPI) UpdatePassword(ctx *httputil.Context) error {
	if err := u.RetrieveUser(ctx); err != nil {
		return err
	}

	var updater pkg.UpdateUserPassword
	if err := json.NewDecoder(ctx.Body()).Decode(&updater); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	user, err := pkg.GetUser(ctx)
	if err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	if err = user.ChangePassword(updater.Password); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusUnauthorized,
		}
	}

	if err := u.Backend.Update(ctx.Context(), user.PublicID, user); err != nil {
		return httputil.HTTPError{
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}

	ctx.Status(http.StatusNoContent)
	return nil
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

// MultiTenantUserBackend is a wrapper to implement userapi.Backend interface methods for
// types.UserBackend.
type MultiUserBackend struct {
	types.UserDBBackend
	Tenants types.TenantDBBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (u MultiUserBackend) Create(ctx context.Context, nt pkg.CreateUser) (pkg.User, error) {
	if err := nt.Validate(true); err != nil {
		return pkg.User{}, err
	}

	// Validate that user's tenant exists.
	if _, err := u.Tenants.Get(ctx, nt.TenantID); err != nil {
		return pkg.User{}, err
	}

	newUser, err := pkg.NewUser(nt)
	if err != nil {
		return newUser, err
	}

	return newUser, u.UserDBBackend.Create(ctx, newUser)
}

// UserBackend is a wrapper to implement userapi.Backend interface methods for
// types.UserBackend.
type UserBackend struct {
	types.UserDBBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (u UserBackend) Create(ctx context.Context, nt pkg.CreateUser) (pkg.User, error) {
	if err := nt.Validate(false); err != nil {
		return pkg.User{}, err
	}

	newUser, err := pkg.NewUser(nt)
	if err != nil {
		return newUser, err
	}

	return newUser, u.UserDBBackend.Create(ctx, newUser)
}
