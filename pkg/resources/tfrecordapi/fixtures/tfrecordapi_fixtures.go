package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "public_id":	"1c4lszq0gvb0oymtf9zd8z6urlronr",

    "tenant_id":	"gx6zaudtmm0kc6qlrn47",

    "created_at":	"2017-12-23 12:42:40.891807 +0000 UTC",

    "updated_at":	"2017-12-23 12:42:40.891816 +0000 UTC",

    "domain":	"Quimm.info",

    "user_id":	"rk8hx0z6ytoq4dwwbyc4",

    "totp":	"cmk9w71r9vljgen2ihcl",

    "code_length":	18,

    "key":	"xtbfdzhosy0wk7ru37hc"

}`

	TFRecordCreateJSON = `{


    "max_length":	8,

    "tenant":	null,

    "user":	null,

    "domain":	"Feednation.name"

}`

	TFRecordUpdateJSON = `{


    "user_id":	"7okn4uet1l8z980c9n9h",

    "totp":	"1ivdaocy9p4fufgwnk5u",

    "code_length":	16,

    "created_at":	"2017-12-23 12:42:40.892377 +0000 UTC",

    "updated_at":	"2017-12-23 12:42:40.892383 +0000 UTC",

    "key":	"zda5qluhmne6o7xvyxe2",

    "domain":	"Feedbug.gov",

    "public_id":	"al1p7eckanl6pgv7j2088764rfwz3l",

    "tenant_id":	"bn9ap7vniside6nox244"

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
