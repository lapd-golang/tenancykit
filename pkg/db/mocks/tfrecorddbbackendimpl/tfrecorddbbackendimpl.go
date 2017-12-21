package tfrecorddbbackendimpl

import (
	context "context"

	pkg "github.com/gokit/tenancykit/pkg"
)

// TFRecordDBBackendImpl defines a concrete struct which implements the methods for the
// TFRecordDBBackend interface. All methods will panic, so add necessary internal logic.
type TFRecordDBBackendImpl struct {
	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem pkg.TFRecord) error

	GetFunc func(ctx context.Context, publicID string) (pkg.TFRecord, error)

	UpdateFunc func(ctx context.Context, publicID string, elem pkg.TFRecord) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]pkg.TFRecord, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (pkg.TFRecord, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.TFRecord, int, error)
}

// Delete implements the TFRecordDBBackend.Delete() method for TFRecordDBBackendImpl.
func (impl TFRecordDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the TFRecordDBBackend.Create() method for TFRecordDBBackendImpl.
func (impl TFRecordDBBackendImpl) Create(ctx context.Context, elem pkg.TFRecord) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the TFRecordDBBackend.Get() method for TFRecordDBBackendImpl.
func (impl TFRecordDBBackendImpl) Get(ctx context.Context, publicID string) (pkg.TFRecord, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the TFRecordDBBackend.Update() method for TFRecordDBBackendImpl.
func (impl TFRecordDBBackendImpl) Update(ctx context.Context, publicID string, elem pkg.TFRecord) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the TFRecordDBBackend.GetAllByOrder() method for TFRecordDBBackendImpl.
func (impl TFRecordDBBackendImpl) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.TFRecord, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the TFRecordDBBackend.GetByField() method for TFRecordDBBackendImpl.
func (impl TFRecordDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (pkg.TFRecord, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the TFRecordDBBackend.GetAll() method for TFRecordDBBackendImpl.
func (impl TFRecordDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.TFRecord, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
