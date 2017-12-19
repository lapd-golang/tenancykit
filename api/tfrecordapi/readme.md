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


    "max_length":	5577006791947779410,

    "tenant":	nil,

    "user":	nil,

    "domain":	"2"

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


    "updated_at":	2017-12-19 13:42:58.59581 +0000 UTC,

    "user_id":	"y",

    "totp":	"q",

    "public_id":	"78ayn01ctf4o9zk",

    "tenant_id":	"k",

    "code_length":	8674665223082153551,

    "created_at":	2017-12-19 13:42:58.595786 +0000 UTC,

    "key":	"z",

    "domain":	"5"

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


    "key":	"5",

    "code_length":	6129484611666145821,

    "created_at":	2017-12-19 13:42:58.596198 +0000 UTC,

    "domain":	"z",

    "user_id":	"k",

    "totp":	"f",

    "public_id":	"9t3n0wpqwy5xctc",

    "tenant_id":	"2",

    "updated_at":	2017-12-19 13:42:58.596202 +0000 UTC

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


    "created_at":	2017-12-19 13:42:58.596563 +0000 UTC,

    "updated_at":	2017-12-19 13:42:58.596567 +0000 UTC,

    "user_id":	"c",

    "domain":	"4",

    "totp":	"h",

    "public_id":	"zgpgt77aois1qsy",

    "tenant_id":	"0",

    "code_length":	4037200794235010051,

    "key":	"v"

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


    "user_id":	"h",

    "code_length":	3916589616287113937,

    "key":	"k",

    "domain":	"8",

    "totp":	"p",

    "public_id":	"zg8uivlyfj7wcu5",

    "tenant_id":	"4",

    "created_at":	2017-12-19 13:42:58.596946 +0000 UTC,

    "updated_at":	2017-12-19 13:42:58.59695 +0000 UTC

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