package usersql_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit/users"

)

var userJSON = `{


    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	nil,

    "updated_at":	nil

}`

var userCreateJSON = `{


    "created_at":	nil,

    "updated_at":	nil,

    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false

}`

var userUpdateJSON = `{


    "created_at":	nil,

    "updated_at":	nil,

    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false

}`

// loadJSONFor returns a new instance of a users.User from the provide json content.
func loadJSONFor(content string) (users.User, error) {
	var elem users.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return users.User{}, err
	}

	return elem, nil
}

