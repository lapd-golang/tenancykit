package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 UserJSON = `{


    "email":	"qui@Livefish.net",

    "public_id":	"vadjda3ajxtmbvk",

    "private_id":	"yqdnhu0dru7uv47",

    "two_factor_auth":	true,

    "created_at":	"2017-12-19 14:01:01.654282 +0000 UTC",

    "username":	"5",

    "hash":	"1",

    "updated_at":	"2017-12-19 14:01:01.654292 +0000 UTC",

    "tenant_id":	"7"

}`
)

// LoadUserJSON returns a new instance of a tenancykit.User.
func LoadUserJSON(content string) (tenancykit.User, error) {
	var elem tenancykit.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.User{}, err
	}

	return elem, nil
}

