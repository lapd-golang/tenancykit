package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserJSON = `{


    "email":	"RichardHarvey@Thoughtbridge.edu",

    "created_at":	"2017-12-21 15:13:27.80115 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:27.801159 +0000 UTC",

    "username":	"accusantium",

    "public_id":	"j4nid6z2gr4yo2meiyyxoocs3u4gnl",

    "tenant_id":	"9pwiulp7ztmuilq9az7j",

    "private_id":	"bv2pfku3jsighwzskzh6qsfwmdyt1x",

    "hash":	"k26be6zpkzaln7gmyzn3",

    "two_factor_auth":	true

}`

	UserCreateJSON = `{


    "tenant_id":	"uuidz60mzoacmohetfsk",

    "email":	"laudantium@Demimbu.biz",

    "username":	"RubyRay",

    "password":	"msszvfmugzqg15j82391",

    "password_confirm":	"idnxkzlp26q9p90ra2f2"

}`

	UserUpdateJSON = `{


    "email":	"LauraAustin@Realblab.biz",

    "password":	"84mvn85dcz1uvwzo481t",

    "password_confirm":	"yyqwki958qfufeqmdp0k"

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
