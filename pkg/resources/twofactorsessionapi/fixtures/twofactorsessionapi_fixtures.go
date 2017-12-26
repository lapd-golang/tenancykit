package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"pui5pvs3pimb0dqvde0d",

    "public_id":	"yxleagdsypxzx3b9tps7v11zradqgk",

    "bool":	true

}`

	TwoFactorSessionCreateJSON = `{


    "user_id":	"vjh718ecupx8muxufe15",

    "public_id":	"uhzn174xcypx8hnsg8vjt50mc8j08l",

    "bool":	true

}`

	TwoFactorSessionUpdateJSON = `{


    "user_id":	"caho84xlgyktqvubcdra",

    "public_id":	"8sak8etobeu6g69499jkjkhatrmwmb",

    "bool":	true

}`
)

// LoadCreateJSON returns a new instance of a pkg.TwoFactorSession.
func LoadCreateJSON(content string) (pkg.TwoFactorSession, error) {
	var elem pkg.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {

		return pkg.TwoFactorSession{}, err

	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.TwoFactorSession.
func LoadUpdateJSON(content string) (pkg.TwoFactorSession, error) {
	var elem pkg.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {

		return pkg.TwoFactorSession{}, err

	}

	return elem, nil
}

// LoadTwoFactorSessionJSON returns a new instance of a pkg.TwoFactorSession.
func LoadTwoFactorSessionJSON(content string) (pkg.TwoFactorSession, error) {
	var elem pkg.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.TwoFactorSession{}, err
	}

	return elem, nil
}
