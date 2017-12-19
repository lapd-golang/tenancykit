package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)


// json fixtures ...
var (
 TFRecordJSON = `{


    "key":	"2",

    "domain":	"a",

    "public_id":	"cbhcz9tw2v43a2j",

    "tenant_id":	"3",

    "user_id":	"l",

    "totp":	"t",

    "code_length":	5577006791947779410,

    "created_at":	"2017-12-19 14:01:02.102097 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.102105 +0000 UTC"

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

