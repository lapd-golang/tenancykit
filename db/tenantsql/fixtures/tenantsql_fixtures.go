package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TenantJSON = `{


    "name":	"Christine Barnes",

    "email":	"JessicaWelch@Podcat.biz",

    "public_id":	"t7oacha8t3eg70k",

    "created_at":	2017-12-19 13:33:46.418556 +0000 UTC,

    "updated_at":	2017-12-19 13:33:46.418563 +0000 UTC

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
