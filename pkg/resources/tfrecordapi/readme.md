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

    "tenant":	null,

    "user":	null,

    "domain":	"Buzzbean.com"

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


    "key":	"mtkyvdxjydj6i1tql0f1",

    "totp":	"5jgpuu36cgu79nm7c7ue",

    "tenant_id":	"622zyuqabhz0m9dofq3y",

    "updated_at":	"2017-12-23 12:57:49.244269 +0000 UTC",

    "domain":	"Youopia.net",

    "user_id":	"tmtm8f1qm7ojkychfsyt",

    "public_id":	"59n8jqbxd9rx0geksh4zxf58oa0oqr",

    "code_length":	11,

    "created_at":	"2017-12-23 12:57:49.244236 +0000 UTC"

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


    "key":	"5m19z0y5p4f2uly56mz6",

    "domain":	"Skalith.biz",

    "totp":	"wtid21hkto4qq7hxvm3u",

    "code_length":	1,

    "created_at":	"2017-12-23 12:57:49.245821 +0000 UTC",

    "user_id":	"ubd568f08y8am1hlcqo1",

    "public_id":	"k84f5mzko4tfrhab4m926w8jxoyjkw",

    "tenant_id":	"03qk7e8vffo87xsbx4f0",

    "updated_at":	"2017-12-23 12:57:49.245833 +0000 UTC"

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


    "domain":	"Yodoo.net",

    "public_id":	"xbcky0iqnznlynoln5kuz24bhji4h1",

    "code_length":	11,

    "created_at":	"2017-12-23 12:57:49.246333 +0000 UTC",

    "key":	"du5izqi4zqsvijwj3pw9",

    "user_id":	"4cqzbt1sdl1s0e4q1jr0",

    "totp":	"pwukchp3el9i354g0l3t",

    "tenant_id":	"ahm9apgwxbrjl2qwnxe0",

    "updated_at":	"2017-12-23 12:57:49.246338 +0000 UTC"

},{


    "key":	"tiat2dxguhq5ie7t1ebh",

    "user_id":	"14c3qsmllpyaz7dpkidd",

    "tenant_id":	"9ug8gtqo7z8q0l0peutu",

    "code_length":	17,

    "created_at":	"2017-12-23 12:57:49.246697 +0000 UTC",

    "domain":	"Tagfeed.info",

    "totp":	"sa465n4qabynuxykew23",

    "public_id":	"l1a983heqsmc5n66cbv1e04o1ax2tq",

    "updated_at":	"2017-12-23 12:57:49.246702 +0000 UTC"

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


    "key":	"tb7q7szbiy04cb3rh2rt",

    "domain":	"Ntag.biz",

    "public_id":	"mhacbvus0mjpq6sjwf5epaq2nwimfi",

    "created_at":	"2017-12-23 12:57:49.247084 +0000 UTC",

    "updated_at":	"2017-12-23 12:57:49.247089 +0000 UTC",

    "user_id":	"ept4e5m3dvol4coliai6",

    "totp":	"ih288syxupoa11l50p25",

    "tenant_id":	"ss18sdmg0k7kwneasp6f",

    "code_length":	0

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

