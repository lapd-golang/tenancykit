package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "totp":	"f",

    "created_at":	2017-12-19 13:33:45.997857 +0000 UTC,

    "code_length":	6129484611666145821,

    "updated_at":	2017-12-19 13:33:45.997865 +0000 UTC,

    "key":	"f",

    "domain":	"k",

    "user_id":	"p",

    "public_id":	"gol43w3nclzgmnl",

    "tenant_id":	"u"

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
