package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TokenJSON = `{


    "public_id":	"uewbbd2vp8gt15lbureee8vojxplx6",

    "target_id":	"jb7rgju2xy6jodjg8q9d",

    "value":	"jp21ic71brcu7fzuew5d"

}`

	TokenCreateJSON = `{


    "value":	"akj31rtcjcqelhldgelf",

    "public_id":	"z3vx8nz44u64j23v6j6due01hsbehh",

    "target_id":	"61dg5hsbla8y9fuiuw7m"

}`

	TokenUpdateJSON = `{


    "public_id":	"6nbqlvoy6nhtqvim0thx7j0tsqoir1",

    "target_id":	"x7zth0uo8lru15vw5uso",

    "value":	"r01u3my79sgtnew5n6jw"

}`
)

// LoadCreateJSON returns a new instance of a pkg.Token.
func LoadCreateJSON(content string) (pkg.Token, error) {
	var elem pkg.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Token{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.Token.
func LoadUpdateJSON(content string) (pkg.Token, error) {
	var elem pkg.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Token{}, err
	}

	return elem, nil
}

// LoadTokenJSON returns a new instance of a pkg.Token.
func LoadTokenJSON(content string) (pkg.Token, error) {
	var elem pkg.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Token{}, err
	}

	return elem, nil
}
