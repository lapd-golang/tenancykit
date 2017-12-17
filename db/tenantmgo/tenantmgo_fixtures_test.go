package tenantmgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var tenantJSON = `{


    "name":	"",

    "email":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:02:35.289237 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:35.289246 +0000 UTC"

}`

var tenantCreateJSON = `{


    "email":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:02:35.289658 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:35.289663 +0000 UTC",

    "name":	""

}`

var tenantUpdateJSON = `{


    "public_id":	"",

    "created_at":	"2017-12-17 19:02:35.28997 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:35.289975 +0000 UTC",

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
