package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserSessionJSON = `{


    "user_id":	"n2c9d72mey3anx1rmza1",

    "public_id":	"x3wtno7b2hcu2si6cvjnj6dqdmjp8a",

    "token":	"ywj9ojiz7lu0hag5phae",

    "expires":	"2017-12-20 14:55:00.164439 +0000 UTC"

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

