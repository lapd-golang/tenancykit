package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TenantJSON = `{


    "name":	"Betty Phillips",

    "email":	"nihil@Pixoboo.mil",

    "public_id":	"wbfp109euoakftlzjw7fakyt05pstc",

    "secret_id":	"1uqylfmldn2trhl4btjy",

    "created_at":	"2017-12-25 14:26:17.571396 +0000 UTC",

    "updated_at":	"2017-12-25 14:26:17.571404 +0000 UTC"

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

