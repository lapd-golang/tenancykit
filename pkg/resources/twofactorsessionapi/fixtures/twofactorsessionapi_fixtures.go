package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"21lbnebcvgkegcjuyvbf",

    "public_id":	"av04zk76rmcydks6roqvqsany26uoq",

    "bool":	true

}`

	TwoFactorSessionCreateJSON = `{


    "user_id":	"9nhmgtqkmbvr3x6coqgt",

    "public_id":	"vx9fwrkuvqtgkilo7eizv6bu1x7obj",

    "bool":	true

}`

	TwoFactorSessionUpdateJSON = `{


    "user_id":	"mo99stxi5im06od40sdk",

    "public_id":	"lgqvr5rlql0uv6t1hfkg04bx4udkl2",

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
