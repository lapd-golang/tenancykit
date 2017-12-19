package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"s",

    "public_id":	"arrryjeppnn3byg",

    "bool":	true

}`

	TwoFactorSessionCreateJSON = `{


    "user_id":	"q",

    "public_id":	"obsio8rliizs8fr",

    "bool":	true

}`

	TwoFactorSessionUpdateJSON = `{


    "bool":	true,

    "user_id":	"d",

    "public_id":	"20r5hxbx7lf8hw0"

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
