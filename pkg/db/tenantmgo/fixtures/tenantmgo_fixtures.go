package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TenantJSON = `{


    "email":	"JoePorter@Vimbo.name",

    "public_id":	"vc7czsx9i08tsbiiocjgjep42y680y",

    "created_at":	"2017-12-21 11:37:52.026295 +0000 UTC",

    "updated_at":	"2017-12-21 11:37:52.026307 +0000 UTC",

    "name":	"Clarence Lawrence"

}`
)

// LoadTenantJSON returns a new instance of a pkg.Tenant.
func LoadTenantJSON(content string) (pkg.Tenant, error) {
	var elem pkg.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Tenant{}, err
	}

	return elem, nil
}
