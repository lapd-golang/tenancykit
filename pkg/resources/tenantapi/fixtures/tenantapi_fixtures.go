package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TenantJSON = `{


    "name":	"Mrs. Ms. Miss Rebecca West",

    "email":	"9Willis@Vidoo.org",

    "public_id":	"mc33e328qr3jji5dplhenbzf4m9bb9",

    "created_at":	"2017-12-23 12:42:41.244379 +0000 UTC",

    "updated_at":	"2017-12-23 12:42:41.244386 +0000 UTC"

}`

	TenantCreateJSON = `{


    "name":	"Stephen Mccoy",

    "email":	"repellendus@Babbleblab.info"

}`

	TenantUpdateJSON = `{


    "email":	"DebraRivera@Cogibox.biz",

    "public_id":	"rha1uqhlrzmvrgahp8wm3xggfty13p",

    "created_at":	"2017-12-23 12:42:41.244872 +0000 UTC",

    "updated_at":	"2017-12-23 12:42:41.244881 +0000 UTC",

    "name":	"Bruce Daniels"

}`
)

// LoadCreateJSON returns a new instance of a pkg.CreateTenant.
func LoadCreateJSON(content string) (pkg.CreateTenant, error) {
	var elem pkg.CreateTenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.CreateTenant{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.Tenant.
func LoadUpdateJSON(content string) (pkg.Tenant, error) {
	var elem pkg.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Tenant{}, err
	}

	return elem, nil
}

// LoadTenantJSON returns a new instance of a pkg.Tenant.
func LoadTenantJSON(content string) (pkg.Tenant, error) {
	var elem pkg.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Tenant{}, err
	}

	return elem, nil
}
