package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "updated_at":	"2017-12-21 12:06:19.110872 +0000 UTC",

    "key":	"aez8aavruefhthngeb3m",

    "domain":	"LiveZ.net",

    "totp":	"aykzitjiezswfeslp9xk",

    "created_at":	"2017-12-21 12:06:19.110844 +0000 UTC",

    "user_id":	"6xt7vqiy96ge3wb8990m",

    "public_id":	"ulow7csjumctb4k0bl5sl1vbozpbvr",

    "tenant_id":	"wkz6whokvqm46g2hpwwy",

    "code_length":	10

}`
)

// LoadTFRecordJSON returns a new instance of a pkg.TFRecord.
func LoadTFRecordJSON(content string) (pkg.TFRecord, error) {
	var elem pkg.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.TFRecord{}, err
	}

	return elem, nil
}
