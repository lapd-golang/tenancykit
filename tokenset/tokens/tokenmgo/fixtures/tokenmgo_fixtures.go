package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/tokenset/tokens"

)


// json fixtures ...
var (
 TokenJSON = `{


    "value":	"9aqhopzkmawcyacw80k2",

    "public_id":	"w98upc13pein1sa19q4k6up6n3flkm",

    "target_id":	"clpipv8jlyyqtj2oyi23"

}`
)

// LoadTokenJSON returns a new instance of a tokens.Token.
func LoadTokenJSON(content string) (tokens.Token, error) {
	var elem tokens.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tokens.Token{}, err
	}

	return elem, nil
}

