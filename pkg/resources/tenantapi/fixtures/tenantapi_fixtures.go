package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TenantJSON = `{


    "name":	"Jane Welch",

    "email":	"explicabo_maxime_omnis@Kaymbo.name",

    "public_id":	"lcgy59fudpd76m114ptjvjw0io3h73",

    "created_at":	"2017-12-20 14:55:00.641027 +0000 UTC",

    "updated_at":	"2017-12-20 14:55:00.641035 +0000 UTC"

}`

	TenantCreateJSON = `{


    "name":	"Mrs. Ms. Miss Rachel Perez",

    "email":	"LindaAllen@Dabvine.biz"

}`

	TenantUpdateJSON = `{


    "updated_at":	"2017-12-20 14:55:00.641517 +0000 UTC",

    "name":	"Jason Sanders Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"qPhillips@Gabtype.com",

    "public_id":	"glxuekt5rbm46ikt22n9hl6gy57ybi",

    "created_at":	"2017-12-20 14:55:00.641511 +0000 UTC"

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
