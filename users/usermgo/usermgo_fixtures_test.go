package usermgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit/users"

)

var userJSON = `{


    "created_at":	nil,

    "updated_at":	nil,

    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false

}`

var userCreateJSON = `{


    "two_factor_auth":	false,

    "created_at":	nil,

    "updated_at":	nil,

    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	""

}`

var userUpdateJSON = `{


    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	nil,

    "updated_at":	nil

}`

// loadJSONFor returns a new instance of a users.User from the provide json content.
func loadJSONFor(content string) (users.User, error) {
	var elem users.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return users.User{}, err
	}

	return elem, nil
}
