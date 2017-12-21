package types

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
)

// UserDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type User.
// @implement
type UserDBBackend interface {
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem pkg.User) error
	Get(ctx context.Context, publicID string) (pkg.User, error)
	Update(ctx context.Context, publicID string, elem pkg.User) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.User, error)
	GetByField(ctx context.Context, key string, value interface{}) (pkg.User, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.User, int, error)
}
