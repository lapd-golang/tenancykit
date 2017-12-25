package pkg

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ActivityName defines a giving activity name used for creating an activty.
type ActivityName string

// Activity embodies data which defines a giving activity of any sort.
// @mongoapi
// @httpapi(New => ActivityName)
type Activity struct {
	Name     string    `json:"name"`
	PublicID string    `json:"public_id"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

// NewActivity returns a new Activity with giving name.
func NewActivity(name ActivityName) Activity {
	created := time.Now()
	return Activity{
		Name:     string(name),
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
