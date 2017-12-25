package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserSessionJSON = `{


    "user_id":	"sp2zikgj0g4wripiiltd",

    "public_id":	"nsiy4vblize2864mqciw6xq9acknr9",

    "token":	"r5dr5ye4zvk8w0rnizw2",

    "expires":	"2017-12-25 14:26:17.496349 +0000 UTC"

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

