package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserJSON = `{


    "username":	"itaque",

    "email":	"WayneYoung@Nlounge.net",

    "public_id":	"lqu613ydeigcmmua2ckeui194ossn9",

    "tenant_id":	"oqstt9bln1h7w7rhm2zy",

    "user_roles":	null,

    "private_id":	"25wnj40hm3v62jegbk6poqabgy547j",

    "hash":	"8r18rdieis395cz25sn5",

    "two_factor_auth":	true,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00"

}`

	UserCreateJSON = `{


    "tenant_id":	"pcyqj60jf0zksdxl5m71",

    "email":	"ducimus_iure@Feedspan.net",

    "username":	"MargaretClark",

    "password":	"vzkv5bcvic22spls2ylw",

    "password_confirm":	"uisiqrui52xza8g10t38",

    "twofactor_auth":	true

}`

	UserUpdateJSON = `{


    "hash":	"af8iahbn48icuvm30ph3",

    "two_factor_auth":	true,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "email":	"vel_dolores@Devify.edu",

    "public_id":	"0hf9cvx8oqtq8evjecvzrfmmhdv1my",

    "tenant_id":	"64f510kphwx6t21617dl",

    "private_id":	"8botl1fnnmhnhb5p7hb642fsdf963h",

    "username":	"sit",

    "user_roles":	null,

    "updated_at":	"2017-12-26T18:57:10+01:00"

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

// LoadUpdateJSON returns a new instance of a pkg.User.
func LoadUpdateJSON(content string) (pkg.User, error) {
	var elem pkg.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {

		return pkg.User{}, err

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
