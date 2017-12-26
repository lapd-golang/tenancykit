package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TenantJSON = `{


    "name":	"Stephen Schmidt",

    "email":	"WalterPhillips@Abatz.com",

    "public_id":	"mqxwizbhw7eqhfjve46y5y36jy69xk",

    "secret_id":	"twn84rgyxql1j4zw3afm",

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00"

}`

	TenantCreateJSON = `{


    "name":	"Betty Franklin",

    "email":	"laborum@Reallinks.org"

}`

	TenantUpdateJSON = `{


    "name":	"Sara Alexander I II III IV V MD DDS PhD DVM",

    "email":	"mHoward@Gigashots.org",

    "public_id":	"rez07fqktq2oxqzbw08yz1hb4fotop",

    "secret_id":	"jj6ag8gkgobrhjg2iefg",

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00"

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
