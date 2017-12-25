package pkg

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Role embodies data which describes permission based on roles able to perform giving
// activities.
// @mongoapi
// @httpapi
type Role struct {
	Name       string    `json:"name"`
	PublicID   string    `json:"public_id"`
	Activities []string  `json:"activities"`
	Created    time.Time `json:"created_at"`
	Updated    time.Time `json:"updated_at"`
}

// NewRole returns a giving Role with provided name and auto-assigned uuid.
func NewRole(name string) Role {
	created := time.Now()
	return Role{
		Name:     name,
		PublicID: uuid.NewV4().String(),
		Created:  created,
		Updated:  created,
	}
}

// AddActivity adds the Activity public id into the giving role activity list.
func (r *Role) AddActivity(activity Activity) {
	r.Activities = append(r.Activities, activity.PublicID)
}

// Consume consumes data from map into instance fields.
func (r *Role) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", r, data)
}

// Fields returns a map containing the instance fields as key-value pairs.
func (r Role) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &r)
}
