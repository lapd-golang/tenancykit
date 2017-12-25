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


    "name":	"Gary Hunt",

    "public_id":	"shfw3c347cm1uhnodkb2xst7cu9whz",

    "created_at":	"2017-12-25 20:33:45.372311 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.37234 +0000 UTC"

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


    "name":	"Mrs. Ms. Miss Elizabeth Daniels",

    "public_id":	"x4az0jx27vr1iwt75prayvu7m3t34x",

    "created_at":	"2017-12-25 20:33:45.373127 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.373136 +0000 UTC"

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


    "name":	"Carolyn Wallace I II III IV V MD DDS PhD DVM",

    "public_id":	"uzo2oyha3zzil3nuh5op3cgvp791df",

    "created_at":	"2017-12-25 20:33:45.373621 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.373631 +0000 UTC"

},{


    "name":	"Carolyn Murray",

    "public_id":	"mib93isclcm1qv5zwnc9333yiinjvc",

    "created_at":	"2017-12-25 20:33:45.374008 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.374016 +0000 UTC"

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


    "name":	"Alan Dean Jr. Sr. I II III IV V MD DDS PhD DVM",

    "public_id":	"4bzmzoepqf25fw42cbro4k4m52cv86",

    "created_at":	"2017-12-25 20:33:45.374781 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.374794 +0000 UTC"

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

