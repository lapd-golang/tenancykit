package usermgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var userJSON = `{


    "two_factor_auth":	false,

    "created_at":	"2017-12-17 19:02:34.865271 +0000 UTC",

    "tenant_id":	"",

    "email":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "updated_at":	"2017-12-17 19:02:34.865282 +0000 UTC",

    "username":	""

}`

var userCreateJSON = `{


    "username":	"",

    "public_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-17 19:02:34.865951 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:34.865958 +0000 UTC",

    "email":	"",

    "tenant_id":	"",

    "private_id":	""

}`

var userUpdateJSON = `{


    "tenant_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "username":	"",

    "email":	"",

    "created_at":	"2017-12-17 19:02:34.866644 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:34.86665 +0000 UTC",

    "public_id":	"",

    "private_id":	""

}`

// loadJSONFor returns a new instance of a tenancykit.User from the provide json content.
func loadJSONFor(content string) (tenancykit.User, error) {
	var elem tenancykit.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.User{}, err
	}

	return elem, nil
}
