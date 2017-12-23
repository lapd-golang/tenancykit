package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"dg9of3npxwvscf8x6q2z",

    "public_id":	"r52qo1wkx2a51p13nvtlyl9u9sf1py",

    "bool":	true

}`

	TwoFactorSessionCreateJSON = `{


    "user_id":	"6axtnaxfnbdqkz5q3qmn",

    "public_id":	"qouvhm7lzf9xexj10v5mv9hior8k5e",

    "bool":	true

}`

	TwoFactorSessionUpdateJSON = `{


    "public_id":	"pjf7i63d7dmihsdvxhptb7t4zw496y",

    "bool":	true,

    "user_id":	"jdjtztvhr3pxzh68b2x4"

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
