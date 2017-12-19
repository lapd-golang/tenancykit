package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserJSON = `{


    "username":	"3",

    "email":	"JulieCarroll@Trilia.edu",

    "public_id":	"8ehezg11g5425go",

    "private_id":	"2hsxiytuocu6tl8",

    "created_at":	2017-12-19 13:33:46.508537 +0000 UTC,

    "tenant_id":	"c",

    "hash":	"w",

    "two_factor_auth":	true,

    "updated_at":	2017-12-19 13:33:46.508543 +0000 UTC

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
