package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "domain":	"Photojam.gov",

    "user_id":	"efgo7jxdjih7b6k58rmp",

    "public_id":	"d9cyz61edqnis5gnrse11ejwb77q7l",

    "tenant_id":	"6s1u4s8mgfwp0lo7epsc",

    "code_length":	13,

    "key":	"56rqtg4udylazt3y0mpd",

    "totp":	"et0th8nfbqgyp44bf2pk",

    "created_at":	"2017-12-21 15:13:27.948135 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:27.948146 +0000 UTC"

}`

	TFRecordCreateJSON = `{


    "max_length":	16,

    "tenant":	null,

    "user":	null,

    "domain":	"Leenti.org"

}`

	TFRecordUpdateJSON = `{


    "user_id":	"x3m581jzokf50kebyhu4",

    "tenant_id":	"evvhr4fko7j2t02zsmvc",

    "code_length":	11,

    "created_at":	"2017-12-21 15:13:27.948842 +0000 UTC",

    "key":	"pfz8m57s4rdyyt8oypwi",

    "domain":	"Aimbo.biz",

    "totp":	"3y0ck5g0aoz1vickh5yc",

    "public_id":	"52adkgture50zt7sbfo0jcxe0kunvj",

    "updated_at":	"2017-12-21 15:13:27.948847 +0000 UTC"

}`
)

// LoadCreateJSON returns a new instance of a pkg.NewTF.
func LoadCreateJSON(content string) (pkg.NewTF, error) {
	var elem pkg.NewTF

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.NewTF{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.TFRecord.
func LoadUpdateJSON(content string) (pkg.TFRecord, error) {
	var elem pkg.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.TFRecord{}, err
	}

	return elem, nil
}

// LoadTFRecordJSON returns a new instance of a pkg.TFRecord.
func LoadTFRecordJSON(content string) (pkg.TFRecord, error) {
	var elem pkg.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.TFRecord{}, err
	}

	return elem, nil
}
