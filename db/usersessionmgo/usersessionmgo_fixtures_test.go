package usersessionmgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var usersessionJSON = `{


    "user_id":	"",

    "public_id":	"",

    "token":	"",

    "expires":	"2017-12-17 19:02:33.841688 +0000 UTC"

}`

var usersessionCreateJSON = `{


    "expires":	"2017-12-17 19:02:33.842003 +0000 UTC",

    "user_id":	"",

    "public_id":	"",

    "token":	""

}`

var usersessionUpdateJSON = `{


    "user_id":	"",

    "public_id":	"",

    "token":	"",

    "expires":	"2017-12-17 19:02:33.842305 +0000 UTC"

}`

// loadJSONFor returns a new instance of a tenancykit.UserSession from the provide json content.
func loadJSONFor(content string) (tenancykit.UserSession, error) {
	var elem tenancykit.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.UserSession{}, err
	}

	return elem, nil
}
