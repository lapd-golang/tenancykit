package types

import (
	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit"
)

// TenantDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type Tenant.
// @implement
type TenantDBBackend interface {
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem tenancykit.Tenant) error
	Get(ctx context.Context, publicID string) (tenancykit.Tenant, error)
	Update(ctx context.Context, publicID string, elem tenancykit.Tenant) error
	GetAllByOrder(ctx context.Context, order, orderBy string) ([]tenancykit.Tenant, error)
	GetByField(ctx context.Context, key string, value interface{}) (tenancykit.Tenant, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.Tenant, int, error)
}
