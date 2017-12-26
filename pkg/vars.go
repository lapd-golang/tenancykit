package pkg

import (
	"errors"
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
	TimeLayout = "2006-01-02T15:04:05Z07:00"
)

var (
	// ErrNotFound identifies giving item has an error not found. Useful has a generic not found error.
	ErrNotFound = errors.New("not found")

	mapper          *reflection.StructMapper
	timeReflectType = reflect.TypeOf((*time.Time)(nil))
)

func init() {
	mapper = reflection.NewStructMapper()
	mapper.AddAdapter(timeReflectType, reflection.TimeMapper(TimeLayout))
	mapper.AddInverseAdapter(timeReflectType, reflection.TimeInverseMapper(TimeLayout))
}
