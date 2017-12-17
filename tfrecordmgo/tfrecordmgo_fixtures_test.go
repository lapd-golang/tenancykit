package tfrecordmgo_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var tfrecordJSON = `{


    "key":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 16:45:11.626025 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:11.626037 +0000 UTC",

    "domain":	"",

    "user_id":	"",

    "totp":	"",

    "tenant_id":	"",

    "code_length":	0

}`

var tfrecordCreateJSON = `{


    "key":	"",

    "domain":	"",

    "user_id":	"",

    "public_id":	"",

    "tenant_id":	"",

    "totp":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 16:45:11.628257 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:11.628269 +0000 UTC"

}`

var tfrecordUpdateJSON = `{


    "tenant_id":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 16:45:11.764203 +0000 UTC",

    "public_id":	"",

    "updated_at":	"2017-12-17 16:45:11.764213 +0000 UTC",

    "key":	"",

    "domain":	"",

    "user_id":	"",

    "totp":	""

}`

// loadJSONFor returns a new instance of a tenancykit.TFRecord from the provide json content.
func loadJSONFor(content string) (tenancykit.TFRecord, error) {
	var elem tenancykit.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TFRecord{}, err
	}

	return elem, nil
}
