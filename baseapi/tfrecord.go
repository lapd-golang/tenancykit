package baseapi

import (
	"github.com/gokit/tenancykit/api/tfrecordapi"
	"github.com/gokit/tenancykit/backends"
	"github.com/gokit/tenancykit/db/types"
	"github.com/influx6/faux/metrics"
)

// TwoFactorAPI exposes the full wrapper over the http api with underline db
// for the database linked in.
type TwoFactorAPI struct {
	tfrecordapi.TFRecordHTTP
}

// NewTwoFactorAPI returns a new instance of TwoFactorAPI.
func NewTwoFactorAPI(m metrics.Metrics, db types.TFRecordDBBackend) TwoFactorAPI {
	return TwoFactorAPI{
		TFRecordHTTP: tfrecordapi.New(m, backends.TFBackend{
			TFRecordDBBackend: db,
		}),
	}
}
