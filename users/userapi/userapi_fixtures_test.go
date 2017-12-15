package userapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit/users"
)

var userJSON = `{


    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-15 21:20:08.367174 +0000 UTC",

    "updated_at":	"2017-12-15 21:20:08.367185 +0000 UTC",

    "username":	""

}`

var userCreateJSON = `{


    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-15 21:20:08.36757 +0000 UTC",

    "updated_at":	"2017-12-15 21:20:08.367578 +0000 UTC",

    "username":	""

}`

var userUpdateJSON = `{


    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-15 21:20:08.367923 +0000 UTC",

    "updated_at":	"2017-12-15 21:20:08.367929 +0000 UTC",

    "username":	"",

    "public_id":	"",

    "private_id":	""

}`

// loadJSONFor returns a new instance of a users.User from the provide json content.
func loadJSONFor(content string) (users.User, error) {
	var elem users.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return users.User{}, err
	}

	return elem, nil
}
