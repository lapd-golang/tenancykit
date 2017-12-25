package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TFRecordJSON = `{


    "user_id":	"t316vkm2lvfbzf2xkbyc",

    "totp":	"zwzlwqjr0bxh2q3wzws6",

    "tenant_id":	"hj4uqwhdczzhl8ehsr5j9cl49l6r65",

    "code_length":	10,

    "created_at":	"2017-12-25 14:26:17.6567 +0000 UTC",

    "updated_at":	"2017-12-25 14:26:17.656709 +0000 UTC",

    "key":	"8ms00oemk6fe4disjwhn",

    "domain":	"LiveZ.edu"

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

