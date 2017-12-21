TwoFactorSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/twofactorsessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/twofactorsessionapi)

TwoFactorSession HTTP API is a auto generated http api for the struct `TwoFactorSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.TwoFactorSession, error)
    Update(context.Context, string, pkg.TwoFactorSession) error
    GetAll(context.Context, string, string, int, int) ([]pkg.TwoFactorSession, int, error)
    Create(context.Context, pkg.TwoFactorSession) (pkg.TwoFactorSession, error)
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


    "user_id":	"yurpr6h5tt2ygznpqcd2",

    "public_id":	"oimai9brl002f2le8p1kthikq09c5q",

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

    "user_id":	"3hirp9xfm61k5oegqjlf",

    "public_id":	"zfwrh0q378qmxujzne9wycvyfq8glw"

}
```

## INFO /TwoFactorSession/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.TwoFactorSession.

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


    "public_id":	"s7sxzjzmr3qxmp34148zik2kcuwyll",

    "bool":	true,

    "user_id":	"xnttlqkksfw8pzg15z18"

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


    "public_id":	"4uqwez4atigwl76yk84vqicneftz9x",

    "bool":	true,

    "user_id":	"1kyabhkdpysl4l8s0phz"

},{


    "user_id":	"n0tqwans2lzp4v2m4gac",

    "public_id":	"pd522ne8haxiz8luyx80s1x5qgset8",

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


    "user_id":	"atboe3pi246mxuht7ypl",

    "public_id":	"uqvux647dv2uv80ea775dw9ry6nafd",

    "bool":	true

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

