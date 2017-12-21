package types

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
)

// UserSessionDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type UserSession.
// @implement
type UserSessionDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem pkg.UserSession) error
	Get(ctx context.Context, publicID string) (pkg.UserSession, error)
	Update(ctx context.Context, publicID string, elem pkg.UserSession) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.UserSession, error)
	GetByField(ctx context.Context, key string, value interface{}) (pkg.UserSession, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.UserSession, int, error)
}
