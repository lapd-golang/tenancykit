package types

import (
	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit"
)

// UserSessionDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type UserSession.
// @implement
type UserSessionDBBackend interface {
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem tenancykit.UserSession) error
	Get(ctx context.Context, publicID string) (tenancykit.UserSession, error)
	Update(ctx context.Context, publicID string, elem tenancykit.UserSession) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]tenancykit.UserSession, error)
	GetByField(ctx context.Context, key string, value interface{}) (tenancykit.UserSession, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.UserSession, int, error)
}
