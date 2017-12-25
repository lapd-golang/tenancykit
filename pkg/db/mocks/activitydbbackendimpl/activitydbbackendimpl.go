package activitydbbackendimpl

import (
	context "context"

	pkg "github.com/gokit/tenancykit/pkg"
)

// ActivityDBBackendImpl defines a concrete struct which implements the methods for the
// ActivityDBBackend interface. All methods will panic, so add necessary internal logic.
type ActivityDBBackendImpl struct {
	CountFunc func(ctx context.Context) (int, error)

	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem pkg.Activity) error

	GetFunc func(ctx context.Context, publicID string) (pkg.Activity, error)

	UpdateFunc func(ctx context.Context, publicID string, elem pkg.Activity) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]pkg.Activity, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (pkg.Activity, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Activity, int, error)
}

// Count implements the ActivityDBBackend.Count() method for ActivityDBBackendImpl.
func (impl ActivityDBBackendImpl) Count(ctx context.Context) (int, error) {

	ret1, ret2 := impl.CountFunc(ctx)
	return ret1, ret2

}

// Delete implements the ActivityDBBackend.Delete() method for ActivityDBBackendImpl.
func (impl ActivityDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the ActivityDBBackend.Create() method for ActivityDBBackendImpl.
func (impl ActivityDBBackendImpl) Create(ctx context.Context, elem pkg.Activity) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the ActivityDBBackend.Get() method for ActivityDBBackendImpl.
func (impl ActivityDBBackendImpl) Get(ctx context.Context, publicID string) (pkg.Activity, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the ActivityDBBackend.Update() method for ActivityDBBackendImpl.
func (impl ActivityDBBackendImpl) Update(ctx context.Context, publicID string, elem pkg.Activity) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the ActivityDBBackend.GetAllByOrder() method for ActivityDBBackendImpl.
func (impl ActivityDBBackendImpl) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Activity, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the ActivityDBBackend.GetByField() method for ActivityDBBackendImpl.
func (impl ActivityDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (pkg.Activity, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the ActivityDBBackend.GetAll() method for ActivityDBBackendImpl.
func (impl ActivityDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Activity, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
