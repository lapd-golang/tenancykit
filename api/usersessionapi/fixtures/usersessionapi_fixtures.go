package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"i",

    "public_id":	"fejeae9cn7a86hz",

    "token":	"q",

    "expires":	"2017-12-19 14:01:02.884914 +0000 UTC"

}`

	UserSessionCreateJSON = `{


    "expiration":	nil,

    "email":	"fuga_eum_atque@Yakidoo.info",

    "password":	"z"

}`

	UserSessionUpdateJSON = `{


    "user_id":	"q",

    "public_id":	"acer551q4gpjb4x",

    "token":	"o",

    "expires":	"2017-12-19 14:01:02.885442 +0000 UTC"

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
