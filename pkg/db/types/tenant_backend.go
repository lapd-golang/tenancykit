package types

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
)

// TenantDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type Tenant.
// @implement
type TenantDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem pkg.Tenant) error
	Get(ctx context.Context, publicID string) (pkg.Tenant, error)
	Update(ctx context.Context, publicID string, elem pkg.Tenant) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Tenant, error)
	GetByField(ctx context.Context, key string, value interface{}) (pkg.Tenant, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Tenant, int, error)
}
