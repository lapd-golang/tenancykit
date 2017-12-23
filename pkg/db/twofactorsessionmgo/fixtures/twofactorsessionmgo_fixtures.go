package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TwoFactorSessionJSON = `{


    "bool":	true,

    "user_id":	"aix2fen9arfz9s8y4bf1",

    "public_id":	"buv3rjksbh528ni9y9vod8hin9d0o8"

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

