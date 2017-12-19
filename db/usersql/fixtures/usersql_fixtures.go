package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 UserJSON = `{


    "private_id":	"5fma0tbskgglz3r",

    "hash":	"m",

    "created_at":	"2017-12-19 14:01:02.244888 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.244897 +0000 UTC",

    "username":	"h",

    "email":	"iBell@Skyble.mil",

    "public_id":	"anr24rfuce644ca",

    "tenant_id":	"y",

    "two_factor_auth":	true

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

