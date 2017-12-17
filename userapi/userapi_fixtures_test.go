package userapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

var userJSON = `{


    "tenant_id":	"",

    "updated_at":	"2017-12-17 16:45:53.114627 +0000 UTC",

    "email":	"",

    "public_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-17 16:45:53.114618 +0000 UTC",

    "username":	"",

    "private_id":	""

}`

var userCreateJSON = `{


    "password":	"",

    "password_confirm":	"",

    "tenant_id":	"",

    "email":	"",

    "username":	""

}`

var userUpdateJSON = `{


    "email":	"",

    "password":	"",

    "password_confirm":	""

}`

// loadJSONFor returns a new instance of a tenancykit.User from the provide json content.
func loadJSONFor(content string) (tenancykit.User, error) {
	var elem tenancykit.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.User{}, err
	}

	return elem, nil
}
