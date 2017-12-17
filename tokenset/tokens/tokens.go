package tokens

import (
	"reflect"
	"time"

	"github.com/influx6/faux/reflection"
	uuid "github.com/satori/go.uuid"
)

const (
	// TimeLayout is the default layout used for mapping a time.Time object to
	// a serialized string.
	TimeLayout = "Mon Jan 2 2006 15:04:05 -0700 MST"
)

var (
	mapper          *reflection.StructMapper
	timeReflectType = reflect.TypeOf((*time.Time)(nil))
)

//go:generate mgokit generate
//go:generate sqlkit generate
//go:generate httpkit generate

// Token defines a specific record with associated id and
// value.
// @sqlapi
// @mongoapi
// @httpapi
type Token struct {
	Value    string `json:"value"`
	PublicID string `json:"public_id"`
	TargetID string `json:"target_id"`
}

// New returns a Token instance with provided targetID and value.
func New(targetID string, value string) Token {
	return Token{
		Value:    value,
		TargetID: targetID,
		PublicID: uuid.NewV4().String(),
	}
}

// Consume consumes data from map into instance fields.
func (t *Token) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", t, data)
}

// Fields returns a map containing the instance fields as key-value pairs.
func (t Token) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &t)
}

func init() {
	mapper = reflection.NewStructMapper()
	mapper.AddAdapter(timeReflectType, reflection.TimeMapper(TimeLayout))
	mapper.AddInverseAdapter(timeReflectType, reflection.TimeInverseMapper(TimeLayout))
}
