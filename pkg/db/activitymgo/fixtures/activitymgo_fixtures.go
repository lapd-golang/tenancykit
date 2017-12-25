package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 ActivityJSON = `{


    "name":	"Jonathan Larson",

    "public_id":	"p7i9q7be3z8m6m7gp9oughcyiikwpt"

}`
)

// LoadActivityJSON returns a new instance of a pkg.Activity.
func LoadActivityJSON(content string) (pkg.Activity, error) {
	var elem pkg.Activity

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Activity{}, err
	}

	return elem, nil
}

