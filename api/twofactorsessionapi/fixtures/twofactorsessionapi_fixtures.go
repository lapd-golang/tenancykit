package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "public_id":	"3jy0ee7zak74tou",

    "bool":	true,

    "user_id":	"p"

}`

	TwoFactorSessionCreateJSON = `{


    "user_id":	"k",

    "public_id":	"autlk2aq9zn3iye",

    "bool":	true

}`

	TwoFactorSessionUpdateJSON = `{


    "bool":	true,

    "user_id":	"x",

    "public_id":	"qmddh41z6n20j0d"

}`
)

// LoadCreateJSON returns a new instance of a tenancykit.TwoFactorSession.
func LoadCreateJSON(content string) (tenancykit.TwoFactorSession, error) {
	var elem tenancykit.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TwoFactorSession{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a tenancykit.TwoFactorSession.
func LoadUpdateJSON(content string) (tenancykit.TwoFactorSession, error) {
	var elem tenancykit.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TwoFactorSession{}, err
	}

	return elem, nil
}

// LoadTwoFactorSessionJSON returns a new instance of a tenancykit.TwoFactorSession.
func LoadTwoFactorSessionJSON(content string) (tenancykit.TwoFactorSession, error) {
	var elem tenancykit.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TwoFactorSession{}, err
	}

	return elem, nil
}
