package usersessiondbbackendimpl

import (
	pkg "github.com/gokit/tenancykit/pkg"

	context "github.com/influx6/faux/context"
)

// UserSessionDBBackendImpl defines a concrete struct which implements the methods for the
// UserSessionDBBackend interface. All methods will panic, so add necessary internal logic.
type UserSessionDBBackendImpl struct {
	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem pkg.UserSession) error

	GetFunc func(ctx context.Context, publicID string) (pkg.UserSession, error)

	UpdateFunc func(ctx context.Context, publicID string, elem pkg.UserSession) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]pkg.UserSession, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (pkg.UserSession, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.UserSession, int, error)
}

// Delete implements the UserSessionDBBackend.Delete() method for UserSessionDBBackendImpl.
func (impl UserSessionDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the UserSessionDBBackend.Create() method for UserSessionDBBackendImpl.
func (impl UserSessionDBBackendImpl) Create(ctx context.Context, elem pkg.UserSession) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the UserSessionDBBackend.Get() method for UserSessionDBBackendImpl.
func (impl UserSessionDBBackendImpl) Get(ctx context.Context, publicID string) (pkg.UserSession, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the UserSessionDBBackend.Update() method for UserSessionDBBackendImpl.
func (impl UserSessionDBBackendImpl) Update(ctx context.Context, publicID string, elem pkg.UserSession) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the UserSessionDBBackend.GetAllByOrder() method for UserSessionDBBackendImpl.
func (impl UserSessionDBBackendImpl) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.UserSession, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the UserSessionDBBackend.GetByField() method for UserSessionDBBackendImpl.
func (impl UserSessionDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (pkg.UserSession, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the UserSessionDBBackend.GetAll() method for UserSessionDBBackendImpl.
func (impl UserSessionDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.UserSession, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
