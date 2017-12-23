package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TFRecordJSON = `{


    "updated_at":	"2017-12-23 12:42:40.280287 +0000 UTC",

    "domain":	"Skaboo.name",

    "user_id":	"c1bizewq19eakea7v29h",

    "totp":	"9wg74r2ghpyp7patgw6l",

    "public_id":	"brd7yf356m63wvh046do47vepicv3u",

    "tenant_id":	"8otb5e5l9e3ey8n0y0d2",

    "key":	"m5qxk8di54uvnq334522",

    "code_length":	10,

    "created_at":	"2017-12-23 12:42:40.280278 +0000 UTC"

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

