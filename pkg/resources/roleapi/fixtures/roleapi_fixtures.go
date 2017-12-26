package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	RoleJSON = `{


    "name":	"Evelyn Warren",

    "public_id":	"4l95rsvrvmscbo51hrveu4a8tymwtg",

    "activities":	null,

    "created_at":	"2017-12-26T18:57:09+01:00",

    "updated_at":	"2017-12-26T18:57:09+01:00"

}`

	RoleCreateJSON = `"wlfij3kv5z8oztcmryp5"`

	RoleUpdateJSON = `{


    "name":	"Jacqueline Powell",

    "public_id":	"7jn0kyv5b8f91sitq8owa1cjkeibih",

    "activities":	null,

    "created_at":	"2017-12-26T18:57:09+01:00",

    "updated_at":	"2017-12-26T18:57:09+01:00"

}`
)

// LoadCreateJSON returns a new instance of a pkg.RoleName.
func LoadCreateJSON(content string) (pkg.RoleName, error) {
	var elem pkg.RoleName

	if err := json.Unmarshal([]byte(content), &elem); err != nil {

		return (*(*pkg.RoleName)(nil)), err

	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.Role.
func LoadUpdateJSON(content string) (pkg.Role, error) {
	var elem pkg.Role

	if err := json.Unmarshal([]byte(content), &elem); err != nil {

		return pkg.Role{}, err

	}

	return elem, nil
}

// LoadRoleJSON returns a new instance of a pkg.Role.
func LoadRoleJSON(content string) (pkg.Role, error) {
	var elem pkg.Role

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Role{}, err
	}

	return elem, nil
}
