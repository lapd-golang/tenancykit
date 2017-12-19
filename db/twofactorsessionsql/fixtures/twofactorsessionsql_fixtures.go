package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TwoFactorSessionJSON = `{


    "user_id":	"u",

    "public_id":	"9debd261tphs1jn",

    "bool":	true

}`
)

// LoadTwoFactorSessionJSON returns a new instance of a tenancykit.TwoFactorSession.
func LoadTwoFactorSessionJSON(content string) (tenancykit.TwoFactorSession, error) {
	var elem tenancykit.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TwoFactorSession{}, err
	}

	return elem, nil
}
