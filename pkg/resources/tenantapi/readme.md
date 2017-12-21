Tenant HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/tenantapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/tenantapi)

Tenant HTTP API is a auto generated http api for the struct `Tenant`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.Tenant, error)
    Update(context.Context, string, pkg.Tenant) error
    GetAll(context.Context, string, string, int, int) ([]pkg.Tenant, int, error)
    Create(context.Context, pkg.CreateTenant) (pkg.Tenant, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Tenant/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the CreateTenant type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "name":	"Timothy Mitchell",

    "email":	"ChristopherDixon@Muxo.mil"

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


    "public_id":	"w5kjjft1hb3d3cxjxylig35qul7sm6",

    "created_at":	"2017-12-21 20:42:41.225976 +0000 UTC",

    "updated_at":	"2017-12-21 20:42:41.225999 +0000 UTC",

    "name":	"Marie Medina",

    "email":	"numquam@Zoonoodle.com"

}
```

## INFO /Tenant/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.Tenant.

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

## GET /Tenant/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the Tenant type from the HTTP API returning received result as a JSON
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


    "created_at":	"2017-12-21 20:42:41.226347 +0000 UTC",

    "updated_at":	"2017-12-21 20:42:41.226353 +0000 UTC",

    "name":	"Sara Perez",

    "email":	"SeanMills@Yoveo.edu",

    "public_id":	"as0w220axs46j4920eyju1huc4bsph"

}
```

## GET /Tenant/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the Tenant type from the HTTP API.

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


    "name":	"Patricia Elliott",

    "email":	"TimothyLee@Tazz.info",

    "public_id":	"cf4wy5q56ch71dtsk1m7rx2xz4txrd",

    "created_at":	"2017-12-21 20:42:41.226617 +0000 UTC",

    "updated_at":	"2017-12-21 20:42:41.226622 +0000 UTC"

},{


    "name":	"Andrea Reed",

    "email":	"nam_molestiae_quia@Photofeed.com",

    "public_id":	"oryujosexg97qb89sre1fivwqgcm4e",

    "created_at":	"2017-12-21 20:42:41.226901 +0000 UTC",

    "updated_at":	"2017-12-21 20:42:41.226907 +0000 UTC"

}]
```

## PUT /Tenant/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the Tenant type from the HTTP API returning received result as a JSON
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


    "updated_at":	"2017-12-21 20:42:41.227343 +0000 UTC",

    "name":	"Cynthia Miller",

    "email":	"beatae_laborum_rerum@Aimbo.org",

    "public_id":	"slo01vpqxwt63nyqio1164yt21iel6",

    "created_at":	"2017-12-21 20:42:41.227337 +0000 UTC"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /Tenant/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the Tenant type from the HTTP API returning received result as a JSON
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

