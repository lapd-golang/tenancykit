package pkg

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Activity embodies data which defines a giving activity of any sort.
// @mongoapi
// @httpapi
type Activity struct {
	Name     string    `json:"name"`
	PublicID string    `json:"public_id"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

// NewActivity returns a new Activity with giving name.
func NewActivity(name string) Activity {
	created := time.Now()
	return Activity{
		Name:     name,
		PublicID: uuid.NewV4().String(),
		Created:  created,
		Updated:  created,
	}
}

// Consume consumes data from map into instance fields.
func (a *Activity) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", a, data)
}

// Fields returns a map containing the instance fields as key-value pairs.
func (a Activity) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &a)
}
