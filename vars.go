package tenancykit

import (
	"reflect"
	"time"

	"github.com/influx6/faux/reflection"
)

const (
	// UserHashComplexity is the value used by the hashing function for generating
	// user password hash.
	UserHashComplexity = 10

	// TimeLayout is the default layout used for mapping a time.Time object to
	// a serialized string.
	TimeLayout = "Mon Jan 2 2006 15:04:05 -0700 MST"
)

var (
	mapper          *reflection.StructMapper
	timeReflectType = reflect.TypeOf((*time.Time)(nil))
)

func init() {
	mapper = reflection.NewStructMapper()
	mapper.AddAdapter(timeReflectType, reflection.TimeMapper(TimeLayout))
	mapper.AddInverseAdapter(timeReflectType, reflection.TimeInverseMapper(TimeLayout))
}
