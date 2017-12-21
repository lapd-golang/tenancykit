package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TenantJSON = `{


    "name":	"Mrs. Ms. Miss Jane Owens",

    "email":	"AngelaBurke@Centimia.name",

    "public_id":	"jo26h9tixwz6ecjz2ztiyrqxx4g0jq",

    "created_at":	"2017-12-21 13:02:17.334862 +0000 UTC",

    "updated_at":	"2017-12-21 13:02:17.334869 +0000 UTC"

}`

	TenantCreateJSON = `{


    "name":	"Michael Wallace",

    "email":	"LillianWard@Realmix.edu"

}`

	TenantUpdateJSON = `{


    "email":	"KellyLee@Trunyx.biz",

    "public_id":	"aefc7hfwqvxwk3liqokkjboay3npjo",

    "created_at":	"2017-12-21 13:02:17.335315 +0000 UTC",

    "updated_at":	"2017-12-21 13:02:17.33532 +0000 UTC",

    "name":	"Martha Martin"

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
