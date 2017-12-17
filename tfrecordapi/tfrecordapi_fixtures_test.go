package tfrecordapi_test

import (
	"encoding/json"

	"github.com/gokit/tenancykit"
)

var tfrecordJSON = `{


    "key":	"",

    "domain":	"",

    "user_id":	"",

    "totp":	"",

    "public_id":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 16:45:51.497432 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:51.497441 +0000 UTC",

    "tenant_id":	""

}`

var tfrecordCreateJSON = `{


    "tenant":	nil,

    "user":	nil,

    "domain":	"",

    "max_length":	0

}`

var tfrecordUpdateJSON = `{


    "domain":	"",

    "tenant_id":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 16:45:51.498252 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:51.49826 +0000 UTC",

    "key":	"",

    "user_id":	"",

    "totp":	"",

    "public_id":	""

}`

// loadJSONFor returns a new instance of a tenancykit.TFRecord from the provide json content.
func loadJSONFor(content string) (tenancykit.TFRecord, error) {
	var elem tenancykit.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TFRecord{}, err
	}

	return elem, nil
}
