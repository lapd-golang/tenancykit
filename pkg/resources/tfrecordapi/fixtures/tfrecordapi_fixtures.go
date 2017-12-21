package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "code_length":	18,

    "domain":	"Linkbuzz.info",

    "user_id":	"8zbgwwehmo14na6hnq9r",

    "totp":	"rbqbxf4f3f6km3pjo2pn",

    "public_id":	"s31qwix0j2ih8ubzmkebsd7dp82y6g",

    "key":	"6ny9ljug0o09agfi6odq",

    "tenant_id":	"51kpmrrcaaw8byaw1tj0",

    "created_at":	"2017-12-21 13:02:17.608105 +0000 UTC",

    "updated_at":	"2017-12-21 13:02:17.608113 +0000 UTC"

}`

	TFRecordCreateJSON = `{


    "tenant":	null,

    "user":	null,

    "domain":	"Divanoodle.gov",

    "max_length":	8

}`

	TFRecordUpdateJSON = `{


    "key":	"gnc6iafh0j3b5hs70ly2",

    "totp":	"qpbzhpurv2v0vkrvl6ss",

    "code_length":	16,

    "created_at":	"2017-12-21 13:02:17.608865 +0000 UTC",

    "domain":	"Skinder.net",

    "user_id":	"fxnwlffime05eb24f8ky",

    "public_id":	"v1b99dpkaa60gfo113xrlx6n28xbgv",

    "tenant_id":	"by86sk9eocnu9bca2s6d",

    "updated_at":	"2017-12-21 13:02:17.60887 +0000 UTC"

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
