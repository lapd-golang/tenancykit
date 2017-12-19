package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TenantJSON = `{


    "updated_at":	2017-12-19 13:42:58.927803 +0000 UTC,

    "name":	"Wanda Jenkins",

    "email":	"4Elliott@Edgeclub.edu",

    "public_id":	"ash2alzvvrygcql",

    "created_at":	2017-12-19 13:42:58.927795 +0000 UTC

}`

	TenantCreateJSON = `{


    "name":	"Phillip Dean",

    "email":	"RichardCarroll@Brainlounge.edu"

}`

	TenantUpdateJSON = `{


    "name":	"Joyce Schmidt",

    "email":	"AshleyMeyer@Rhybox.edu",

    "public_id":	"ei12ftu91stcf28",

    "created_at":	2017-12-19 13:42:58.928399 +0000 UTC,

    "updated_at":	2017-12-19 13:42:58.928405 +0000 UTC

}`
)

// LoadCreateJSON returns a new instance of a tenancykit.CreateTenant.
func LoadCreateJSON(content string) (tenancykit.CreateTenant, error) {
	var elem tenancykit.CreateTenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.CreateTenant{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a tenancykit.Tenant.
func LoadUpdateJSON(content string) (tenancykit.Tenant, error) {
	var elem tenancykit.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.Tenant{}, err
	}

	return elem, nil
}

// LoadTenantJSON returns a new instance of a tenancykit.Tenant.
func LoadTenantJSON(content string) (tenancykit.Tenant, error) {
	var elem tenancykit.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.Tenant{}, err
	}

	return elem, nil
}
