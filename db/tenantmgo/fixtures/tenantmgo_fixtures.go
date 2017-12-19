package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 TenantJSON = `{


    "name":	"Ruby Burke",

    "email":	"LoisRuiz@Oyonder.org",

    "public_id":	"9y763pf6wmx96qs",

    "created_at":	"2017-12-19 14:01:01.33187 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:01.331892 +0000 UTC"

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

