package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"yzo28l51oeb38rsgrlv7",

    "public_id":	"0a7zmzgfutk5d3wcyitl6lshnco5hs",

    "token":	"o0iglz559jgppa7uzoeq",

    "expires":	"2017-12-23 12:42:41.08402 +0000 UTC"

}`

	UserSessionCreateJSON = `{


    "email":	"eos_sapiente@Avamba.org",

    "password":	"biet0n3e0s328j1elkh0",

    "expiration":	null

}`

	UserSessionUpdateJSON = `{


    "user_id":	"e6kutrngtzm075frcomq",

    "public_id":	"1pajin78ln87cebqldn1rohxwi9uav",

    "token":	"jot4rx51nfc9x6lg979g",

    "expires":	"2017-12-23 12:42:41.084635 +0000 UTC"

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
