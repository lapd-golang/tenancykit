package tenantmgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var tenantJSON = `{


    "name":	"",

    "email":	"",

    "serial":	nil,

    "public_id":	"",

    "created_at":	"2017-12-17 16:45:10.985666 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:11.049923 +0000 UTC"

}`

var tenantCreateJSON = `{


    "name":	"",

    "email":	"",

    "serial":	nil,

    "public_id":	"",

    "created_at":	"2017-12-17 16:45:11.123914 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:11.123926 +0000 UTC"

}`

var tenantUpdateJSON = `{


    "name":	"",

    "email":	"",

    "serial":	nil,

    "public_id":	"",

    "created_at":	"2017-12-17 16:45:11.12463 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:11.124637 +0000 UTC"

}`

// loadJSONFor returns a new instance of a tenancykit.Tenant from the provide json content.
func loadJSONFor(content string) (tenancykit.Tenant, error) {
	var elem tenancykit.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.Tenant{}, err
	}

	return elem, nil
}
