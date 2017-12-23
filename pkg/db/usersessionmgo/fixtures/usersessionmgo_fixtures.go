package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserSessionJSON = `{


    "user_id":	"tw6hzwlh9acyust8n53p",

    "public_id":	"1bz7pguhlh2kzizlwhw93p1hzg7txj",

    "token":	"snl7bqg5il9e30ctscd1",

    "expires":	"2017-12-23 12:42:40.425298 +0000 UTC"

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

