package tfrecordsql_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var tfrecordJSON = `{


    "domain":	"",

    "totp":	"",

    "public_id":	"",

    "tenant_id":	"",

    "created_at":	"2017-12-17 19:02:46.677838 +0000 UTC",

    "key":	"",

    "code_length":	0,

    "updated_at":	"2017-12-17 19:02:46.677864 +0000 UTC",

    "user_id":	""

}`

var tfrecordCreateJSON = `{


    "user_id":	"",

    "created_at":	"2017-12-17 19:02:46.678528 +0000 UTC",

    "key":	"",

    "totp":	"",

    "public_id":	"",

    "tenant_id":	"",

    "code_length":	0,

    "updated_at":	"2017-12-17 19:02:46.678537 +0000 UTC",

    "domain":	""

}`

var tfrecordUpdateJSON = `{


    "created_at":	"2017-12-17 19:02:46.679119 +0000 UTC",

    "updated_at":	"2017-12-17 19:02:46.679127 +0000 UTC",

    "public_id":	"",

    "tenant_id":	"",

    "code_length":	0,

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

