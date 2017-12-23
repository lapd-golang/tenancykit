package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserJSON = `{


    "username":	"8Diaz",

    "two_factor_auth":	true,

    "tenant_id":	"40s4yp7el6xz08uv9wuy",

    "private_id":	"3getznvv33bu0ezfyl7k5vd5cr2vlx",

    "hash":	"a0179i7bk1idsn8srubl",

    "created_at":	"2017-12-23 12:42:40.497337 +0000 UTC",

    "updated_at":	"2017-12-23 12:42:40.497346 +0000 UTC",

    "email":	"uJordan@Rhynyx.net",

    "public_id":	"832xt4fzxrlxjalsu7ryb3s8vfleip"

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

