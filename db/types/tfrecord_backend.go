package types

import (
	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit"
)

// TFRecordBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type TFRecord.
type TFRecordBackend interface {
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem tenancykit.TFRecord) error
	Get(ctx context.Context, publicID string) (tenancykit.TFRecord, error)
	Update(ctx context.Context, publicID string, elem tenancykit.TFRecord) error
	GetAllByOrder(ctx context.Context, order, orderBy string) ([]tenancykit.TFRecord, error)
	GetByField(ctx context.Context, key string, value interface{}) (tenancykit.TFRecord, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.TFRecord, int, error)
}
