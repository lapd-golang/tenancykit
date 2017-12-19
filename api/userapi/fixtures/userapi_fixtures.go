package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserJSON = `{


    "email":	"sStephens@Trunyx.biz",

    "tenant_id":	"x",

    "private_id":	"qul4fd2dn6mucx2",

    "created_at":	"2017-12-19 14:01:02.742844 +0000 UTC",

    "username":	"y",

    "hash":	"i",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-19 14:01:02.742851 +0000 UTC",

    "public_id":	"nxoq9fp12v15twf"

}`

	UserCreateJSON = `{


    "password":	"z",

    "password_confirm":	"m",

    "tenant_id":	"d",

    "email":	"LillianBradley@Ainyx.name",

    "username":	"w"

}`

	UserUpdateJSON = `{


    "password":	"f",

    "password_confirm":	"v",

    "is_password_update":	true,

    "email":	"LindaLawson@Vinte.net"

}`
)

// LoadCreateJSON returns a new instance of a tenancykit.CreateUser.
func LoadCreateJSON(content string) (tenancykit.CreateUser, error) {
	var elem tenancykit.CreateUser

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.CreateUser{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a tenancykit.UpdateUser.
func LoadUpdateJSON(content string) (tenancykit.UpdateUser, error) {
	var elem tenancykit.UpdateUser

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.UpdateUser{}, err
	}

	return elem, nil
}

// LoadUserJSON returns a new instance of a tenancykit.User.
func LoadUserJSON(content string) (tenancykit.User, error) {
	var elem tenancykit.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.User{}, err
	}

	return elem, nil
}
