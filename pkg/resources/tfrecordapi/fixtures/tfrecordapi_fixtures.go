package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "key":	"0ogwmjqwb64tke9sl33n",

    "domain":	"Fiveclub.com",

    "user_id":	"0qsy8nvwge06f6jw0vkr",

    "public_id":	"i2d3objoxpivskjzpe2s3lc1frtp7i",

    "tenant_id":	"08amn64b509ofk7pitl4",

    "code_length":	11,

    "created_at":	"2017-12-20 14:55:01.015307 +0000 UTC",

    "totp":	"s4ymf484eueyri3gwcha",

    "updated_at":	"2017-12-20 14:55:01.015318 +0000 UTC"

}`

	TFRecordCreateJSON = `{


    "max_length":	17,

    "tenant":	nil,

    "user":	nil,

    "domain":	"Tagcat.biz"

}`

	TFRecordUpdateJSON = `{


    "totp":	"ddyz8rxl8xlqjtbiznt8",

    "code_length":	18,

    "created_at":	"2017-12-20 14:55:01.017186 +0000 UTC",

    "updated_at":	"2017-12-20 14:55:01.017198 +0000 UTC",

    "domain":	"Dabshots.edu",

    "user_id":	"o1dd4khv2i5fm1ayjwt8",

    "public_id":	"4vxn576cthx4ncs1bhyvlcyt7is9lp",

    "tenant_id":	"cfhp7pwaqhqgqekf4w0c",

    "key":	"5hrvs68p0w12xhhqsdq5"

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
