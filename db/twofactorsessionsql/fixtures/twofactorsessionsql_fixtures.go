package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 TwoFactorSessionJSON = `{


    "public_id":	"xhlxij84yq7znb1",

    "bool":	true,

    "user_id":	"r"

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

