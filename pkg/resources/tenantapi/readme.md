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


    "name":	"Mary Fowler I II III IV V MD DDS PhD DVM",

    "email":	"eum_qui@Plambee.info"

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


    "email":	"fOlson@Midel.net",

    "public_id":	"2hkzweet1jp8rc03r8gy9sn58ipdaa",

    "created_at":	"2017-12-23 12:57:49.46011 +0000 UTC",

    "updated_at":	"2017-12-23 12:57:49.460118 +0000 UTC",

    "name":	"Phillip James"

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


    "email":	"nCampbell@Babbleopia.com",

    "public_id":	"qr4n8rj8nl0m0gv9wel0yg8hbteb4h",

    "created_at":	"2017-12-23 12:57:49.460771 +0000 UTC",

    "updated_at":	"2017-12-23 12:57:49.460777 +0000 UTC",

    "name":	"Ruth Fuller"

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


    "email":	"aut@Eazzy.com",

    "public_id":	"h06pwo5yzes8pcsjntxmcrcxlgf370",

    "created_at":	"2017-12-23 12:57:49.461303 +0000 UTC",

    "updated_at":	"2017-12-23 12:57:49.461311 +0000 UTC",

    "name":	"Kimberly Fields"

},{


    "created_at":	"2017-12-23 12:57:49.461834 +0000 UTC",

    "updated_at":	"2017-12-23 12:57:49.461844 +0000 UTC",

    "name":	"Arthur Anderson",

    "email":	"BrianGrant@Avavee.mil",

    "public_id":	"mqfvmo4qb09jarbekew95m94ci084h"

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


    "updated_at":	"2017-12-23 12:57:49.462601 +0000 UTC",

    "name":	"Maria Holmes",

    "email":	"doloribus@Vipe.org",

    "public_id":	"5uvjqkhkgt6ydl50s59kbcef19cjdq",

    "created_at":	"2017-12-23 12:57:49.46259 +0000 UTC"

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

