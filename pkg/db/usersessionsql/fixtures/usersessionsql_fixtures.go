package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"senwrxhx04fcu7mtww2p",

    "public_id":	"4s75g387mufe7f70cls53lsv9sycc6",

    "token":	"m3n12na0lx9s6m255rea",

    "expires":	"2017-12-21 12:06:19.378529 +0000 UTC"

}`
)

// LoadUserSessionJSON returns a new instance of a pkg.UserSession.
func LoadUserSessionJSON(content string) (pkg.UserSession, error) {
	var elem pkg.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.UserSession{}, err
	}

	return elem, nil
}
