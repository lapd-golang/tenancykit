package fixtures

import (
     "encoding/json"


     "github.com/gokit/tenancykit/pkg"

)


// json fixtures ...
var (
 UserJSON = `{


    "username":	"est_hic",

    "private_id":	"sjyemnr822t712ggdi8v0rvf0xbc7a",

    "updated_at":	"2017-12-20 14:55:00.218095 +0000 UTC",

    "email":	"RobertPalmer@Blogtag.name",

    "public_id":	"9iczvfoso3r56mspzqagzszt7lw4yp",

    "tenant_id":	"mysdrmd26b0ql947zcvt",

    "hash":	"kra5djgenb4wvcmpp6c3",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 14:55:00.218084 +0000 UTC"

}`
)

// LoadUserJSON returns a new instance of a pkg.User.
func LoadUserJSON(content string) (pkg.User, error) {
	var elem pkg.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.User{}, err
	}

	return elem, nil
}

