package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TokenJSON = `{


    "value":	"ah89fk9lm73rbhffkw3z",

    "public_id":	"rtxfyj1cc8t44xbpzjrwym2mt9s1hl",

    "target_id":	"rts72z1s2xrevgfakwzb"

}`
)

// LoadTokenJSON returns a new instance of a pkg.Token.
func LoadTokenJSON(content string) (pkg.Token, error) {
	var elem pkg.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Token{}, err
	}

	return elem, nil
}
