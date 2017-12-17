package tenantsql_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var tenantJSON = `{


    "public_id":	"",

    "created_at":	"2017-12-17 16:45:22.574916 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:22.57493 +0000 UTC",

    "name":	"",

    "email":	"",

    "serial":	nil

}`

var tenantCreateJSON = `{


    "serial":	nil,

    "public_id":	"",

    "created_at":	"2017-12-17 16:45:22.575795 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:22.575803 +0000 UTC",

    "name":	"",

    "email":	""

}`

var tenantUpdateJSON = `{


    "updated_at":	"2017-12-17 16:45:22.57621 +0000 UTC",

    "name":	"",

    "email":	"",

    "serial":	nil,

    "public_id":	"",

    "created_at":	"2017-12-17 16:45:22.576204 +0000 UTC"

}`

// loadJSONFor returns a new instance of a tenancykit.Tenant from the provide json content.
func loadJSONFor(content string) (tenancykit.Tenant, error) {
	var elem tenancykit.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.Tenant{}, err
	}

	return elem, nil
}

