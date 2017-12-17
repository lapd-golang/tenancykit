package usermgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var userJSON = `{


    "public_id":	"",

    "tenant_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-17 16:45:13.217268 +0000 UTC",

    "username":	"",

    "email":	"",

    "updated_at":	"2017-12-17 16:45:13.217299 +0000 UTC"

}`

var userCreateJSON = `{


    "email":	"",

    "public_id":	"",

    "tenant_id":	"",

    "private_id":	"",

    "two_factor_auth":	false,

    "username":	"",

    "hash":	"",

    "created_at":	"2017-12-17 16:45:13.218323 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:13.218334 +0000 UTC"

}`

var userUpdateJSON = `{


    "public_id":	"",

    "tenant_id":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-17 16:45:13.218938 +0000 UTC",

    "email":	"",

    "private_id":	"",

    "hash":	"",

    "updated_at":	"2017-12-17 16:45:13.218946 +0000 UTC",

    "username":	""

}`

// loadJSONFor returns a new instance of a tenancykit.User from the provide json content.
func loadJSONFor(content string) (tenancykit.User, error) {
	var elem tenancykit.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.User{}, err
	}

	return elem, nil
}
