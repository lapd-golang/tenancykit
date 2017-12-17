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


    "max_length":	0,

    "tenant":	nil,

    "user":	nil,

    "domain":	""

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


    "user_id":	"",

    "public_id":	"",

    "code_length":	0,

    "updated_at":	"2017-12-17 19:11:07.184008 +0000 UTC",

    "key":	"",

    "totp":	"",

    "tenant_id":	"",

    "created_at":	"2017-12-17 19:11:07.183989 +0000 UTC",

    "domain":	""

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


    "max_length":	0,

    "tenant":	nil,

    "user":	nil,

    "domain":	""

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


    "domain":	"",

    "max_length":	0,

    "tenant":	nil,

    "user":	nil

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


    "key":	"",

    "public_id":	"",

    "code_length":	0,

    "created_at":	"2017-12-17 19:11:07.192333 +0000 UTC",

    "updated_at":	"2017-12-17 19:11:07.192342 +0000 UTC",

    "domain":	"",

    "user_id":	"",

    "totp":	"",

    "tenant_id":	""

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