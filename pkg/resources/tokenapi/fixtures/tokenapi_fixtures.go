package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TokenJSON = `{


    "value":	"nfrs6d79y4jjzg8qaay4",

    "public_id":	"hbb06xrr84pd9uhofdn20lke706sl2",

    "target_id":	"alw094r3ivl7gv2yct22"

}`

	TokenCreateJSON = `{


    "value":	"r8fsazn7hkjzuvtxgcq0",

    "public_id":	"fkdmmlxks13m9oxzh3g6c1ckzvfv9i",

    "target_id":	"cpc11jo9d27msz475k8d"

}`

	TokenUpdateJSON = `{


    "value":	"qchqk23c7s815536bves",

    "public_id":	"os4tcsrpxjcpcq8gnoz62ag50e0uo4",

    "target_id":	"h10tvs0ae3qafx5t54uc"

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
