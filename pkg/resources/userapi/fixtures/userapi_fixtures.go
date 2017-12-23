package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserJSON = `{


    "created_at":	"2017-12-23 12:42:41.164561 +0000 UTC",

    "updated_at":	"2017-12-23 12:42:41.16457 +0000 UTC",

    "email":	"cMurray@Gabtune.biz",

    "public_id":	"v9zrmguw8j5nyhhwrf33w0shajh25i",

    "tenant_id":	"p8ty5mewtiqehog4d32o",

    "two_factor_auth":	true,

    "username":	"nLewis",

    "private_id":	"p6lw8bozql6e2v0hifpikid9hyjpf1",

    "hash":	"8p31bqr1z8u80kpfkhgi"

}`

	UserCreateJSON = `{


    "tenant_id":	"9d8vu5g5cmzm0w0txrgy",

    "email":	"itaque_quam@Flashpoint.mil",

    "username":	"itaque_sint",

    "password":	"qa8jv14s90109p8vn1pl",

    "password_confirm":	"pf82x82k88b80a55c96h",

    "twofactor_auth":	true

}`

	UserUpdateJSON = `{


    "email":	"PatrickOlson@Gevee.edu",

    "password":	"i9lfy6xewuko9kgx3own",

    "password_confirm":	"qdbyl780tt3nj02mvgc5"

}`
)

// LoadCreateJSON returns a new instance of a pkg.CreateUser.
func LoadCreateJSON(content string) (pkg.CreateUser, error) {
	var elem pkg.CreateUser

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.CreateUser{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.UpdateUser.
func LoadUpdateJSON(content string) (pkg.UpdateUser, error) {
	var elem pkg.UpdateUser

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.UpdateUser{}, err
	}

	return elem, nil
}

// LoadUserJSON returns a new instance of a pkg.User.
func LoadUserJSON(content string) (pkg.User, error) {
	var elem pkg.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.User{}, err
	}

	return elem, nil
}
