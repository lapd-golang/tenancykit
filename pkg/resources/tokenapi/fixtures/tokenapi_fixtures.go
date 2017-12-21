package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TokenJSON = `{


    "public_id":	"ymxzecah79njjvg4vp45llyapzjrij",

    "target_id":	"xvzpqb481u34ns2ik16j",

    "value":	"gjnsb24wupxaj7ef5lv3"

}`

	TokenCreateJSON = `{


    "value":	"1m6zb5tv4eax5ihr1y04",

    "public_id":	"yjbfca36db3b27dd85yju04adcqm5l",

    "target_id":	"mtxfsm7gj6m6giw366os"

}`

	TokenUpdateJSON = `{


    "value":	"bqleoh1qzspo3ord59oc",

    "public_id":	"2ort4l8sohn6felyx5xcx13axa671k",

    "target_id":	"ujplbanh17pywww16i4w"

}`
)

// LoadCreateJSON returns a new instance of a pkg.Token.
func LoadCreateJSON(content string) (pkg.Token, error) {
	var elem pkg.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Token{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.Token.
func LoadUpdateJSON(content string) (pkg.Token, error) {
	var elem pkg.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Token{}, err
	}

	return elem, nil
}

// LoadTokenJSON returns a new instance of a pkg.Token.
func LoadTokenJSON(content string) (pkg.Token, error) {
	var elem pkg.Token

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Token{}, err
	}

	return elem, nil
}
