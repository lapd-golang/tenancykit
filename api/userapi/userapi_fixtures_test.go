package userapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

var userJSON = `{


    "created_at":	"2017-12-17 19:11:27.031164 +0000 UTC",

    "updated_at":	"2017-12-17 19:11:27.031174 +0000 UTC",

    "username":	"",

    "email":	"",

    "public_id":	"",

    "tenant_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false

}`

var userCreateJSON = `{


    "tenant_id":	"",

    "email":	"",

    "username":	"",

    "password":	"",

    "password_confirm":	""

}`

var userUpdateJSON = `{


    "password_confirm":	"",

    "email":	"",

    "password":	""

}`

// loadJSONFor returns a new instance of a tenancykit.User from the provide json content.
func loadJSONFor(content string) (tenancykit.User, error) {
	var elem tenancykit.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.User{}, err
	}

	return elem, nil
}
