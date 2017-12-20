package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"u8kya5g50adlhr9hugb1",

    "public_id":	"lwmx3z4mgv9scflofw1ulm8rmt6p3n",

    "bool":	true

}`

	TwoFactorSessionCreateJSON = `{


    "user_id":	"4d4etgs0vfidp5f2rwmd",

    "public_id":	"0xf9939exv8t1011ngmoy2h698cfsj",

    "bool":	true

}`

	TwoFactorSessionUpdateJSON = `{


    "bool":	true,

    "user_id":	"mu5e5ri8kiy700j803kd",

    "public_id":	"60e16ex7s6scbw402ckw79ce2btsrt"

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
