package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserJSON = `{


    "created_at":	"2017-12-20 14:54:59.724698 +0000 UTC",

    "updated_at":	"2017-12-20 14:54:59.724708 +0000 UTC",

    "username":	"RoseGarrett",

    "email":	"facere@Yodel.name",

    "tenant_id":	"48mcqhw4cc0bnrrmtjxl",

    "two_factor_auth":	true,

    "public_id":	"nz7glu2oo8g0a4w2zcf6ov1cm35n4i",

    "private_id":	"pl6xdonhbu8if1cdem90vqed5hhtdp",

    "hash":	"62hss6a6q9wv8dx1u41c"

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

