package types

import (
	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit/pkg"
)

// TwoFactorSessionDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type TwoFactorSession.
// @implement
type TwoFactorSessionDBBackend interface {
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem pkg.TwoFactorSession) error
	Get(ctx context.Context, publicID string) (pkg.TwoFactorSession, error)
	Update(ctx context.Context, publicID string, elem pkg.TwoFactorSession) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.TwoFactorSession, error)
	GetByField(ctx context.Context, key string, value interface{}) (pkg.TwoFactorSession, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.TwoFactorSession, int, error)
}
