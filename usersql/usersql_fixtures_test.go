package usersql_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var userJSON = `{


    "hash":	"",

    "two_factor_auth":	false,

    "updated_at":	"2017-12-17 16:45:21.904285 +0000 UTC",

    "public_id":	"",

    "tenant_id":	"",

    "private_id":	"",

    "created_at":	"2017-12-17 16:45:21.904272 +0000 UTC",

    "username":	"",

    "email":	""

}`

var userCreateJSON = `{


    "updated_at":	"2017-12-17 16:45:21.905086 +0000 UTC",

    "username":	"",

    "email":	"",

    "tenant_id":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-17 16:45:21.905074 +0000 UTC",

    "public_id":	"",

    "private_id":	"",

    "hash":	""

}`

var userUpdateJSON = `{


    "username":	"",

    "email":	"",

    "tenant_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "public_id":	"",

    "created_at":	"2017-12-17 16:45:21.905859 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:21.90587 +0000 UTC"

}`

// loadJSONFor returns a new instance of a tenancykit.User from the provide json content.
func loadJSONFor(content string) (tenancykit.User, error) {
	var elem tenancykit.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.User{}, err
	}

	return elem, nil
}

