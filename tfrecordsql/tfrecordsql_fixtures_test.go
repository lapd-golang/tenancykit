package tfrecordsql_test

import (
     "encoding/json"


     "github.com/gokit/tenancykit"

)

var tfrecordJSON = `{


    "domain":	"",

    "public_id":	"",

    "code_length":	0,

    "updated_at":	"2017-12-17 16:45:20.536714 +0000 UTC",

    "key":	"",

    "user_id":	"",

    "totp":	"",

    "tenant_id":	"",

    "created_at":	"2017-12-17 16:45:20.536687 +0000 UTC"

}`

var tfrecordCreateJSON = `{


    "domain":	"",

    "user_id":	"",

    "totp":	"",

    "public_id":	"",

    "tenant_id":	"",

    "created_at":	"2017-12-17 16:45:20.53717 +0000 UTC",

    "key":	"",

    "code_length":	0,

    "updated_at":	"2017-12-17 16:45:20.537177 +0000 UTC"

}`

var tfrecordUpdateJSON = `{


    "public_id":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 16:45:20.537493 +0000 UTC",

    "updated_at":	"2017-12-17 16:45:20.537499 +0000 UTC",

    "domain":	"",

    "user_id":	"",

    "totp":	"",

    "tenant_id":	"",

    "key":	""

}`

// loadJSONFor returns a new instance of a tenancykit.TFRecord from the provide json content.
func loadJSONFor(content string) (tenancykit.TFRecord, error) {
	var elem tenancykit.TFRecord

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return tenancykit.TFRecord{}, err
	}

	return elem, nil
}

