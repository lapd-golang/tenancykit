package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TokenJSON = `{


    "value":	"ikd3amgzpqecm1zg1iy2",

    "public_id":	"bdll7a95mjlasy5dsuq1s355e1fs4u",

    "target_id":	"8esszpp387uupr0pae4m"

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

