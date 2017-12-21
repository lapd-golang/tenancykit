package types

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
)

// TokenDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type Token.
// @implement
type TokenDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem pkg.Token) error
	Get(ctx context.Context, publicID string) (pkg.Token, error)
	Update(ctx context.Context, publicID string, elem pkg.Token) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Token, error)
	GetByField(ctx context.Context, key string, value interface{}) (pkg.Token, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Token, int, error)
}
