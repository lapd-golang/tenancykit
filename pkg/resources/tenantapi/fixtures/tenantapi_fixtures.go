package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TenantJSON = `{


    "name":	"Timothy Gilbert",

    "email":	"veritatis_consectetur_adipisci@Skinix.gov",

    "public_id":	"apn06wr9xmdesq9lx3jre4e6avmnx5",

    "created_at":	"2017-12-20 19:23:27.329125 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.329133 +0000 UTC"

}`

	TenantCreateJSON = `{


    "email":	"JenniferKnight@Yadel.info",

    "name":	"Sharon Gonzales I II III IV V MD DDS PhD DVM"

}`

	TenantUpdateJSON = `{


    "name":	"Evelyn Baker",

    "email":	"placeat_hic_in@Jetpulse.org",

    "public_id":	"6b5135ppx77f5zd8r6b901vnkfvdbq",

    "created_at":	"2017-12-20 19:23:27.329745 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.329753 +0000 UTC"

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
