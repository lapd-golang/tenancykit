package tokendbbackendimpl

import (
	context "context"

	pkg "github.com/gokit/tenancykit/pkg"
)

// TokenDBBackendImpl defines a concrete struct which implements the methods for the
// TokenDBBackend interface. All methods will panic, so add necessary internal logic.
type TokenDBBackendImpl struct {
	CountFunc func(ctx context.Context) (int, error)

	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem pkg.Token) error

	GetFunc func(ctx context.Context, publicID string) (pkg.Token, error)

	UpdateFunc func(ctx context.Context, publicID string, elem pkg.Token) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]pkg.Token, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (pkg.Token, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Token, int, error)
}

// Count implements the TokenDBBackend.Count() method for TokenDBBackendImpl.
func (impl TokenDBBackendImpl) Count(ctx context.Context) (int, error) {

	ret1, ret2 := impl.CountFunc(ctx)
	return ret1, ret2

}

// Delete implements the TokenDBBackend.Delete() method for TokenDBBackendImpl.
func (impl TokenDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the TokenDBBackend.Create() method for TokenDBBackendImpl.
func (impl TokenDBBackendImpl) Create(ctx context.Context, elem pkg.Token) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the TokenDBBackend.Get() method for TokenDBBackendImpl.
func (impl TokenDBBackendImpl) Get(ctx context.Context, publicID string) (pkg.Token, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the TokenDBBackend.Update() method for TokenDBBackendImpl.
func (impl TokenDBBackendImpl) Update(ctx context.Context, publicID string, elem pkg.Token) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the TokenDBBackend.GetAllByOrder() method for TokenDBBackendImpl.
func (impl TokenDBBackendImpl) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Token, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the TokenDBBackend.GetByField() method for TokenDBBackendImpl.
func (impl TokenDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (pkg.Token, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the TokenDBBackend.GetAll() method for TokenDBBackendImpl.
func (impl TokenDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Token, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
