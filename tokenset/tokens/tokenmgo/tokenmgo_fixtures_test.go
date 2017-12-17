package tokenmgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit/tokenset/tokens"

)

var tokenJSON = `{


    "value":	"",

    "public_id":	"",

    "target_id":	""

}`

var tokenCreateJSON = `{


    "value":	"",

    "public_id":	"",

    "target_id":	""

}`

var tokenUpdateJSON = `{


    "value":	"",

    "public_id":	"",

    "target_id":	""

}`

// loadJSONFor returns a new instance of a tokens.Token from the provide json content.
func loadJSONFor(content string) (tokens.Token, error) {
	var elem tokens.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tokens.Token{}, err
	}

	return elem, nil
}
