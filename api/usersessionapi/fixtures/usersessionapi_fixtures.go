package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"9",

    "public_id":	"r296d7pajtwrsjz",

    "token":	"i",

    "expires":	2017-12-19 13:42:59.076742 +0000 UTC

}`

	UserSessionCreateJSON = `{


    "email":	"lBrown@Agivu.gov",

    "password":	"o",

    "expiration":	nil

}`

	UserSessionUpdateJSON = `{


    "user_id":	"g",

    "public_id":	"u9rp1fz0rv9vvbe",

    "token":	"u",

    "expires":	2017-12-19 13:42:59.07727 +0000 UTC

}`
)

// LoadCreateJSON returns a new instance of a tenancykit.CreateUserSession.
func LoadCreateJSON(content string) (tenancykit.CreateUserSession, error) {
	var elem tenancykit.CreateUserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.CreateUserSession{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a tenancykit.UserSession.
func LoadUpdateJSON(content string) (tenancykit.UserSession, error) {
	var elem tenancykit.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.UserSession{}, err
	}

	return elem, nil
}

// LoadUserSessionJSON returns a new instance of a tenancykit.UserSession.
func LoadUserSessionJSON(content string) (tenancykit.UserSession, error) {
	var elem tenancykit.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.UserSession{}, err
	}

	return elem, nil
}
