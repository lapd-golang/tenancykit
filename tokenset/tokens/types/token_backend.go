package types

import (
	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit/tokenset/tokens"
)

// TokenDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type Token.
// @implement
type TokenDBBackend interface {
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem tokens.Token) error
	Get(ctx context.Context, publicID string) (tokens.Token, error)
	Update(ctx context.Context, publicID string, elem tokens.Token) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]tokens.Token, error)
	GetByField(ctx context.Context, key string, value interface{}) (tokens.Token, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tokens.Token, int, error)
}
