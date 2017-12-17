package tenantapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

var tenantJSON = `{


    "name":	"",

    "email":	"",

    "serial":	nil,

    "public_id":	"",

    "created_at":	"2017-12-17 16:45:49.624871 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:49.624882 +0000 UTC"

}`

var tenantCreateJSON = `{


    "name":	"",

    "email":	""

}`

var tenantUpdateJSON = `{


    "public_id":	"",

    "created_at":	"2017-12-17 16:45:49.625603 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:49.625609 +0000 UTC",

    "name":	"",

    "email":	"",

    "serial":	nil

}`

// loadJSONFor returns a new instance of a tenancykit.Tenant from the provide json content.
func loadJSONFor(content string) (tenancykit.Tenant, error) {
	var elem tenancykit.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.Tenant{}, err
	}

	return elem, nil
}
