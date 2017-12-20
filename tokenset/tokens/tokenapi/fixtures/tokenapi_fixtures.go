package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/tokenset/tokens"
)

// json fixtures ...
var (
	TokenJSON = `{


    "target_id":	"e08w7yfhno2troa2kixb",

    "value":	"r7gus9prgokrugd9a3wy",

    "public_id":	"s3mmmq1npzpkn0tiblwme12zk4s2in"

}`

	TokenCreateJSON = `{


    "target_id":	"evq736jg98toa0ye6gah",

    "value":	"hizl0hd0l2vmi6ca1anm",

    "public_id":	"11p866zpvamc3okdifelh0wvh2nua4"

}`

	TokenUpdateJSON = `{


    "public_id":	"efvbpkdnp558ydiaumx3tj9d2sh3hc",

    "target_id":	"dq2jhl0lfbzwkc5e7fee",

    "value":	"agvfbkhp27f7gmd8ayzh"

}`
)

// LoadCreateJSON returns a new instance of a tokens.Token.
func LoadCreateJSON(content string) (tokens.Token, error) {
	var elem tokens.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tokens.Token{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a tokens.Token.
func LoadUpdateJSON(content string) (tokens.Token, error) {
	var elem tokens.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tokens.Token{}, err
	}

	return elem, nil
}

// LoadTokenJSON returns a new instance of a tokens.Token.
func LoadTokenJSON(content string) (tokens.Token, error) {
	var elem tokens.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tokens.Token{}, err
	}

	return elem, nil
}
