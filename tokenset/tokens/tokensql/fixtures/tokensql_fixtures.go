package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/tokenset/tokens"

)


// json fixtures ...
var (
 TokenJSON = `{


    "value":	"f35ho5fpx5e5m8mbs3zh",

    "public_id":	"40sxxusdlrw5hp70p1ucb4lpmbzybr",

    "target_id":	"v0yvvvexso2293sbg13d"

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

