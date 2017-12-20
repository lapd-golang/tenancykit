package userdbbackendimpl

import (
	context "github.com/influx6/faux/context"

	tenancykit "github.com/gokit/tenancykit"
)

// UserDBBackendImpl defines a concrete struct which implements the methods for the
// UserDBBackend interface. All methods will panic, so add necessary internal logic.
type UserDBBackendImpl struct {
	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem tenancykit.User) error

	GetFunc func(ctx context.Context, publicID string) (tenancykit.User, error)

	UpdateFunc func(ctx context.Context, publicID string, elem tenancykit.User) error

	GetAllByOrderFunc func(ctx context.Context, order string) ([]tenancykit.User, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (tenancykit.User, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.User, int, error)
}

// Delete implements the UserDBBackend.Delete() method for UserDBBackendImpl.
func (impl UserDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the UserDBBackend.Create() method for UserDBBackendImpl.
func (impl UserDBBackendImpl) Create(ctx context.Context, elem tenancykit.User) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the UserDBBackend.Get() method for UserDBBackendImpl.
func (impl UserDBBackendImpl) Get(ctx context.Context, publicID string) (tenancykit.User, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the UserDBBackend.Update() method for UserDBBackendImpl.
func (impl UserDBBackendImpl) Update(ctx context.Context, publicID string, elem tenancykit.User) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the UserDBBackend.GetAllByOrder() method for UserDBBackendImpl.
func (impl UserDBBackendImpl) GetAllByOrder(ctx context.Context, order string) ([]tenancykit.User, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order)
	return ret1, ret2

}

// GetByField implements the UserDBBackend.GetByField() method for UserDBBackendImpl.
func (impl UserDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (tenancykit.User, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the UserDBBackend.GetAll() method for UserDBBackendImpl.
func (impl UserDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.User, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
