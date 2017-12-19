package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 TFRecordJSON = `{


    "totp":	"n",

    "tenant_id":	"a",

    "code_length":	5577006791947779410,

    "created_at":	"2017-12-19 14:01:01.404278 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:01.40429 +0000 UTC",

    "key":	"b",

    "domain":	"2",

    "user_id":	"c",

    "public_id":	"93yp8q8glyo4jzi"

}`
)

// LoadTFRecordJSON returns a new instance of a tenancykit.TFRecord.
func LoadTFRecordJSON(content string) (tenancykit.TFRecord, error) {
	var elem tenancykit.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TFRecord{}, err
	}

	return elem, nil
}

