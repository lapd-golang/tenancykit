package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserSessionJSON = `{


    "token":	"aykbsy0s1mc7nh8h9hg3",

    "expires":	"2017-12-20 14:54:59.646817 +0000 UTC",

    "user_id":	"5wz4btzxn350ykolacyp",

    "public_id":	"ru3ncd06g0a3v3jjo76g8afir4epur"

}`
)

// LoadUserSessionJSON returns a new instance of a pkg.UserSession.
func LoadUserSessionJSON(content string) (pkg.UserSession, error) {
	var elem pkg.UserSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.UserSession{}, err
	}

	return elem, nil
}

