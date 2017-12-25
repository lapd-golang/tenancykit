package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	ActivityJSON = `{


    "name":	"Fred Vasquez",

    "public_id":	"38408mxsch2h4jv19z2461ndpezi6a",

    "created_at":	"2017-12-25 16:32:45.965233 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.96524 +0000 UTC"

}`

	ActivityCreateJSON = `{


    "name":	"Patricia Kim",

    "public_id":	"f00p3ndzag650nodaapcr7wznda0n2",

    "created_at":	"2017-12-25 16:32:45.965521 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.965526 +0000 UTC"

}`

	ActivityUpdateJSON = `{


    "name":	"Annie Olson",

    "public_id":	"mbkri28apje7f46ybtf7j3ub8pjvb2",

    "created_at":	"2017-12-25 16:32:45.965765 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.965769 +0000 UTC"

}`
)

// LoadCreateJSON returns a new instance of a pkg.Activity.
func LoadCreateJSON(content string) (pkg.Activity, error) {
	var elem pkg.Activity

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Activity{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.Activity.
func LoadUpdateJSON(content string) (pkg.Activity, error) {
	var elem pkg.Activity

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Activity{}, err
	}

	return elem, nil
}

// LoadActivityJSON returns a new instance of a pkg.Activity.
func LoadActivityJSON(content string) (pkg.Activity, error) {
	var elem pkg.Activity

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Activity{}, err
	}

	return elem, nil
}
