package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "totp":	"q",

    "public_id":	"ask9w3l0imldnic",

    "updated_at":	2017-12-19 13:42:58.600691 +0000 UTC,

    "created_at":	2017-12-19 13:42:58.600683 +0000 UTC,

    "key":	"y",

    "domain":	"t",

    "user_id":	"i",

    "tenant_id":	"j",

    "code_length":	6334824724549167320

}`

	TFRecordCreateJSON = `{


    "user":	nil,

    "domain":	"6",

    "max_length":	605394647632969758,

    "tenant":	nil

}`

	TFRecordUpdateJSON = `{


    "key":	"j",

    "public_id":	"gjikw7xo55o5sw9",

    "tenant_id":	"h",

    "domain":	"x",

    "user_id":	"3",

    "totp":	"4",

    "code_length":	1443635317331776148,

    "created_at":	2017-12-19 13:42:58.601343 +0000 UTC,

    "updated_at":	2017-12-19 13:42:58.601347 +0000 UTC

}`
)

// LoadCreateJSON returns a new instance of a tenancykit.NewTF.
func LoadCreateJSON(content string) (tenancykit.NewTF, error) {
	var elem tenancykit.NewTF

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.NewTF{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a tenancykit.TFRecord.
func LoadUpdateJSON(content string) (tenancykit.TFRecord, error) {
	var elem tenancykit.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TFRecord{}, err
	}

	return elem, nil
}

// LoadTFRecordJSON returns a new instance of a tenancykit.TFRecord.
func LoadTFRecordJSON(content string) (tenancykit.TFRecord, error) {
	var elem tenancykit.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TFRecord{}, err
	}

	return elem, nil
}
