TwoFactorSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/twofactorsessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/twofactorsessionapi)

TwoFactorSession HTTP API is a auto generated http api for the struct `TwoFactorSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.TwoFactorSession, error)
    Update(context.Context, string, tenancykit.TwoFactorSession) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.TwoFactorSession, int, error)
    Create(context.Context, tenancykit.TwoFactorSession) (tenancykit.TwoFactorSession, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /TwoFactorSession/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the TwoFactorSession type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "user_id":	"c8d7chpo1eygi3ihsorg",

    "public_id":	"7aou38llg28s3okomuzt83meotkyyf",

    "bool":	true

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


    "bool":	true,

    "user_id":	"slhfkpcvzat1hachhcnf",

    "public_id":	"6pglpqz9rribfrtsb4iilhxdrb1zwp"

}
```

## GET /TwoFactorSession/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the TwoFactorSession type from the HTTP API returning received result as a JSON
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


    "public_id":	"p16bpv7xvhph0snmscomn14xwjfw6k",

    "bool":	true,

    "user_id":	"fetzy4230mx7s23r2pqo"

}
```

## GET /TwoFactorSession/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the TwoFactorSession type from the HTTP API.

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


    "user_id":	"mdb6jb3g9nae1n8t9i37",

    "public_id":	"d92burmb648thta0iqkqfuy5gopeof",

    "bool":	true

},{


    "user_id":	"xql159ptvtqgswcsh6s6",

    "public_id":	"c4hkkemz9v64qqzebbdgkrc5uq4xfd",

    "bool":	true

}]
```

## PUT /TwoFactorSession/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the TwoFactorSession type from the HTTP API returning received result as a JSON
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


    "public_id":	"7ill77rkub70lf99dls8w371omghgp",

    "bool":	true,

    "user_id":	"hptl5bx330y5dems4um3"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /TwoFactorSession/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the TwoFactorSession type from the HTTP API returning received result as a JSON
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

