package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TenantJSON = `{


    "updated_at":	"2017-12-21 12:06:19.426188 +0000 UTC",

    "name":	"Dennis Hawkins",

    "email":	"quibusdam@Kazio.name",

    "public_id":	"f291cpq3l9k1155b4blmva9of3s1re",

    "created_at":	"2017-12-21 12:06:19.426179 +0000 UTC"

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
