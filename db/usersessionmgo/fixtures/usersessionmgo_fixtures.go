package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserSessionJSON = `{


    "user_id":	"3",

    "public_id":	"y329apz04ts0awn",

    "token":	"s",

    "expires":	2017-12-19 13:33:45.732992 +0000 UTC

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
