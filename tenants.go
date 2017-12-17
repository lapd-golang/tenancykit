package tenancykit

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// NewTenant embodies the expected data used to create
// a Tenant.
type NewTenant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Tenant defines a structure to represent a giving tenant.
// @mongoapi
// @sqlapi
// @httpapi(New => NewTenant)
type Tenant struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	PublicID string    `json:"public_id"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

// New returns a new Tenant from provided NewTenant.
func New(nt NewTenant) (Tenant, error) {
	var t Tenant

	t.Name = nt.Name
	t.Email = nt.Email
	t.PublicID = uuid.NewV4().String()
	t.Created = time.Now()
	t.Updated = t.Created

	return t, nil
}

// Consume consumes data from map into instance fields.
func (t *Tenant) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", t, data)
}

// Fields returns a map containing the instance fields as key-value pairs.
func (t Tenant) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &t)
}
