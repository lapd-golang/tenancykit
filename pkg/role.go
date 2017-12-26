package pkg

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// RoleName defines a giving role name used for creating an Role.
type RoleName string

// Role embodies data which describes permission based on roles able to perform giving
// activities.
// @sqlapi
// @mongoapi
// @httpapi(New => RoleName)
type Role struct {
	Name       string    `json:"name"`
	PublicID   string    `json:"public_id"`
	Activities []string  `json:"activities"`
	Created    time.Time `json:"created_at"`
	Updated    time.Time `json:"updated_at"`
}

// NewRole returns a giving Role with provided name and auto-assigned uuid.
func NewRole(name RoleName) Role {
	created := time.Now()
	return Role{
		Name:     string(name),
		PublicID: uuid.NewV4().String(),
		Created:  created,
		Updated:  created,
	}
}

// hasID returns true/false if giving id exists in role's Activities.
func (r Role) hasID(id string) bool {
	for _, id := range r.Activities {
		if id == id {
			return true
		}
	}
	return false
}

// CanPerformAny returns true for all provided activities can by performed by the role.
func (r Role) CanPerformAll(activities ...Activity) bool {
	for _, act := range activities {
		if !r.hasID(act.PublicID) {
			return false
		}
	}
	return true
}

// CanPerformAny returns true for any provided activities can by performed by the role.
func (r Role) CanPerformAny(activities ...Activity) bool {
	for _, act := range activities {
		if r.hasID(act.PublicID) {
			return true
		}
	}
	return false
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
