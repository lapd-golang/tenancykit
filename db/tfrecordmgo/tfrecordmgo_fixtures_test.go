package tfrecordmgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var tfrecordJSON = `{


    "public_id":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 19:02:25.704121 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:25.741546 +0000 UTC",

    "key":	"",

    "domain":	"",

    "user_id":	"",

    "totp":	"",

    "tenant_id":	""

}`

var tfrecordCreateJSON = `{


    "tenant_id":	"",

    "code_length":	0,

    "domain":	"",

    "user_id":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:02:25.778131 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:25.778144 +0000 UTC",

    "key":	"",

    "totp":	""

}`

var tfrecordUpdateJSON = `{


    "tenant_id":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 19:02:25.778676 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:25.778683 +0000 UTC",

    "key":	"",

    "public_id":	"",

    "totp":	"",

    "domain":	"",

    "user_id":	""

}`

// loadJSONFor returns a new instance of a tenancykit.TFRecord from the provide json content.
func loadJSONFor(content string) (tenancykit.TFRecord, error) {
	var elem tenancykit.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TFRecord{}, err
	}

	return elem, nil
}
