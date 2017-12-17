package usersessionapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

var usersessionJSON = `{


    "user_id":	"",

    "public_id":	"",

    "token":	"",

    "expires":	"2017-12-17 19:11:25.287672 +0000 UTC"

}`

var usersessionCreateJSON = `{


    "user_id":	"",

    "public_id":	"",

    "token":	"",

    "expires":	"2017-12-17 19:11:25.288088 +0000 UTC"

}`

var usersessionUpdateJSON = `{


    "public_id":	"",

    "token":	"",

    "expires":	"2017-12-17 19:11:25.288352 +0000 UTC",

    "user_id":	""

}`

// loadJSONFor returns a new instance of a tenancykit.UserSession from the provide json content.
func loadJSONFor(content string) (tenancykit.UserSession, error) {
	var elem tenancykit.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.UserSession{}, err
	}

	return elem, nil
}
