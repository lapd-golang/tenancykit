package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserJSON = `{


    "username":	"cLewis",

    "tenant_id":	"cl0hc7bjq6j54uoz46ls",

    "private_id":	"5ggc9clurwx7c0kdcg81qae5o6ejfd",

    "hash":	"sc63j1w71z8zsprzxrhk",

    "created_at":	"2017-12-20 19:23:27.252689 +0000 UTC",

    "email":	"dolores_eum_aliquid@Linkbridge.net",

    "public_id":	"9f7ega7d9ag3dd7t8vhig27ew5n0hy",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-20 19:23:27.252699 +0000 UTC"

}`

	UserCreateJSON = `{


    "password_confirm":	"84afzu9ygh3h80rpr69w",

    "tenant_id":	"a156s71p62o5oveioqol",

    "email":	"PamelaMorales@Twitterbridge.biz",

    "username":	"mFranklin",

    "password":	"5cqzwbe6ybt5eugvd0ii"

}`

	UserUpdateJSON = `{


    "email":	"CraigRamos@Mybuzz.net",

    "password":	"78iv8wes28sgu83xkwdd",

    "password_confirm":	"ep4yjq55o9lic9dm39i4",

    "is_password_update":	true

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
