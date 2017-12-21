package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"949z0qlive0tfux9ymxy",

    "public_id":	"mbkjboierea48s88s4l5omztgpnlb5",

    "bool":	true

}`

	TwoFactorSessionCreateJSON = `{


    "user_id":	"lsstbqu3aoqrfe5rknql",

    "public_id":	"gyzzhdgb9rxa5aqu2ggonmnv6q6qby",

    "bool":	true

}`

	TwoFactorSessionUpdateJSON = `{


    "user_id":	"1lsreltmcdqknw7d5mw2",

    "public_id":	"h942sf6cd3eb3z7jjzz2fed6m58k54",

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
