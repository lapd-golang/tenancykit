package baseapi

import (
	"github.com/gokit/tenancykit/api/tfrecordapi"
)

// TwoFactorAPI augments the existing tfrecordapi.TFRecordHTTP with other
// methods that expose more functionality.
type TwoFactorAPI struct {
	tfrecordapi.TFRecordHTTP
}
