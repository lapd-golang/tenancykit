package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "token":	"1v1wlt9f1e6av4n2z262",

    "expires":	"2017-12-26T18:57:10+01:00",

    "user_id":	"l4fowopofpcas88dlqnf",

    "public_id":	"mw3sm52qxi3m2d23m4n6beoisxchmm"

}`

	UserSessionCreateJSON = `{


    "email":	"voluptatem@Dabshots.info",

    "password":	"x59i1v4heuztmbveae5m",

    "expiration":	null

}`

	UserSessionUpdateJSON = `{


    "token":	"l7d184dvs1jzevfe9qhv",

    "expires":	"2017-12-26T18:57:10+01:00",

    "user_id":	"ru12xfn4737onhpv3p7q",

    "public_id":	"jv9rjufbojcxcfgtyhzgzg4ld9duh7"

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
