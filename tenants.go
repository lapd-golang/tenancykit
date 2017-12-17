package tenancykit

import (
	"crypto/rand"
	"math/big"
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
	Serial   *big.Int  `json:"serial"`
	PublicID string    `json:"public_id"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

// New returns a new Tenant from provided NewTenant.
func New(nt NewTenant) (Tenant, error) {
	var t Tenant

	serialNumber := new(big.Int).Lsh(big.NewInt(1), 128)
	serial, err := rand.Int(rand.Reader, serialNumber)
	if err != nil {
		return t, err
	}

	t.Name = nt.Name
	t.Email = nt.Email
	t.Serial = serial
	t.PublicID = uuid.NewV4().String()
	t.Created = time.Now()
	t.Updated = t.Created

	return t, nil
}
