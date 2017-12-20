package types

import (
	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit"
)

// UserDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type User.
// @implement
type UserDBBackend interface {
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem tenancykit.User) error
	Get(ctx context.Context, publicID string) (tenancykit.User, error)
	Update(ctx context.Context, publicID string, elem tenancykit.User) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]tenancykit.User, error)
	GetByField(ctx context.Context, key string, value interface{}) (tenancykit.User, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.User, int, error)
}
