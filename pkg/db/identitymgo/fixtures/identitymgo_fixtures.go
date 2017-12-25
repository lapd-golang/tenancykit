package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"

)


// json fixtures ...
var (
 IdentityJSON = `{


    "target_id":	"n86noim4baj6uocnckjf",

    "last_login":	1,

    "refresh_interval":	null,

    "public_id":	"y24nohhs4drwnn9smb613cp4pyoncs",

    "refresh_token":	"qjyvyr6qppfd6f7h3hd2",

    "expires":	10

}`
)

// LoadIdentityJSON returns a new instance of a userclaimjwt.Identity.
func LoadIdentityJSON(content string) (userclaimjwt.Identity, error) {
	var elem userclaimjwt.Identity

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return userclaimjwt.Identity{}, err
	}

	return elem, nil
}

