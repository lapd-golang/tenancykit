package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserJSON = `{


    "username":	"sed_accusamus_et",

    "private_id":	"g58mdfkvh171vxys47hjxjm88rp39m",

    "hash":	"ax7ojr81ha3s5av16ajf",

    "updated_at":	"2017-12-23 12:42:39.682299 +0000 UTC",

    "email":	"uGriffin@Yabox.net",

    "public_id":	"1qiganbhog8vv994qkbt0l0kgdqaez",

    "tenant_id":	"a0tx7ibjnr3n4obnjlk3",

    "two_factor_auth":	true,

    "created_at":	"2017-12-23 12:42:39.682291 +0000 UTC"

}`
)

// LoadUserJSON returns a new instance of a pkg.User.
func LoadUserJSON(content string) (pkg.User, error) {
	var elem pkg.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.User{}, err
	}

	return elem, nil
}

