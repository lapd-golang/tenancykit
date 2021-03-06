package pkg

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

// CreateTenant embodies the expected data used to create
// a Tenant.
type CreateTenant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Validate returns error if any field does not match requirements.
func (cu CreateTenant) Validate() error {
	if cu.Name == "" {
		return errors.New("Name is required")
	}

	if cu.Email == "" {
		return errors.New("Email is required")
	}

	return nil
}

// ITenant defines a tenant copy version with security details removed.
type ITenant struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	PublicID string    `json:"public_id"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

// Tenant defines a structure to represent a giving tenant.
// @mongoapi(ENVName => TENANCYKIT)
// @httpapi(New => CreateTenant)
type Tenant struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	PublicID string    `json:"public_id"`
	SecretID string    `json:"secret_id"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

// NewTenant returns a new Tenant from provided NewTenant.
func NewTenant(nt CreateTenant) Tenant {
	var t Tenant
	t.Name = nt.Name
	t.Email = nt.Email
	t.Created = time.Now()
	t.Updated = t.Created
	t.PublicID = uuid.NewV4().String()
	t.SecretID = uuid.NewV4().String()

	return t
}

// ITenant returns a ITenant instance which contains public fields of
// a tenant record.
func (t Tenant) ITenant() ITenant {
	return ITenant{
		Name:     t.Name,
		Email:    t.Email,
		PublicID: t.PublicID,
		Created:  t.Created,
		Updated:  t.Updated,
	}
}

// Consume consumes data from map into instance fields.
func (t *Tenant) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", t, data)
}

// Fields returns a map containing the instance fields as key-value pairs.
func (t Tenant) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &t)
}
