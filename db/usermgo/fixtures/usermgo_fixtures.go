package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	UserJSON = `{


    "username":	"g",

    "public_id":	"pxt9fzpsi96mvds",

    "private_id":	"x17n3smx7pcqu7g",

    "hash":	"u",

    "created_at":	2017-12-19 13:33:45.817234 +0000 UTC,

    "updated_at":	2017-12-19 13:33:45.817241 +0000 UTC,

    "email":	"5Scott@Meetz.org",

    "tenant_id":	"9",

    "two_factor_auth":	true

}`
)

// LoadUserJSON returns a new instance of a tenancykit.User.
func LoadUserJSON(content string) (tenancykit.User, error) {
	var elem tenancykit.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.User{}, err
	}

	return elem, nil
}
