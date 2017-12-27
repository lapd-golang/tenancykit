Activity HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/activityapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/activityapi)

Activity HTTP API is a auto generated http api for the struct `Activity`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.Activity, error)
    Update(context.Context, string, pkg.Activity) error
    GetAll(context.Context, string, string, int, int) ([]pkg.Activity, int, error)
    Create(context.Context, pkg.ActivityName) (pkg.Activity, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Activity/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the ActivityName type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
pkg.ActivityName
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


    "name":	"Andrea Parker",

    "public_id":	"di41bl72xwz68oconjd4ezplb54ec9",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

}
```

## INFO /Activity/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.Activity.

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

## GET /Activity/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the Activity type from the HTTP API returning received result as a JSON
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


    "name":	"Mr. Dr. Albert Clark",

    "public_id":	"32gb1oe0g26tyq64cnoidcoaue1ims",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

}
```

## GET /Activity/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the Activity type from the HTTP API.

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


    "public_id":	"f65xw813sof7bc6h5olae91kb66d7t",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "name":	"George Medina"

},{


    "name":	"Gregory White",

    "public_id":	"gcl037sxm233i3m4khex0lt21k6209",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

}]
```

## PUT /Activity/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the Activity type from the HTTP API returning received result as a JSON
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


    "public_id":	"zl3ycsmybu3ytewtpruc2o5nqx2xux",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "name":	"Timothy Banks"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /Activity/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the Activity type from the HTTP API returning received result as a JSON
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

