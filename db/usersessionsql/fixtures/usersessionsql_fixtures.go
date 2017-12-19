package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 UserSessionJSON = `{


    "expires":	"2017-12-19 14:01:02.151075 +0000 UTC",

    "user_id":	"j",

    "public_id":	"xryy5n0j3cydnmk",

    "token":	"o"

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

