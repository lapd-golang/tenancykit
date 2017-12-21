package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TokenJSON = `{


    "target_id":	"8qhlr7kx4cjzx958t24m",

    "value":	"965e2rx0juzdg348x6p4",

    "public_id":	"bmbp7plnjg846jjnpq9hbild6oulz6"

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
