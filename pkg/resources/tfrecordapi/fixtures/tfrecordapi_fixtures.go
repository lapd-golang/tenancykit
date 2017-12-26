package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "updated_at":	"2017-12-26T18:57:10+01:00",

    "key":	"nsd6lhcyvfy5yrol1ma3",

    "domain":	"Twitterwire.edu",

    "user_id":	"0fmttoo2e5z5wkgaeuuy",

    "totp":	"jl75lzow7qp39ccvpv84",

    "tenant_id":	"yl23qke276cqhzw93jxes89a8tsa49",

    "code_length":	13,

    "created_at":	"2017-12-26T18:57:10+01:00"

}`

	TFRecordCreateJSON = `{


    "max_length":	8,

    "user":	null,

    "domain":	"Ozu.name"

}`

	TFRecordUpdateJSON = `{


    "user_id":	"0ay9kuakbwij8vkihrhf",

    "totp":	"howjqyo4kx4ltvht2qr4",

    "tenant_id":	"1hluc42g0h3ximykp0sjumljepzvig",

    "code_length":	11,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "key":	"sw5uhnr0rff9408wc3lr",

    "domain":	"Roomm.name"

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
