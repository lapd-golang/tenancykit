package twofactorsessiondbbackendimpl

import (
	context "github.com/influx6/faux/context"

	tenancykit "github.com/gokit/tenancykit"
)

// TwoFactorSessionDBBackendImpl defines a concrete struct which implements the methods for the
// TwoFactorSessionDBBackend interface. All methods will panic, so add necessary internal logic.
type TwoFactorSessionDBBackendImpl struct {
	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem tenancykit.TwoFactorSession) error

	GetFunc func(ctx context.Context, publicID string) (tenancykit.TwoFactorSession, error)

	UpdateFunc func(ctx context.Context, publicID string, elem tenancykit.TwoFactorSession) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]tenancykit.TwoFactorSession, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (tenancykit.TwoFactorSession, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.TwoFactorSession, int, error)
}

// Delete implements the TwoFactorSessionDBBackend.Delete() method for TwoFactorSessionDBBackendImpl.
func (impl TwoFactorSessionDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the TwoFactorSessionDBBackend.Create() method for TwoFactorSessionDBBackendImpl.
func (impl TwoFactorSessionDBBackendImpl) Create(ctx context.Context, elem tenancykit.TwoFactorSession) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the TwoFactorSessionDBBackend.Get() method for TwoFactorSessionDBBackendImpl.
func (impl TwoFactorSessionDBBackendImpl) Get(ctx context.Context, publicID string) (tenancykit.TwoFactorSession, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the TwoFactorSessionDBBackend.Update() method for TwoFactorSessionDBBackendImpl.
func (impl TwoFactorSessionDBBackendImpl) Update(ctx context.Context, publicID string, elem tenancykit.TwoFactorSession) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the TwoFactorSessionDBBackend.GetAllByOrder() method for TwoFactorSessionDBBackendImpl.
func (impl TwoFactorSessionDBBackendImpl) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]tenancykit.TwoFactorSession, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the TwoFactorSessionDBBackend.GetByField() method for TwoFactorSessionDBBackendImpl.
func (impl TwoFactorSessionDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (tenancykit.TwoFactorSession, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the TwoFactorSessionDBBackend.GetAll() method for TwoFactorSessionDBBackendImpl.
func (impl TwoFactorSessionDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.TwoFactorSession, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
