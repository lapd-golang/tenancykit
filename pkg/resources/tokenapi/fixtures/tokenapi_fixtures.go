package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TokenJSON = `{


    "value":	"zkzxwr028rpg69x87cxx",

    "public_id":	"rbrsjexibibvapk9godsb2kxl7dma8",

    "target_id":	"xfgdo405a9ykkxzw660d"

}`

	TokenCreateJSON = `{


    "value":	"5szjy5fpkb4zac57gjug",

    "public_id":	"8hwb8m30fndo45yfmudsplpb6qf7h0",

    "target_id":	"8yqpyike6th28mfjkr4r"

}`

	TokenUpdateJSON = `{


    "target_id":	"yxgfvw9iype00p8d5ruh",

    "value":	"u62ejwa2ap8dosjhszak",

    "public_id":	"02ec3gz00z53topkld1a8to2xt5swn"

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
