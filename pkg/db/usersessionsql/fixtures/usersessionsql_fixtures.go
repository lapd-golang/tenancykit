package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserSessionJSON = `{


    "token":	"c7dxmoef7oy5t7q331zk",

    "expires":	"2017-12-23 12:42:39.604943 +0000 UTC",

    "user_id":	"itwzj21bv9wtv09f7k0i",

    "public_id":	"1mhzdgnjam5kodbhmkio234bt3irjh"

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

