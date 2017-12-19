package tokenapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit/tokenset/tokens"
)

var tokenJSON = `{


    "public_id":	"",

    "target_id":	"",

    "value":	""

}`

var tokenCreateJSON = `{


    "public_id":	"",

    "target_id":	"",

    "value":	""

}`

var tokenUpdateJSON = `{


    "public_id":	"",

    "target_id":	"",

    "value":	""

}`

// loadJSONFor returns a new instance of a tokens.Token from the provide json content.
func loadJSONFor(content string) (tokens.Token, error) {
	var elem tokens.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tokens.Token{}, err
	}

	return elem, nil
}
