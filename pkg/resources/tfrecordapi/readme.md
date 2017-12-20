TFRecord HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/tfrecordapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/tfrecordapi)

TFRecord HTTP API is a auto generated http api for the struct `TFRecord`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.TFRecord, error)
    Update(context.Context, string, pkg.TFRecord) error
    GetAll(context.Context, string, string, int, int) ([]pkg.TFRecord, int, error)
    Create(context.Context, pkg.NewTF) (pkg.TFRecord, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /TFRecord/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the NewTF type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "max_length":	10,

    "tenant":	nil,

    "user":	nil,

    "domain":	"Zazio.info"

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


    "public_id":	"5du3qhlqullok2s8ggk4u3zz8gcx99",

    "created_at":	"2017-12-20 15:06:24.610045 +0000 UTC",

    "updated_at":	"2017-12-20 15:06:24.610054 +0000 UTC",

    "key":	"002br4efph17x1qe7zuk",

    "domain":	"Zoombox.biz",

    "user_id":	"9bpwv8fpqkuua2eox5pj",

    "totp":	"2bq2lfk1fnw86x64lkhr",

    "tenant_id":	"z28y23svb40riirkvxso",

    "code_length":	11

}
```

## GET /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the TFRecord type from the HTTP API returning received result as a JSON
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


    "user_id":	"3342xurai4ijrcs3la7u",

    "totp":	"440ag17zq0x8iepog6jg",

    "created_at":	"2017-12-20 15:06:24.610601 +0000 UTC",

    "tenant_id":	"aa894brg2m12h8i1nikt",

    "code_length":	1,

    "updated_at":	"2017-12-20 15:06:24.610607 +0000 UTC",

    "key":	"4fbffh5lht65gpgq1mrh",

    "domain":	"Skivee.name",

    "public_id":	"pywp9j7g4wku1xrv6boaa4wyc93ssb"

}
```

## GET /TFRecord/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the TFRecord type from the HTTP API.

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


    "code_length":	11,

    "key":	"tyn746nk5vpe27g8zc5d",

    "totp":	"1h9my4l9ltcigbzzsjl9",

    "public_id":	"b6ok5plrrpt76bxgja5nx55rnsbus4",

    "tenant_id":	"pom20qq3m3tiafc32w3h",

    "created_at":	"2017-12-20 15:06:24.611208 +0000 UTC",

    "updated_at":	"2017-12-20 15:06:24.611214 +0000 UTC",

    "domain":	"Tanoodle.name",

    "user_id":	"4emvyy84e83hugudefel"

},{


    "tenant_id":	"mbt8u787fu2w65n1gown",

    "created_at":	"2017-12-20 15:06:24.61257 +0000 UTC",

    "updated_at":	"2017-12-20 15:06:24.612582 +0000 UTC",

    "domain":	"Twitternation.net",

    "user_id":	"vu8rkl07xrsnmefvy1bx",

    "public_id":	"mk0ovv9c6e8rnmlst1vvk9mxji39n9",

    "code_length":	17,

    "key":	"5ixjsl62sxv3jwzwgyx7",

    "totp":	"acegxyl74dikbav7q9p3"

}]
```

## PUT /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the TFRecord type from the HTTP API returning received result as a JSON
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


    "tenant_id":	"k4y37fgt60uxbm6nx4jo",

    "code_length":	0,

    "updated_at":	"2017-12-20 15:06:24.613926 +0000 UTC",

    "key":	"aqr4h1ozoq1g0nxvmvkb",

    "domain":	"Skippad.name",

    "user_id":	"k3nv2lbnw7v9s75ntp0w",

    "totp":	"t42ej5oxpys1coqk9gto",

    "public_id":	"w8rxluw1qpmbuu95limvta6ij2p1j7",

    "created_at":	"2017-12-20 15:06:24.613917 +0000 UTC"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the TFRecord type from the HTTP API returning received result as a JSON
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

