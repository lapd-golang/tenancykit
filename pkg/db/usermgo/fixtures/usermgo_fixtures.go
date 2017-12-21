package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	UserJSON = `{


    "username":	"JohnnyReyes",

    "public_id":	"jbej6d5xlxz5tuggtse1rfqdxi1l98",

    "private_id":	"e6w5u90dw2mme4ww6dzj0ciqskdipa",

    "hash":	"au62swxuhfwug7is8b4m",

    "updated_at":	"2017-12-21 11:37:51.943651 +0000 UTC",

    "email":	"dolores_non@Skaboo.biz",

    "tenant_id":	"l42ejpnzy5ol758nv6o0",

    "two_factor_auth":	true,

    "created_at":	"2017-12-21 11:37:51.924575 +0000 UTC"

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
