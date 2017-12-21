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


    "max_length":	4,

    "tenant":	null,

    "user":	null,

    "domain":	"Jayo.biz"

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


    "key":	"bryhqf4xtny6oftvqbb7",

    "domain":	"Blogspan.name",

    "user_id":	"srridgzc7dqygmvugglq",

    "totp":	"mfnommqi5zbcwmh6jxad",

    "public_id":	"fqxrl0geqevr4reryuv2yyu4u8rowu",

    "code_length":	7,

    "tenant_id":	"o2ruy97w54hbjd4i9ik2",

    "created_at":	"2017-12-21 13:40:49.368449 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.368458 +0000 UTC"

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


    "tenant_id":	"vuj0rlj1dl7jpufqndw8",

    "code_length":	14,

    "updated_at":	"2017-12-21 13:40:49.368995 +0000 UTC",

    "public_id":	"i2als8paozye7yh6fett5q2b2fzanb",

    "created_at":	"2017-12-21 13:40:49.368991 +0000 UTC",

    "key":	"6z5rvb6hhuj99f4s5z4x",

    "domain":	"Oba.org",

    "user_id":	"fy7mq8zqw7zfmdasrays",

    "totp":	"ahzbecdwcxc475eck2dl"

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


    "domain":	"Pixope.gov",

    "user_id":	"j8j8s7rgydxnuqbpk1fh",

    "tenant_id":	"xtrgqs53k0rlt3an7nbr",

    "code_length":	16,

    "updated_at":	"2017-12-21 13:40:49.369583 +0000 UTC",

    "key":	"l7u6q5zr5sxqosqt1qjv",

    "totp":	"zpkirie263vnnlmjqmqa",

    "public_id":	"u3tba7oamxz24kfojyxih1aytf9zbc",

    "created_at":	"2017-12-21 13:40:49.369578 +0000 UTC"

},{


    "created_at":	"2017-12-21 13:40:49.370126 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.370131 +0000 UTC",

    "totp":	"dr5sezwqx33s05tbartm",

    "tenant_id":	"v9u5wvxcc8iqe5h380le",

    "user_id":	"iie59zejbpe7rb03py4p",

    "public_id":	"1nhwe18qq36zyqh00huuuk85efl4f0",

    "code_length":	15,

    "key":	"m9h5r84l00tqapvqzrr6",

    "domain":	"Meembee.name"

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


    "tenant_id":	"t8kxnn4iwo7afivkmonp",

    "created_at":	"2017-12-21 13:40:49.370582 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.370588 +0000 UTC",

    "key":	"8ugxdig0tyu4swzhzgpu",

    "user_id":	"mgw6toiu6dgxhzkhet08",

    "public_id":	"s1bc2lp14ifuju2lpbl57zqea8pdk9",

    "domain":	"Kazu.name",

    "totp":	"7jw5xhioodtpp6pwkuma",

    "code_length":	13

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

