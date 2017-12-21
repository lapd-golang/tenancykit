package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"nno485yh1l9m5uj0mgqa",

    "public_id":	"d98cm1t8zmend1j0xakh3wzc8qes6o",

    "token":	"1c519tebg4val0nmqbw8",

    "expires":	"2017-12-21 11:37:52.184226 +0000 UTC"

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
