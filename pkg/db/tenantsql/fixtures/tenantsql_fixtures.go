package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TenantJSON = `{


    "created_at":	"2017-12-23 12:42:39.092215 +0000 UTC",

    "updated_at":	"2017-12-23 12:42:39.092239 +0000 UTC",

    "name":	"Norma Greene",

    "email":	"iste_officiis_amet@Oyonder.info",

    "public_id":	"69wehc58iw63ck7mq46xnui7wl8n83"

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

