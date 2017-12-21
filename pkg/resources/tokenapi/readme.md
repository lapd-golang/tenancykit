Token HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/tokenapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/tokenapi)

Token HTTP API is a auto generated http api for the struct `Token`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.Token, error)
    Update(context.Context, string, pkg.Token) error
    GetAll(context.Context, string, string, int, int) ([]pkg.Token, int, error)
    Create(context.Context, pkg.Token) (pkg.Token, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Token/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the Token type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "value":	"mq9vuvsow8y8p5lw8m3s",

    "public_id":	"47u15d22jb2fhqa6pvediarvmiwcca",

    "target_id":	"lbqscx7av7m9w5bt8who"

}
```

- Expected Status Code

```
Failure: 500
Success: 201
```

- Expected Response Body

```http
    Content-Type: application/json
```

```json
{


    "value":	"v40jdpes2bpfmwg2k6dx",

    "public_id":	"6xar18fedquo3dxlr4i2dwck0xd2qn",

    "target_id":	"qzste18jhcjd2ictuxc0"

}
```

## GET /Token/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the Token type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Request Parameters

```
    :public_id
```

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```http
    Content-Type: application/json
```

```json
{


    "value":	"8d8hxnonxhgh6lof3e23",

    "public_id":	"5m4f08eirty93v7rdssucdig1wosck",

    "target_id":	"w5qdgallzsfiayvf8mmn"

}
```

## GET /Token/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the Token type from the HTTP API.

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```http
    Content-Type: application/json
```

```json
[{


    "value":	"f3lkwbngijvpcqgwby05",

    "public_id":	"0p8k6hvkqzflbso7qukrm6miyamr3i",

    "target_id":	"ylk7nfr52mwlud4jq0xm"

},{


    "value":	"jf1f5gmr5si1n2cmnt4t",

    "public_id":	"fe04rrwdhk645h2x2aa2170rh685g9",

    "target_id":	"c53oose8pup9ht4fc25x"

}]
```

## PUT /Token/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the Token type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record with the provided JSON request body.

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "value":	"2f0sf849hysnvfba0ntn",

    "public_id":	"pipxxn1snvqjgwzq5d7ldmzzqdh06w",

    "target_id":	"kqtulxiotdyvw1xdsna6"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /Token/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the Token type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Request Parameters

```
    :public_id
```

- Expected Status Code

```
Failure: 500
Success: 204
```

