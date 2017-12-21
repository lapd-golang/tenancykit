package types

import (
	"context"

	"github.com/gokit/tenancykit/pkg"
)

// TFRecordDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type TFRecord.
// @implement
type TFRecordDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem pkg.TFRecord) error
	Get(ctx context.Context, publicID string) (pkg.TFRecord, error)
	Update(ctx context.Context, publicID string, elem pkg.TFRecord) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.TFRecord, error)
	GetByField(ctx context.Context, key string, value interface{}) (pkg.TFRecord, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.TFRecord, int, error)
}
