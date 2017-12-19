package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 UserSessionJSON = `{


    "user_id":	"d",

    "public_id":	"1vxmkl52k4ejbl1",

    "token":	"l",

    "expires":	"2017-12-19 14:01:01.569226 +0000 UTC"

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

