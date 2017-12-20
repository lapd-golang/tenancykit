package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"6723lh7pu04cqhgvd8b2",

    "public_id":	"s12auas7pkqvepk4a4iekhcrdc5nzs",

    "token":	"h7o8juca1wy95cmjcdl2",

    "expires":	"2017-12-20 14:55:01.084024 +0000 UTC"

}`

	UserSessionCreateJSON = `{


    "expiration":	nil,

    "email":	"CarlosHughes@Browsezoom.gov",

    "password":	"6q3bzytbvqqgw3yn2x3p"

}`

	UserSessionUpdateJSON = `{


    "user_id":	"kw6zkzem84qgj96e3xb3",

    "public_id":	"smxonj1xl6a59nfyvhw5afq584wwsl",

    "token":	"ur439n53jlhhor9rd9wb",

    "expires":	"2017-12-20 14:55:01.084486 +0000 UTC"

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
