package roledbbackendimpl

import (
	context "context"

	pkg "github.com/gokit/tenancykit/pkg"
)

// RoleDBBackendImpl defines a concrete struct which implements the methods for the
// RoleDBBackend interface. All methods will panic, so add necessary internal logic.
type RoleDBBackendImpl struct {
	CountFunc func(ctx context.Context) (int, error)

	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem pkg.Role) error

	GetFunc func(ctx context.Context, publicID string) (pkg.Role, error)

	UpdateFunc func(ctx context.Context, publicID string, elem pkg.Role) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]pkg.Role, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (pkg.Role, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Role, int, error)
}

// Count implements the RoleDBBackend.Count() method for RoleDBBackendImpl.
func (impl RoleDBBackendImpl) Count(ctx context.Context) (int, error) {

	ret1, ret2 := impl.CountFunc(ctx)
	return ret1, ret2

}

// Delete implements the RoleDBBackend.Delete() method for RoleDBBackendImpl.
func (impl RoleDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the RoleDBBackend.Create() method for RoleDBBackendImpl.
func (impl RoleDBBackendImpl) Create(ctx context.Context, elem pkg.Role) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the RoleDBBackend.Get() method for RoleDBBackendImpl.
func (impl RoleDBBackendImpl) Get(ctx context.Context, publicID string) (pkg.Role, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the RoleDBBackend.Update() method for RoleDBBackendImpl.
func (impl RoleDBBackendImpl) Update(ctx context.Context, publicID string, elem pkg.Role) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the RoleDBBackend.GetAllByOrder() method for RoleDBBackendImpl.
func (impl RoleDBBackendImpl) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Role, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the RoleDBBackend.GetByField() method for RoleDBBackendImpl.
func (impl RoleDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (pkg.Role, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the RoleDBBackend.GetAll() method for RoleDBBackendImpl.
func (impl RoleDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Role, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
