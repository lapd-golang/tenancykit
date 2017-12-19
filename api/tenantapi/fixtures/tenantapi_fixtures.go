package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TenantJSON = `{


    "name":	"Brenda Hill",

    "email":	"CarlRose@Wordtune.info",

    "public_id":	"x2odzoi97qviu7k",

    "created_at":	"2017-12-19 14:01:02.524405 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.524413 +0000 UTC"

}`

	TenantCreateJSON = `{


    "name":	"Jack Meyer Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"harum_expedita_eum@Brainsphere.biz"

}`

	TenantUpdateJSON = `{


    "name":	"Ryan Ferguson",

    "email":	"esse_veniam_et@Skajo.mil",

    "public_id":	"m20ihypo5r1d3iz",

    "created_at":	"2017-12-19 14:01:02.524885 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.524891 +0000 UTC"

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
