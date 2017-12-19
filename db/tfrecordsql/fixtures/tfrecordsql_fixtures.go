package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "updated_at":	2017-12-19 13:33:46.463268 +0000 UTC,

    "key":	"0",

    "domain":	"t",

    "totp":	"q",

    "public_id":	"rggz5h2fgkeml1y",

    "tenant_id":	"u",

    "user_id":	"3",

    "code_length":	8674665223082153551,

    "created_at":	2017-12-19 13:33:46.463261 +0000 UTC

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
