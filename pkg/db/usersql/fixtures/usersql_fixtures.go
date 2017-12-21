package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserJSON = `{


    "email":	"wDaniels@Topiclounge.gov",

    "public_id":	"2fn8zsgjmfi2loosrejyv861obmg16",

    "tenant_id":	"dug7o5smw8skun3is9ef",

    "private_id":	"84mg43epfhk4ra2m1rdx64qt9qfnti",

    "hash":	"d1cf670oere0yone7hkn",

    "created_at":	"2017-12-21 12:06:19.518979 +0000 UTC",

    "updated_at":	"2017-12-21 12:06:19.518988 +0000 UTC",

    "username":	"vMorrison",

    "two_factor_auth":	true

}`
)

// LoadUserJSON returns a new instance of a pkg.User.
func LoadUserJSON(content string) (pkg.User, error) {
	var elem pkg.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.User{}, err
	}

	return elem, nil
}
