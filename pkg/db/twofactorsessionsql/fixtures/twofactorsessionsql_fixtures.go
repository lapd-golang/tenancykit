package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"zchkj0a0kbkck4jfv8hz",

    "public_id":	"6yc7b33wpr9y2l75ygfzysadlf4758",

    "bool":	true

}`
)

// LoadTwoFactorSessionJSON returns a new instance of a pkg.TwoFactorSession.
func LoadTwoFactorSessionJSON(content string) (pkg.TwoFactorSession, error) {
	var elem pkg.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.TwoFactorSession{}, err
	}

	return elem, nil
}
