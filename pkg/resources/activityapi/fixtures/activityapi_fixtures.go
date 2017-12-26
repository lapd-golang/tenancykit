package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	ActivityJSON = `{


    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "name":	"Mr. Dr. Jimmy Jenkins",

    "public_id":	"4gfm4i2fxb8chtrubvjsdqfrsr4j6b"

}`

	ActivityCreateJSON = `"ge1nxopbl4m63nojl3jb"`

	ActivityUpdateJSON = `{


    "name":	"Sean Franklin",

    "public_id":	"rvhl6af6roh3xzumkjeyeh9levpw1e",

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00"

}`
)

// LoadCreateJSON returns a new instance of a pkg.ActivityName.
func LoadCreateJSON(content string) (pkg.ActivityName, error) {
	var elem pkg.ActivityName

	if err := json.Unmarshal([]byte(content), &elem); err != nil {

		return (*(*pkg.ActivityName)(nil)), err

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
