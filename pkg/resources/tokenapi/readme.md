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


    "value":	"qlads8bozg74xwgv60ql",

    "public_id":	"wr9rf3ck4uu713okcpggjo7key2hp3",

    "target_id":	"a0t0wvmsgev0xl5p5fnq"

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


    "public_id":	"j1j3wherhdjf515l965gm891rj7h6f",

    "target_id":	"ml4835wpnhiz8kwetz2v",

    "value":	"khmnzzxh8xkybzowjh81"

}
```

## INFO /Token/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.Token.

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
    "total": 10,
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


    "value":	"c0ak4nmmgb2u0jwtrsji",

    "public_id":	"hq933jgfidcgrhof3qu8epbm71l8ue",

    "target_id":	"3wzs7z8j331fkeakak0h"

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


    "value":	"jiou9vvpq4xejpe887i2",

    "public_id":	"l1phfwaznul3cbxpurv5i1z9mmquwa",

    "target_id":	"l73afegvwd4yxz1442eg"

},{


    "value":	"puf7cpi70p8xwuj7tcc5",

    "public_id":	"gj8swmzk61f5qgblhua071ktefj1g4",

    "target_id":	"u3cmakgmpwrlltggyxl9"

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


    "value":	"3l8xy4w5srjlqsyshmcy",

    "public_id":	"x287zx00zrm1qnag3yy34ja8xy82ki",

    "target_id":	"r1fz06mh9f2n68zmb416"

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

