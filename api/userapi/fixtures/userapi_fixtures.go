package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserJSON = `{


    "hash":	"vlfhvc9opr8xsdairh6t",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 12:28:53.779193 +0000 UTC",

    "username":	"et",

    "public_id":	"nyatdn8sogvh5lo141ivihjbxd6w1t",

    "tenant_id":	"bpd5cfzwlq65n6l1o0ie",

    "private_id":	"m96hdcsbsczpa7bpi2m75szpbouj29",

    "email":	"CarolynFernandez@Skyble.org",

    "updated_at":	"2017-12-20 12:28:53.779211 +0000 UTC"

}`

	UserCreateJSON = `{


    "email":	"9Graham@Babbleset.edu",

    "username":	"qCollins",

    "password":	"k0b2mnx7enhuq2vygigj",

    "password_confirm":	"2lpojwe4a224aca7f5kf",

    "tenant_id":	"jjza0qvdc3kjcrwqbabj"

}`

	UserUpdateJSON = `{


    "password":	"b6fmtv3jiglst8r9y0ul",

    "password_confirm":	"v5tr32np8ubd7z20q6h6",

    "is_password_update":	true,

    "email":	"unde@Divape.net"

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
