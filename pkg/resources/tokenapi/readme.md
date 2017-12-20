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


    "value":	"7vijhiebog2g32j3dbnx",

    "public_id":	"9kxl30dcrrq7ywcwnyw17coiosbsb6",

    "target_id":	"wdbfo84227u9uvlxyskm"

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


    "value":	"1e0a4o9mhrg18jvi5uz4",

    "public_id":	"2m1eb98pfl80t8dglijwuy9piyguwu",

    "target_id":	"bmduf1pzfsnd5jvkm9te"

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


    "value":	"9ehx1oxak95ivj8tbp37",

    "public_id":	"hlbi57tlyr4d62uj1o75t9h9e5ihiw",

    "target_id":	"5ndq2dfl1vqbcidhl5pz"

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


    "value":	"6yha4qspibetz4dac68i",

    "public_id":	"2h25sz4vum3liilskwao9jkwf23bql",

    "target_id":	"4dd67nrcche83zkuk0jb"

},{


    "value":	"zeohotf9yhl13f3asonf",

    "public_id":	"4nldh7a5n2ypkbe75ur78tw3jvkved",

    "target_id":	"3yjdsc50hnq1ei12skp3"

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


    "value":	"duqc176l35wcw2tn3bzf",

    "public_id":	"87soge7hx7iz7glxe0h2rc4omu5ky0",

    "target_id":	"fr60bggd50nnrxye8e1h"

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

