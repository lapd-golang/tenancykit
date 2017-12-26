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


    "updated_at":	"2017-12-26 15:18:04.283408 +0000 UTC",

    "name":	"Kevin Porter",

    "public_id":	"bs6jigt7v63oeje3scrnzszzez49o9",

    "created_at":	"2017-12-26 15:18:04.283382 +0000 UTC"

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


    "name":	"Edward Hill",

    "public_id":	"lm89opljd2h2ef4e34fli9iz9g95vt",

    "created_at":	"2017-12-26 15:18:04.283961 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.283968 +0000 UTC"

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


    "name":	"Arthur Henry",

    "public_id":	"u19y07whh7gicgrw2mgv5c05pkv7x1",

    "created_at":	"2017-12-26 15:18:04.284324 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.284331 +0000 UTC"

},{


    "name":	"Carolyn Bishop",

    "public_id":	"9ahv7r44vnx0duzr1s2mr3m2304rio",

    "created_at":	"2017-12-26 15:18:04.284931 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.284938 +0000 UTC"

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


    "name":	"Doris Romero",

    "public_id":	"paqeevlvgwq91u5k4puqyuo19nwiif",

    "created_at":	"2017-12-26 15:18:04.285275 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.285281 +0000 UTC"

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

