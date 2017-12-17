package tenantsql_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var tenantJSON = `{


    "email":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:02:47.395247 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:47.395264 +0000 UTC",

    "name":	""

}`

var tenantCreateJSON = `{


    "name":	"",

    "email":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:02:47.395701 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:47.395714 +0000 UTC"

}`

var tenantUpdateJSON = `{


    "name":	"",

    "email":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:02:47.396076 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:47.396081 +0000 UTC"

}`

// loadJSONFor returns a new instance of a tenancykit.Tenant from the provide json content.
func loadJSONFor(content string) (tenancykit.Tenant, error) {
	var elem tenancykit.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.Tenant{}, err
	}

	return elem, nil
}

