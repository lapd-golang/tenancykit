package pkg

import (
	uuid "github.com/satori/go.uuid"
)

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
