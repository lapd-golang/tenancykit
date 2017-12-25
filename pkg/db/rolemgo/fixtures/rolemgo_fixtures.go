package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 RoleJSON = `{


    "name":	"Louise Hamilton",

    "public_id":	"qkvr352yc2b6w5x60q5krrl9csvfxl",

    "activities":	null

}`
)

// LoadRoleJSON returns a new instance of a pkg.Role.
func LoadRoleJSON(content string) (pkg.Role, error) {
	var elem pkg.Role

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Role{}, err
	}

	return elem, nil
}

