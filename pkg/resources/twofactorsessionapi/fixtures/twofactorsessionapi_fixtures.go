package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"znece3qporg34d4r81ph",

    "public_id":	"3h6mqo26c77uy33itcrqy64j3murqg",

    "bool":	true

}`

	TwoFactorSessionCreateJSON = `{


    "user_id":	"jwfyrie26gq0ldm2gnxw",

    "public_id":	"ewdmivvyvcm2wh2w71qvo64cgg2vuw",

    "bool":	true

}`

	TwoFactorSessionUpdateJSON = `{


    "user_id":	"i08bdyf74tdqvzy0m1rs",

    "public_id":	"yvmct00yvxqzphrt9ytpdjk0wa6c35",

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
