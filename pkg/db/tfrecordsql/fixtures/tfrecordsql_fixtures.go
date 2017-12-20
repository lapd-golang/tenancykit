package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TFRecordJSON = `{


    "domain":	"Yombu.org",

    "public_id":	"73hii1hmy9c7v9y62erl61p4jmdm1p",

    "tenant_id":	"9buu9go6rrgn1s8hdmpg",

    "code_length":	11,

    "key":	"qknad6d2t72lumrdx7me",

    "user_id":	"93s7vvbihpdr7nmscbww",

    "totp":	"k6ayc0ginljxjb65r8iq",

    "created_at":	"2017-12-20 14:55:00.309336 +0000 UTC",

    "updated_at":	"2017-12-20 14:55:00.309345 +0000 UTC"

}`
)

// LoadTFRecordJSON returns a new instance of a pkg.TFRecord.
func LoadTFRecordJSON(content string) (pkg.TFRecord, error) {
	var elem pkg.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.TFRecord{}, err
	}

	return elem, nil
}

