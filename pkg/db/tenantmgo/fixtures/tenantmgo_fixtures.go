package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TenantJSON = `{


    "name":	"Dorothy Young",

    "email":	"sint_quia_unde@Fatz.net",

    "public_id":	"nrwjzzt99vw68s5yx88zo5lxaf47x2",

    "created_at":	"2017-12-23 12:42:40.211882 +0000 UTC",

    "updated_at":	"2017-12-23 12:42:40.211902 +0000 UTC"

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

