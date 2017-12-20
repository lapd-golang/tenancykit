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


    "max_length":	0,

    "tenant":	nil,

    "user":	nil,

    "domain":	"Zoozzy.biz"

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


    "updated_at":	"2017-12-20 10:14:32.026383 +0000 UTC",

    "key":	"zwouakt1fiqjhbhqnjou",

    "domain":	"Browseblab.biz",

    "user_id":	"s1953k2n50546vhb4fhy",

    "totp":	"9t00qawl0k1lfbk7xeoq",

    "created_at":	"2017-12-20 10:14:32.026374 +0000 UTC",

    "public_id":	"yaxdlgcqsgxmmajgds90udu1sb0wkl",

    "tenant_id":	"n5c880gy9db37dreq7g7",

    "code_length":	18

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


    "key":	"i8n5ajdjo5lj0c10kkav",

    "totp":	"mm64mkoxtgxnhisr2onu",

    "code_length":	8,

    "domain":	"Twinte.gov",

    "user_id":	"htlrfhm66euq8w8qz6im",

    "public_id":	"paadx3n7nzb2v88rgax1vwhllum5yp",

    "tenant_id":	"59ggiu8kr1t2bln8bdzi",

    "created_at":	"2017-12-20 10:14:32.026857 +0000 UTC",

    "updated_at":	"2017-12-20 10:14:32.026861 +0000 UTC"

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


    "domain":	"Kwimbee.edu",

    "totp":	"c0llku1w1p55ardu7c3m",

    "public_id":	"m7oo11f9urqjn0bmo3gcnbylbfdb6q",

    "tenant_id":	"jc9kp323mql8iuhyr2et",

    "key":	"qdp684adcv43pzgcwxcs",

    "user_id":	"28sxn36svpxg0yr0sdsn",

    "code_length":	16,

    "created_at":	"2017-12-20 10:14:32.027364 +0000 UTC",

    "updated_at":	"2017-12-20 10:14:32.027368 +0000 UTC"

},{


    "user_id":	"6lh3fgajpymz5j6zajsm",

    "code_length":	9,

    "key":	"a2jcl1x48jzl1tkiyrez",

    "totp":	"0ksm8kf6tvite3683whh",

    "public_id":	"yna389mw052lo9yasgepkj3pb839xf",

    "tenant_id":	"o0rcddd2js83bb6vfcqx",

    "created_at":	"2017-12-20 10:14:32.027837 +0000 UTC",

    "updated_at":	"2017-12-20 10:14:32.027843 +0000 UTC",

    "domain":	"Twimbo.net"

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


    "key":	"igxppn9fmwfptxi6scsa",

    "totp":	"6qmjuw8a124ekubv79tr",

    "updated_at":	"2017-12-20 10:14:32.028372 +0000 UTC",

    "code_length":	4,

    "created_at":	"2017-12-20 10:14:32.028367 +0000 UTC",

    "domain":	"Zoonoodle.edu",

    "user_id":	"02s5rbp3chctm663xf6l",

    "public_id":	"04tf465stdf2rhaotpoezmyp4krqt0",

    "tenant_id":	"lc9w0eyhchqlft0i4k7q"

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

