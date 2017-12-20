package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"aq6hvwb46hndd94xnpw6",

    "public_id":	"l4rg6e9rxi3vni843i9507hmqskoga",

    "token":	"h9ez2fr226bnxatlynxp",

    "expires":	"2017-12-20 19:23:27.473238 +0000 UTC"

}`

	UserSessionCreateJSON = `{


    "password":	"rhmrp3mu7zvib4lel215",

    "expiration":	null,

    "email":	"hBowman@Vitz.biz"

}`

	UserSessionUpdateJSON = `{


    "user_id":	"8bi0mwleuyqvl2kp3r1f",

    "public_id":	"twuy3dv33spvvdeqgyawlaeun28vhe",

    "token":	"qkop9n13lefezttpc7or",

    "expires":	"2017-12-20 19:23:27.47386 +0000 UTC"

}`
)

// LoadCreateJSON returns a new instance of a pkg.CreateUserSession.
func LoadCreateJSON(content string) (pkg.CreateUserSession, error) {
	var elem pkg.CreateUserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.CreateUserSession{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.UserSession.
func LoadUpdateJSON(content string) (pkg.UserSession, error) {
	var elem pkg.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.UserSession{}, err
	}

	return elem, nil
}

// LoadUserSessionJSON returns a new instance of a pkg.UserSession.
func LoadUserSessionJSON(content string) (pkg.UserSession, error) {
	var elem pkg.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.UserSession{}, err
	}

	return elem, nil
}
