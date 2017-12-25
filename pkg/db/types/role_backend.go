package types

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
)

// RoleDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type Role.
// @implement
type RoleDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem pkg.Role) error
	Get(ctx context.Context, publicID string) (pkg.Role, error)
	Update(ctx context.Context, publicID string, elem pkg.Role) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Role, error)
	GetByField(ctx context.Context, key string, value interface{}) (pkg.Role, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Role, int, error)
}
