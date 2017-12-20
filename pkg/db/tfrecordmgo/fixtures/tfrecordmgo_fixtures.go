package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 TFRecordJSON = `{


    "domain":	"Gabcube.gov",

    "user_id":	"cc30dmqu2ki3hwx1rs83",

    "totp":	"idwx30q82zvg5zi9nhyp",

    "public_id":	"z8tiqc8wkdvh064qwehn6xybv7yann",

    "updated_at":	"2017-12-20 14:54:59.811171 +0000 UTC",

    "key":	"y4rmiq28vjkzzs86ddja",

    "tenant_id":	"gpz5mvegh4mw975u573e",

    "code_length":	1,

    "created_at":	"2017-12-20 14:54:59.811162 +0000 UTC"

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

