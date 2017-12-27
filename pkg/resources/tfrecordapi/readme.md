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

    "user":	null,

    "domain":	"Linkbuzz.net"

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


    "key":	"uz7abjv8k1xjlq28qypu",

    "domain":	"Ooba.name",

    "user_id":	"dsztneqtomm6d4ovdj3v",

    "totp":	"uhgcrisznj0yyvfhnlot",

    "tenant_id":	"3t67oxw0ss2zysjswwz9ajtciscbmo",

    "code_length":	11,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

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


    "domain":	"Dynabox.net",

    "user_id":	"r2a7b1wq0nk7atg8fmbh",

    "totp":	"axw2autfhsljj4jt10bx",

    "tenant_id":	"ea5qad5iovkqnojywco5lvt55ix5ol",

    "code_length":	1,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "key":	"sh4cpk1lsf3xl6qi427o"

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


    "user_id":	"qaozbisjl7vcrkmtaw0g",

    "totp":	"0ch01ma4rpbert97iht4",

    "tenant_id":	"mom3lzjquuzx39pydq3kyqzmthc88l",

    "code_length":	11,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "key":	"lpjo8vkk60bjrwd6eajj",

    "domain":	"Meevee.edu"

},{


    "tenant_id":	"voiut2h6eulpq7p0zzlpfkjvjd2oy4",

    "code_length":	17,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "key":	"yp4d46hq6flw9372kd7a",

    "domain":	"Edgetag.net",

    "user_id":	"etsqdwq45e76r8fhw0b6",

    "totp":	"05dwpzz488ebtpkv6lwj"

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


    "totp":	"ppcqdoveh13clfin1e83",

    "tenant_id":	"rw6bu8cqsx56cmsm05ek3f6g3usg4i",

    "code_length":	0,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "key":	"3nrzng9dve02xxq5tajm",

    "domain":	"Izio.com",

    "user_id":	"mandejk0pklpbd89hiom"

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

