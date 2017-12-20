package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "user_id":	"u6iq43x7gkrd2gq8qm5c",

    "totp":	"imk5xhbrnd9om6irknl5",

    "public_id":	"1d4swfm10e8ra1w6f0yyvhvy7vxc0q",

    "created_at":	"2017-12-20 19:23:27.404064 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.40407 +0000 UTC",

    "domain":	"Feedfire.gov",

    "tenant_id":	"j1nji36nrrkbs6thhqwq",

    "code_length":	11,

    "key":	"pp73fy6zbrwm486uhno2"

}`

	TFRecordCreateJSON = `{


    "max_length":	17,

    "tenant":	null,

    "user":	null,

    "domain":	"Fatz.mil"

}`

	TFRecordUpdateJSON = `{


    "key":	"mf68lkihupldifde2zei",

    "tenant_id":	"e4b193tbeey48fnd2a5u",

    "code_length":	18,

    "created_at":	"2017-12-20 19:23:27.404696 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.404702 +0000 UTC",

    "domain":	"Youspan.edu",

    "user_id":	"1h69mfgo256ej6p0t7pm",

    "totp":	"9ntfnai6d5dnz83rz8ux",

    "public_id":	"mmzo4d5acmby0yxk7j9uj6r5zj0sk1"

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
