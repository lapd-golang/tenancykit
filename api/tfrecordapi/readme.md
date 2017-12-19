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

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
```

- Expected Request Body

```json
{


    "max_length":	2933568871211445515,

    "tenant":	nil,

    "user":	nil,

    "domain":	"9"

}
```

- Expected Status Code

```
Failure: 500
Success: 201
```

- Expected Response Body

```json
{


    "tenant_id":	"f",

    "created_at":	"2017-12-19 14:01:02.812613 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.812623 +0000 UTC",

    "domain":	"h",

    "totp":	"e",

    "public_id":	"u3y2h3eihfxaqqf",

    "code_length":	4324745483838182873,

    "key":	"r",

    "user_id":	"a"

}
```

## GET /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the TFRecord type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```json
{


    "key":	"w",

    "domain":	"v",

    "public_id":	"9ao7po9ttthrs7f",

    "code_length":	2610529275472644968,

    "created_at":	"2017-12-19 14:01:02.813168 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.813174 +0000 UTC",

    "user_id":	"h",

    "totp":	"g",

    "tenant_id":	"d"

}
```

## GET /TFRecord/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the TFRecord type from the HTTP API.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```json
[{


    "tenant_id":	"a",

    "created_at":	"2017-12-19 14:01:02.813742 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.813747 +0000 UTC",

    "key":	"9",

    "domain":	"f",

    "user_id":	"c",

    "totp":	"x",

    "public_id":	"k8hf2s14bibawi6",

    "code_length":	2703387474910584091

}]
```

## PUT /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the TFRecord type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record with the provided JSON request body.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
{


    "totp":	"f",

    "public_id":	"d6kbntcat50kccr",

    "tenant_id":	"4",

    "key":	"g",

    "domain":	"e",

    "created_at":	"2017-12-19 14:01:02.814167 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.814172 +0000 UTC",

    "user_id":	"7",

    "code_length":	6263450610539110790

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```


- Expected Response Body

```json
```

## DELETE /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the TFRecord type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 204
```

- Expected Response Body

```json
```