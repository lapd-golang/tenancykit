package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TenantJSON = `{


    "updated_at":	"2017-12-20 14:55:00.042053 +0000 UTC",

    "name":	"Todd Schmidt",

    "email":	"6Alexander@Lajo.name",

    "public_id":	"vr2r0kmazn8yjqdzxyl43ijrp3ujqz",

    "created_at":	"2017-12-20 14:55:00.042033 +0000 UTC"

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

