package tenantapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

var tenantJSON = `{


    "name":	"",

    "email":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:11:28.667576 +0000 UTC",

    "updated_at":	"2017-12-17 19:11:28.667593 +0000 UTC"

}`

var tenantCreateJSON = `{


    "name":	"",

    "email":	""

}`

var tenantUpdateJSON = `{


    "public_id":	"",

    "created_at":	"2017-12-17 19:11:28.669104 +0000 UTC",

    "updated_at":	"2017-12-17 19:11:28.669115 +0000 UTC",

    "name":	"",

    "email":	""

}`

// loadJSONFor returns a new instance of a tenancykit.Tenant from the provide json content.
func loadJSONFor(content string) (tenancykit.Tenant, error) {
	var elem tenancykit.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.Tenant{}, err
	}

	return elem, nil
}
