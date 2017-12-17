package types

import (
	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit"
)

// TwoFactorSessionBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type TwoFactorSession.
type TwoFactorSessionBackend interface {
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem tenancykit.TwoFactorSession) error
	Get(ctx context.Context, publicID string) (tenancykit.TwoFactorSession, error)
	Update(ctx context.Context, publicID string, elem tenancykit.TwoFactorSession) error
	GetAllByOrder(ctx context.Context, order, orderBy string) ([]tenancykit.TwoFactorSession, error)
	GetByField(ctx context.Context, key string, value interface{}) (tenancykit.TwoFactorSession, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.TwoFactorSession, int, error)
}
