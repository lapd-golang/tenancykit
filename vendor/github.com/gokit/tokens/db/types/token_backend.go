package types

import (
	"context"

	"github.com/gokit/tokens"
)

// TokenDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type Token.
// @implement
type TokenDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem tokens.Token) error
	Get(ctx context.Context, publicID string) (tokens.Token, error)
	Update(ctx context.Context, publicID string, elem tokens.Token) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]tokens.Token, error)
	GetByField(ctx context.Context, key string, value interface{}) (tokens.Token, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tokens.Token, int, error)
}
