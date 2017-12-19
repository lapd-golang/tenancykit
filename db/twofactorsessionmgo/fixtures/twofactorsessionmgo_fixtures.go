package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 TwoFactorSessionJSON = `{


    "user_id":	"4",

    "public_id":	"o7r9hjxplnzecyy",

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

