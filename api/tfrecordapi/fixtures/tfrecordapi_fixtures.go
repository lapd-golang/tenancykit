package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "domain":	"a",

    "user_id":	"d",

    "totp":	"3",

    "created_at":	"2017-12-19 14:01:02.816138 +0000 UTC",

    "key":	"f",

    "public_id":	"h6fi44vd4gmbb0r",

    "tenant_id":	"7",

    "code_length":	2015796113853353331,

    "updated_at":	"2017-12-19 14:01:02.816145 +0000 UTC"

}`

	TFRecordCreateJSON = `{


    "max_length":	1874068156324778273,

    "tenant":	nil,

    "user":	nil,

    "domain":	"a"

}`

	TFRecordUpdateJSON = `{


    "updated_at":	"2017-12-19 14:01:02.816664 +0000 UTC",

    "user_id":	"u",

    "totp":	"n",

    "tenant_id":	"9",

    "code_length":	3328451335138149956,

    "key":	"f",

    "domain":	"4",

    "public_id":	"8dnqoukfl31812n",

    "created_at":	"2017-12-19 14:01:02.816659 +0000 UTC"

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
