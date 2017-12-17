package tfrecordapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

var tfrecordJSON = `{


    "key":	"",

    "domain":	"",

    "totp":	"",

    "public_id":	"",

    "code_length":	0,

    "user_id":	"",

    "tenant_id":	"",

    "created_at":	"2017-12-17 19:11:22.125252 +0000 UTC",

    "updated_at":	"2017-12-17 19:11:22.125264 +0000 UTC"

}`

var tfrecordCreateJSON = `{


    "max_length":	0,

    "tenant":	nil,

    "user":	nil,

    "domain":	""

}`

var tfrecordUpdateJSON = `{


    "key":	"",

    "user_id":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 19:11:22.126158 +0000 UTC",

    "domain":	"",

    "totp":	"",

    "public_id":	"",

    "tenant_id":	"",

    "updated_at":	"2017-12-17 19:11:22.126163 +0000 UTC"

}`

// loadJSONFor returns a new instance of a tenancykit.TFRecord from the provide json content.
func loadJSONFor(content string) (tenancykit.TFRecord, error) {
	var elem tenancykit.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TFRecord{}, err
	}

	return elem, nil
}
