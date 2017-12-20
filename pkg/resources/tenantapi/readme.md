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


    "name":	"Mrs. Ms. Miss Amanda Ruiz",

    "email":	"DennisHansen@Mynte.biz"

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


    "email":	"jCollins@Eabox.com",

    "public_id":	"ckmarje9rnw2fslt22u3qornol99f3",

    "created_at":	"2017-12-20 15:06:24.489028 +0000 UTC",

    "updated_at":	"2017-12-20 15:06:24.489489 +0000 UTC",

    "name":	"Paula Richardson"

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


    "name":	"Jacqueline Gonzales",

    "email":	"TheresaLane@Zoonder.com",

    "public_id":	"o1f08d5wz40kbzjfkol33a6vsgebaz",

    "created_at":	"2017-12-20 15:06:24.490351 +0000 UTC",

    "updated_at":	"2017-12-20 15:06:24.490359 +0000 UTC"

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


    "name":	"Susan Patterson",

    "email":	"StevePerry@Tazzy.biz",

    "public_id":	"v4d15f6fazp8l9fimt82zdvdf2z964",

    "created_at":	"2017-12-20 15:06:24.494795 +0000 UTC",

    "updated_at":	"2017-12-20 15:06:24.494807 +0000 UTC"

},{


    "updated_at":	"2017-12-20 15:06:24.49524 +0000 UTC",

    "name":	"Margaret Young",

    "email":	"qRuiz@Wordware.org",

    "public_id":	"h3hsgc5lfdo0h6ylaax0p7s36q6yy6",

    "created_at":	"2017-12-20 15:06:24.495233 +0000 UTC"

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


    "name":	"Shawn Kim",

    "email":	"SeanMorrison@Lajo.name",

    "public_id":	"a1p0vl3lh74sfh8v1a1t4132s1viie",

    "created_at":	"2017-12-20 15:06:24.495687 +0000 UTC",

    "updated_at":	"2017-12-20 15:06:24.495694 +0000 UTC"

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

