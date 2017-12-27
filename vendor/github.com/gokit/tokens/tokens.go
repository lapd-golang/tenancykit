package tokens

import (
	"errors"

	"github.com/influx6/faux/reflection"

	uuid "github.com/satori/go.uuid"
)

var (
	mapper *reflection.StructMapper

	// ErrNotFound identifies giving item has an error not found. Useful has a generic not found error.
	ErrNotFound = errors.New("not found")
)

// Token defines a specific record with associated id and
// value.
// @sqlapi
// @mongoapi(ENVName => TENANCYKIT)
// @httpapi
type Token struct {
	Value    string `json:"value"`
	PublicID string `json:"public_id"`
	TargetID string `json:"target_id"`
}

// NewToken returns a Token instance with provided targetID and value.
func NewToken(targetID string, value string) Token {
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
}
