package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TwoFactorSessionJSON = `{


    "user_id":	"v7qo7380tlg7tgtdps64",

    "public_id":	"mg22tg8r9jpqaaf9oye9uu4ymncvwe",

    "bool":	true

}`
)

// LoadTwoFactorSessionJSON returns a new instance of a pkg.TwoFactorSession.
func LoadTwoFactorSessionJSON(content string) (pkg.TwoFactorSession, error) {
	var elem pkg.TwoFactorSession

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.TwoFactorSession{}, err
	}

	return elem, nil
}

