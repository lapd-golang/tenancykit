package usersql_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var userJSON = `{


    "public_id":	"",

    "tenant_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "updated_at":	"2017-12-17 19:02:47.277622 +0000 UTC",

    "email":	"",

    "private_id":	"",

    "created_at":	"2017-12-17 19:02:47.277613 +0000 UTC",

    "username":	""

}`

var userCreateJSON = `{


    "email":	"",

    "public_id":	"",

    "tenant_id":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-17 19:02:47.278191 +0000 UTC",

    "username":	"",

    "private_id":	"",

    "hash":	"",

    "updated_at":	"2017-12-17 19:02:47.278196 +0000 UTC"

}`

var userUpdateJSON = `{


    "public_id":	"",

    "hash":	"",

    "email":	"",

    "tenant_id":	"",

    "private_id":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-17 19:02:47.278649 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:47.278655 +0000 UTC",

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

