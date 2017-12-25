package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserJSON = `{


    "public_id":	"flb4yf8mcqok6z7anpr25l3ykp5t0a",

    "tenant_id":	"v1fr7xdt7feg7ko3zjlu",

    "hash":	"m36tc6h5ssl27h7xngak",

    "updated_at":	"2017-12-25 14:26:17.843023 +0000 UTC",

    "email":	"illo@Eidel.org",

    "private_id":	"5qbe7c92bmtjrz73e4oa8rh9bw1a9a",

    "two_factor_auth":	true,

    "created_at":	"2017-12-25 14:26:17.843014 +0000 UTC",

    "username":	"MarilynHansen"

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

