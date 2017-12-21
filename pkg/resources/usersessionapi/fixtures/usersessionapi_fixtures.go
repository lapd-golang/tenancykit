package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"kddlpvr93j97iw85va39",

    "public_id":	"p0y4s38y1dv400yfr1giufnolz6tnq",

    "token":	"b9cjagocqh4i4p56gbqc",

    "expires":	"2017-12-21 13:02:17.67844 +0000 UTC"

}`

	UserSessionCreateJSON = `{


    "email":	"lLee@Podcat.edu",

    "password":	"th824z1rgz8843kz9xis",

    "expiration":	null

}`

	UserSessionUpdateJSON = `{


    "user_id":	"qcy98t0iyzovms8fqfzu",

    "public_id":	"y3kla23wj4ifqcczgjiqnk34k9nb7g",

    "token":	"xmu9vbjh148ulzsw4mr0",

    "expires":	"2017-12-21 13:02:17.679077 +0000 UTC"

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
