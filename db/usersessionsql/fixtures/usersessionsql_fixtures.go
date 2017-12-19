package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"c",

    "public_id":	"en2vlnpimaxuzbv",

    "token":	"c",

    "expires":	2017-12-19 13:33:46.363574 +0000 UTC

}`
)

// LoadUserSessionJSON returns a new instance of a tenancykit.UserSession.
func LoadUserSessionJSON(content string) (tenancykit.UserSession, error) {
	var elem tenancykit.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.UserSession{}, err
	}

	return elem, nil
}
