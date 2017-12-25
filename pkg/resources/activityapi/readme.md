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
    Create(context.Context, pkg.Activity) (pkg.Activity, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Activity/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the Activity type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "name":	"Shirley Allen",

    "public_id":	"8x8vjpwj7poheoa887z3fjqxdppm1s",

    "created_at":	"2017-12-25 16:32:45.961329 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.961337 +0000 UTC"

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


    "created_at":	"2017-12-25 16:32:45.961833 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.961839 +0000 UTC",

    "name":	"Arthur Ramirez",

    "public_id":	"awv7ptorqfmps6c5cxtj4uoo2pyu6e"

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


    "public_id":	"yev67cubuqhjwcblwa6skqwdu0dknp",

    "created_at":	"2017-12-25 16:32:45.962345 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.962351 +0000 UTC",

    "name":	"Elizabeth Morales"

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


    "name":	"Amanda Wheeler",

    "public_id":	"g1t0h5f30wm6t3t6tcet3fdfwza8ka",

    "created_at":	"2017-12-25 16:32:45.962807 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.962815 +0000 UTC"

},{


    "created_at":	"2017-12-25 16:32:45.963131 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.963135 +0000 UTC",

    "name":	"Joshua Watkins",

    "public_id":	"9rf7lxd3kyuitix40ocyavaq6plm1x"

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


    "updated_at":	"2017-12-25 16:32:45.963503 +0000 UTC",

    "name":	"Sharon Hill",

    "public_id":	"o7hap9gkqtg1w7n3ngdlgjysimliec",

    "created_at":	"2017-12-25 16:32:45.963498 +0000 UTC"

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

