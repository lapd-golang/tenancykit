package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/tokenset/tokens"

)


// json fixtures ...
var (
 TokenJSON = `{


    "value":	"z697t26e2uuzxdf6oh4e",

    "public_id":	"1hfjzth3c2aiam719xo1rku2xu7m8x",

    "target_id":	"hlaik66m7dhcms0bto4z"

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

