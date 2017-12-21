package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserJSON = `{


    "private_id":	"of4ocka8bp8cbh9xf2mwg1kip65g7w",

    "created_at":	"2017-12-21 13:02:17.813825 +0000 UTC",

    "updated_at":	"2017-12-21 13:02:17.813833 +0000 UTC",

    "username":	"JaneBennett",

    "email":	"RonaldKennedy@Jetwire.info",

    "public_id":	"o3l3k3udxi3adsl33lo51ownwos9a7",

    "tenant_id":	"oa1jpref1d1kv1jlzb85",

    "hash":	"jm6l68ip7nm5feqvykgz",

    "two_factor_auth":	true

}`

	UserCreateJSON = `{


    "tenant_id":	"1nm1awrrrch8rixruyxo",

    "email":	"8Alvarez@Vinder.mil",

    "username":	"et_ratione_ut",

    "password":	"he058p6xdrv28ad4r1ls",

    "password_confirm":	"a9fhmzkos0gjni7guzp4"

}`

	UserUpdateJSON = `{


    "email":	"fugiat_et@Wikibox.mil",

    "password":	"bst707dthxk7c0lvb3ea",

    "password_confirm":	"ulsxcgooelzjbt958auc"

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
