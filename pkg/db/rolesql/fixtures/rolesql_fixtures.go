package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
	RoleJSON = `{


    "name":	"Catherine Wilson",

    "public_id":	"lrq14wvaxsp4dh14kc1u7fn4xlf670",

    "activities":	null,

    "created_at":	"2017-12-26 15:45:31.631558 +0000 UTC",

    "updated_at":	"2017-12-26 15:45:31.631588 +0000 UTC"

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

