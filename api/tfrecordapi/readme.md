TFRecord HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/tfrecordapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/tfrecordapi)

TFRecord HTTP API is a auto generated http api for the struct `TFRecord`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.TFRecord, error)
    Update(context.Context, string, tenancykit.TFRecord) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.TFRecord, int, error)
    Create(context.Context, tenancykit.NewTF) (tenancykit.TFRecord, error)
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


    "domain":	"BlogXS.edu",

    "max_length":	7,

    "tenant":	nil,

    "user":	nil

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


    "public_id":	"eed6chcyaqlveruhtjsfp29zfhz4ok",

    "code_length":	14,

    "key":	"aj036y77oa06dc8djmy9",

    "domain":	"Snaptags.net",

    "tenant_id":	"8hz13jlwbdv1r3dl9gx6",

    "created_at":	"2017-12-20 08:05:05.584616 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.584628 +0000 UTC",

    "user_id":	"z0ngfg4sjc9ioowo78zs",

    "totp":	"5obj8halpcu4md3m81jf"

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


    "tenant_id":	"566d5ymq2d3oy15eru99",

    "created_at":	"2017-12-20 08:05:05.585907 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.585916 +0000 UTC",

    "key":	"sqjfnpn8ns6eb07m5tpo",

    "domain":	"Eire.edu",

    "user_id":	"0jwu3441z09hzitdr4or",

    "totp":	"vjtg96iicoemonok46dq",

    "public_id":	"oo7du7k7wz6xi1d84n8kdedmscda2n",

    "code_length":	16

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


    "totp":	"b9da98p8hm4ay6bfrba5",

    "public_id":	"h6pnri4t05uqkxltlvbf7dgzj9oxbx",

    "tenant_id":	"s2vui5ofj780147hk31i",

    "code_length":	15,

    "created_at":	"2017-12-20 08:05:05.587228 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.587244 +0000 UTC",

    "key":	"tm964lzvuv95zbmqgpt8",

    "domain":	"Realbridge.name",

    "user_id":	"aavim3k11binmtrbv6us"

},{


    "public_id":	"s1mdj0xa8zessyylgnjsyv2gq5e7b5",

    "tenant_id":	"5l1plhkn8l2q63pct3p4",

    "code_length":	13,

    "domain":	"Photospace.mil",

    "user_id":	"wecki257sjlowtinluos",

    "totp":	"nvp02tay9y6tvyffum02",

    "key":	"trmagtv2zchyrwgdzwax",

    "created_at":	"2017-12-20 08:05:05.587942 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.587948 +0000 UTC"

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


    "public_id":	"kj76j601faatbfleukcbqlivmj7uqq",

    "tenant_id":	"vx419aulhfnnkt4e3acn",

    "updated_at":	"2017-12-20 08:05:05.588468 +0000 UTC",

    "key":	"r3ayot0q1ljtabh0mbr7",

    "domain":	"Kwilith.edu",

    "code_length":	8,

    "created_at":	"2017-12-20 08:05:05.588462 +0000 UTC",

    "user_id":	"pbmbzu9t757wkncohqqv",

    "totp":	"udss80vy1j3v9zo73oj3"

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

