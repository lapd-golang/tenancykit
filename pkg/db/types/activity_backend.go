package types

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
)

// ActivityDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type Activity.
// @implement
type ActivityDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem pkg.Activity) error
	Get(ctx context.Context, publicID string) (pkg.Activity, error)
	Update(ctx context.Context, publicID string, elem pkg.Activity) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Activity, error)
	GetByField(ctx context.Context, key string, value interface{}) (pkg.Activity, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Activity, int, error)
}
