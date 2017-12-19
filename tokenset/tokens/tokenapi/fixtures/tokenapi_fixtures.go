package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/tokenset/tokens"
)

// json fixtures ...
var (
	TokenJSON = `{


    "value":	"dktbnwyvyhf54o9mo353",

    "public_id":	"d1qdmmpcoq7aat7qyijz4ethdwbtv6",

    "target_id":	"t8oae33fz3vsvdmpmbqq"

}`

	TokenCreateJSON = `{


    "value":	"6qn333jy225u9aqgke0t",

    "public_id":	"fgug0xbk57cx0p3cwulzkfn8pn2gj9",

    "target_id":	"0k6s7y59vcbcfa6utc76"

}`

	TokenUpdateJSON = `{


    "value":	"g3fcihccu8o6wb15t1do",

    "public_id":	"l4dvt04csnm3ief2puiema71jnqtc4",

    "target_id":	"jrkbowiol3ccdj9pcjsk"

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
