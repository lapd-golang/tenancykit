package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	RoleJSON = `{


    "public_id":	"2q4jhe9tltwnqzsgcjvukgglaywwox",

    "activities":	null,

    "created_at":	"2017-12-25 16:32:46.047905 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.047916 +0000 UTC",

    "name":	"Bobby Young"

}`

	RoleCreateJSON = `{


    "name":	"Steve Edwards",

    "public_id":	"0hhd8w1lehhuf85iqs3pubp9nlc50o",

    "activities":	null,

    "created_at":	"2017-12-25 16:32:46.049433 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.049443 +0000 UTC"

}`

	RoleUpdateJSON = `{


    "updated_at":	"2017-12-25 16:32:46.049961 +0000 UTC",

    "name":	"Steve Willis",

    "public_id":	"xby41dpvdoagy89wkj5i37n4ctxj87",

    "activities":	null,

    "created_at":	"2017-12-25 16:32:46.049954 +0000 UTC"

}`
)

// LoadCreateJSON returns a new instance of a pkg.Role.
func LoadCreateJSON(content string) (pkg.Role, error) {
	var elem pkg.Role

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Role{}, err
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
