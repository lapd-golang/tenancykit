package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/tokenset/tokens"

)


// json fixtures ...
var (
 TokenJSON = `{


    "value":	"890754pjn9j18vbdmtay",

    "public_id":	"wwmjttrlrxgg42lm05rvstjjkj3nh4",

    "target_id":	"ci8lvar0z06of7xkijou"

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

