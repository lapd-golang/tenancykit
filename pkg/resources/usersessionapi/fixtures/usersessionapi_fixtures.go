package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"eu42zd82od4fcyuif5hd",

    "public_id":	"xtfmbuc0lw1t66hris662elh452dy7",

    "token":	"11gwrtz3qy4hnqqk05bt",

    "expires":	"2017-12-21 15:13:27.714976 +0000 UTC"

}`

	UserSessionCreateJSON = `{


    "expiration":	null,

    "email":	"0Stephens@Jetwire.org",

    "password":	"w09fjgo2tr6tukqkeqtr"

}`

	UserSessionUpdateJSON = `{


    "public_id":	"lgd1o5cg3nhrpl58tlrl1917er1pab",

    "token":	"jzv04vd3xfp8qsmus12s",

    "expires":	"2017-12-21 15:13:27.715615 +0000 UTC",

    "user_id":	"hxq8qva6p3pyybexcpke"

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
