package twofactorsessionapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

var twofactorsessionJSON = `{


    "user_id":	"",

    "public_id":	"",

    "bool":	false

}`

var twofactorsessionCreateJSON = `{


    "user_id":	"",

    "public_id":	"",

    "bool":	false

}`

var twofactorsessionUpdateJSON = `{


    "user_id":	"",

    "public_id":	"",

    "bool":	false

}`

// loadJSONFor returns a new instance of a tenancykit.TwoFactorSession from the provide json content.
func loadJSONFor(content string) (tenancykit.TwoFactorSession, error) {
	var elem tenancykit.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TwoFactorSession{}, err
	}

	return elem, nil
}
