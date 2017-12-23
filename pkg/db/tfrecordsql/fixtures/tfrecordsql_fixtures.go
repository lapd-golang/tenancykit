package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TFRecordJSON = `{


    "key":	"sgopp73sywduk590lj50",

    "domain":	"Blogtag.gov",

    "totp":	"f96vn2ob4729quo753ij",

    "code_length":	10,

    "created_at":	"2017-12-23 12:42:39.455123 +0000 UTC",

    "user_id":	"eagl2onrnl0shctilauq",

    "public_id":	"72stxs8lsqsdpx8zif5f9bphbrtgj8",

    "tenant_id":	"6xe2bs4m5ulg2gb9vil7",

    "updated_at":	"2017-12-23 12:42:39.455131 +0000 UTC"

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

