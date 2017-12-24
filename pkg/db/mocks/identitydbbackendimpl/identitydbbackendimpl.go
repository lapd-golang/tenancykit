package identitydbbackendimpl

import (
	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"

	context "context"
)

// IdentityDBBackendImpl defines a concrete struct which implements the methods for the
// IdentityDBBackend interface. All methods will panic, so add necessary internal logic.
type IdentityDBBackendImpl struct {
	CountFunc func(ctx context.Context) (int, error)

	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem userclaimjwt.Identity) error

	GetFunc func(ctx context.Context, publicID string) (userclaimjwt.Identity, error)

	UpdateFunc func(ctx context.Context, publicID string, elem userclaimjwt.Identity) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]userclaimjwt.Identity, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (userclaimjwt.Identity, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]userclaimjwt.Identity, int, error)
}

// Count implements the IdentityDBBackend.Count() method for IdentityDBBackendImpl.
func (impl IdentityDBBackendImpl) Count(ctx context.Context) (int, error) {

	ret1, ret2 := impl.CountFunc(ctx)
	return ret1, ret2

}

// Delete implements the IdentityDBBackend.Delete() method for IdentityDBBackendImpl.
func (impl IdentityDBBackendImpl) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the IdentityDBBackend.Create() method for IdentityDBBackendImpl.
func (impl IdentityDBBackendImpl) Create(ctx context.Context, elem userclaimjwt.Identity) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the IdentityDBBackend.Get() method for IdentityDBBackendImpl.
func (impl IdentityDBBackendImpl) Get(ctx context.Context, publicID string) (userclaimjwt.Identity, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the IdentityDBBackend.Update() method for IdentityDBBackendImpl.
func (impl IdentityDBBackendImpl) Update(ctx context.Context, publicID string, elem userclaimjwt.Identity) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the IdentityDBBackend.GetAllByOrder() method for IdentityDBBackendImpl.
func (impl IdentityDBBackendImpl) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]userclaimjwt.Identity, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the IdentityDBBackend.GetByField() method for IdentityDBBackendImpl.
func (impl IdentityDBBackendImpl) GetByField(ctx context.Context, key string, value interface{}) (userclaimjwt.Identity, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the IdentityDBBackend.GetAll() method for IdentityDBBackendImpl.
func (impl IdentityDBBackendImpl) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]userclaimjwt.Identity, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}
