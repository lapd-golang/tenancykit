package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TokenJSON = `{


    "value":	"v24pp5hx3zur9dqfl6t2",

    "public_id":	"22vi2nryekbni39i07ovjkyn4xbmxb",

    "target_id":	"eimdcx66zgf7wtlbzkim"

}`
)

// LoadTokenJSON returns a new instance of a pkg.Token.
func LoadTokenJSON(content string) (pkg.Token, error) {
	var elem pkg.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Token{}, err
	}

	return elem, nil
}

