package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TenantJSON = `{


    "name":	"Jean Rodriguez I II III IV V MD DDS PhD DVM",

    "email":	"JoanDavis@Buzzdog.com",

    "public_id":	"65cfu4uixht740hzqpuvphyglt7nig",

    "created_at":	"2017-12-20 14:54:59.41682 +0000 UTC",

    "updated_at":	"2017-12-20 14:54:59.416843 +0000 UTC"

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

