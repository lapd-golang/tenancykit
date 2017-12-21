package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TFRecordJSON = `{


    "public_id":	"s4bgv1sstqid6sdm4369d3731k8s2e",

    "tenant_id":	"d1abgprwuymwuro0xl3l",

    "totp":	"iwsxwgn7gzwh7kxkzhk8",

    "domain":	"Abatz.info",

    "user_id":	"tiada7zdzr6im6q6p6a5",

    "code_length":	1,

    "created_at":	"2017-12-21 11:37:52.109066 +0000 UTC",

    "updated_at":	"2017-12-21 11:37:52.109075 +0000 UTC",

    "key":	"5r395v6pskn9wtaf0ws4"

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
