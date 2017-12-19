package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TenantJSON = `{


    "name":	"Randy Mills",

    "email":	"vMiller@Eamia.gov",

    "public_id":	"t5tynlz1v8ukn3w",

    "created_at":	2017-12-19 13:33:45.908212 +0000 UTC,

    "updated_at":	2017-12-19 13:33:45.908219 +0000 UTC

}`
)

// LoadTenantJSON returns a new instance of a tenancykit.Tenant.
func LoadTenantJSON(content string) (tenancykit.Tenant, error) {
	var elem tenancykit.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.Tenant{}, err
	}

	return elem, nil
}
