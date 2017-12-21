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


    "domain":	"Yabox.org",

    "max_length":	17,

    "tenant":	null,

    "user":	null

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


    "user_id":	"40yadorc96s0ey9upflz",

    "totp":	"ntmplq4hrz27vzidlex2",

    "public_id":	"mn18bth1gcb2fj4w3afqm246hiixir",

    "code_length":	0,

    "updated_at":	"2017-12-21 20:42:41.484353 +0000 UTC",

    "key":	"hr7ii5po497yeklwqfbv",

    "domain":	"Dynabox.net",

    "tenant_id":	"d0a9xp9souv8ju2jl7y9",

    "created_at":	"2017-12-21 20:42:41.484345 +0000 UTC"

}
```

## INFO /TFRecord/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.TFRecord.

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


    "domain":	"Thoughtworks.name",

    "user_id":	"noi6plisz6gpb8kdtd91",

    "totp":	"zf45mppimt1hfhr9jsq8",

    "public_id":	"dyead9loze0crxlhtpwnq9t8xtw2ao",

    "code_length":	18,

    "key":	"ulgoo2bi945tmk1qb27l",

    "tenant_id":	"vn6f8mawn4x13j5qxabb",

    "created_at":	"2017-12-21 20:42:41.485218 +0000 UTC",

    "updated_at":	"2017-12-21 20:42:41.485224 +0000 UTC"

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


    "tenant_id":	"uu8wl5et02ezmd3gcacv",

    "created_at":	"2017-12-21 20:42:41.485911 +0000 UTC",

    "updated_at":	"2017-12-21 20:42:41.485918 +0000 UTC",

    "user_id":	"89uxoy03o3my4khvmxwf",

    "domain":	"Ntag.gov",

    "totp":	"jaewqu1mpqhq6kjj284n",

    "public_id":	"uqpuoaltb1f6856u10juzb27fjgv1a",

    "code_length":	8,

    "key":	"opqrikrs6an45mme0id2"

},{


    "updated_at":	"2017-12-21 20:42:41.487085 +0000 UTC",

    "tenant_id":	"faqxzrcpqxfy8foer8qy",

    "code_length":	16,

    "created_at":	"2017-12-21 20:42:41.487074 +0000 UTC",

    "totp":	"6aq2pysap0zjr7jl3bdh",

    "public_id":	"f9c141m4gjokddijqxliie7d1sdtqd",

    "key":	"nmk4y5c7jdeovdw0vy70",

    "domain":	"Blogpad.gov",

    "user_id":	"a5izaw91o8du6ai3ogch"

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


    "key":	"rs9yk3yayfgsdhugp1qy",

    "domain":	"Tazz.gov",

    "public_id":	"b1hl4a3lu9v8aaojq3ihri291aciaa",

    "tenant_id":	"t2cqev29zpbw7n2099pa",

    "code_length":	9,

    "user_id":	"0j5m3q8crqtopfhvcmiu",

    "totp":	"e493w4wot40pj3dy98ws",

    "created_at":	"2017-12-21 20:42:41.48881 +0000 UTC",

    "updated_at":	"2017-12-21 20:42:41.48882 +0000 UTC"

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

