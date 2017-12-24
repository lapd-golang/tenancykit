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


    "name":	"Brandon Pierce",

    "email":	"LarryRomero@Avamba.edu"

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


    "name":	"Helen Clark",

    "email":	"eum@Jabberstorm.org",

    "public_id":	"zuofrv7xc0hr26v0lp38tu6axt7i63",

    "secret_id":	"9lbdlethdfznsxhnk2i4",

    "created_at":	"2017-12-24 18:16:09.033469 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.0335 +0000 UTC"

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


    "updated_at":	"2017-12-24 18:16:09.034648 +0000 UTC",

    "name":	"Jesse Young",

    "email":	"PhilipRyan@Skidoo.biz",

    "public_id":	"bq2dsxw1kof9cpcri489tm4lmsrlli",

    "secret_id":	"8irzsbpw6djfz9fusrjy",

    "created_at":	"2017-12-24 18:16:09.03464 +0000 UTC"

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


    "email":	"eius_est@Linkbuzz.edu",

    "public_id":	"azsdzdghhh2ksirj0geg6gikngmztl",

    "secret_id":	"rjrjfauo0xmngpq99e4b",

    "created_at":	"2017-12-24 18:16:09.035689 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.035697 +0000 UTC",

    "name":	"Thomas Lopez"

},{


    "created_at":	"2017-12-24 18:16:09.036775 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.036783 +0000 UTC",

    "name":	"Mr. Dr. Jack Cox",

    "email":	"xGarcia@Miboo.gov",

    "public_id":	"wwlljrqydtaf3on2bu8y00kazirfxs",

    "secret_id":	"36x172w84ukjs0j5e3hs"

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


    "secret_id":	"r22lu4hyocv143wo1hn9",

    "created_at":	"2017-12-24 18:16:09.037277 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.037283 +0000 UTC",

    "name":	"Tammy Lee",

    "email":	"ea_provident@Quatz.net",

    "public_id":	"e721iock3el8phadh1pddciue03rnm"

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

