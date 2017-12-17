package usersessionsql_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var usersessionJSON = `{


    "public_id":	"",

    "token":	"",

    "expires":	"2017-12-17 19:02:47.206652 +0000 UTC",

    "user_id":	""

}`

var usersessionCreateJSON = `{


    "user_id":	"",

    "public_id":	"",

    "token":	"",

    "expires":	"2017-12-17 19:02:47.207281 +0000 UTC"

}`

var usersessionUpdateJSON = `{


    "user_id":	"",

    "public_id":	"",

    "token":	"",

    "expires":	"2017-12-17 19:02:47.207772 +0000 UTC"

}`

// loadJSONFor returns a new instance of a tenancykit.UserSession from the provide json content.
func loadJSONFor(content string) (tenancykit.UserSession, error) {
	var elem tenancykit.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.UserSession{}, err
	}

	return elem, nil
}

