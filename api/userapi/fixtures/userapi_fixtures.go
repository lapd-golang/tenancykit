package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserJSON = `{


    "email":	"magni_neque_ullam@Dynava.mil",

    "private_id":	"338ud2dhfiycc77",

    "two_factor_auth":	true,

    "created_at":	2017-12-19 13:42:58.85745 +0000 UTC,

    "username":	"a",

    "public_id":	"4m5fym0vkznw6zo",

    "tenant_id":	"s",

    "hash":	"d",

    "updated_at":	2017-12-19 13:42:58.857457 +0000 UTC

}`

	UserCreateJSON = `{


    "tenant_id":	"a",

    "email":	"EvelynPorter@Skilith.biz",

    "username":	"r",

    "password":	"6",

    "password_confirm":	"h"

}`

	UserUpdateJSON = `{


    "password":	"n",

    "password_confirm":	"6",

    "is_password_update":	true,

    "email":	"BonnieOwens@Yotz.com"

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
