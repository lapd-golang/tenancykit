package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserJSON = `{


    "updated_at":	"2017-12-20 14:55:00.86284 +0000 UTC",

    "username":	"DavidLopez",

    "public_id":	"ewecc624bna8fwp2hokpuq1gx89fb2",

    "private_id":	"3udjxgk9ecpgxcy25haloaacs7mwm2",

    "hash":	"zq7d6ljg2u4v996rf0my",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 14:55:00.862829 +0000 UTC",

    "email":	"AlanMendoza@Leenti.biz",

    "tenant_id":	"yfpa8ps5ou9erbkgaca4"

}`

	UserCreateJSON = `{


    "tenant_id":	"vh3fsmm9ski2yc9z24a0",

    "email":	"CarlosPatterson@Mycat.com",

    "username":	"jWashington",

    "password":	"nljmiqaam215zfxzwnc2",

    "password_confirm":	"a261wj2viub5js4wfcz5"

}`

	UserUpdateJSON = `{


    "password":	"55y9c4fc54ruu1813r35",

    "password_confirm":	"4166gjxbq39pf188ijwz",

    "is_password_update":	true,

    "email":	"kBurns@Roombo.net"

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
