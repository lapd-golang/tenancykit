package tenantdbbackendimpl

import (
	context "github.com/influx6/faux/context"

	pkg "github.com/gokit/tenancykit/pkg"
)

// TenantDBBackendImpl defines a concrete struct which implements the methods for the
// TenantDBBackend interface. All methods will panic, so add necessary internal logic.
type TenantDBBackendImpl struct {
	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem pkg.Tenant) error

	GetFunc func(ctx context.Context, publicID string) (pkg.Tenant, error)

	UpdateFunc func(ctx context.Context, publicID string, elem pkg.Tenant) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]pkg.Tenant, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (pkg.Tenant, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Tenant, int, error)
}

// Delete implements the TenantDBBackend.Delete() method for TenantDBBackendImpl.
func (impl TenantDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the TenantDBBackend.Create() method for TenantDBBackendImpl.
func (impl TenantDBBackendImpl) Create(ctx context.Context, elem pkg.Tenant) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the TenantDBBackend.Get() method for TenantDBBackendImpl.
func (impl TenantDBBackendImpl) Get(ctx context.Context, publicID string) (pkg.Tenant, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the TenantDBBackend.Update() method for TenantDBBackendImpl.
func (impl TenantDBBackendImpl) Update(ctx context.Context, publicID string, elem pkg.Tenant) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the TenantDBBackend.GetAllByOrder() method for TenantDBBackendImpl.
func (impl TenantDBBackendImpl) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Tenant, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the TenantDBBackend.GetByField() method for TenantDBBackendImpl.
func (impl TenantDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (pkg.Tenant, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the TenantDBBackend.GetAll() method for TenantDBBackendImpl.
func (impl TenantDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Tenant, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
