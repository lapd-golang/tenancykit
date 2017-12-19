package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 TenantJSON = `{


    "name":	"Roy Ramos",

    "email":	"wMurphy@Edgetag.name",

    "public_id":	"6ven4ple62wdeej",

    "created_at":	"2017-12-19 14:01:01.959706 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:01.959726 +0000 UTC"

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

