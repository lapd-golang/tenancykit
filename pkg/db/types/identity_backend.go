package types

import (
	"context"

	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"
)

// IdentityDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type Identity.
// @implement
type IdentityDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem userclaimjwt.Identity) error
	Get(ctx context.Context, publicID string) (userclaimjwt.Identity, error)
	Update(ctx context.Context, publicID string, elem userclaimjwt.Identity) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]userclaimjwt.Identity, error)
	GetByField(ctx context.Context, key string, value interface{}) (userclaimjwt.Identity, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]userclaimjwt.Identity, int, error)
}
