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


    "name":	"Beverly Harper I II III IV V MD DDS PhD DVM",

    "email":	"eius_autem_et@Jabbersphere.biz"

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


    "name":	"Lisa Campbell",

    "email":	"DeniseBurke@Viva.net",

    "public_id":	"5gl3b52vlktqggtf3ic5a9vn7nxen5",

    "created_at":	"2017-12-24 06:38:45.110807 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.11084 +0000 UTC"

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


    "created_at":	"2017-12-24 06:38:45.112373 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.112382 +0000 UTC",

    "name":	"Alice Anderson",

    "email":	"KeithReyes@Eare.mil",

    "public_id":	"iruooqnlzj56fcl6wwttdu22pezt5k"

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


    "email":	"at_placeat@Voonder.mil",

    "public_id":	"t5031h21njafpcd893tw20mf23caa4",

    "created_at":	"2017-12-24 06:38:45.113832 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.113844 +0000 UTC",

    "name":	"Fred Freeman"

},{


    "public_id":	"yxucdachxv43bsfrqz6ao0oo5kkea8",

    "created_at":	"2017-12-24 06:38:45.114797 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.11481 +0000 UTC",

    "name":	"Matthew Duncan",

    "email":	"ShirleyDavis@Quinu.name"

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


    "name":	"Debra Porter I II III IV V MD DDS PhD DVM",

    "email":	"RonaldJackson@Skyvu.edu",

    "public_id":	"l50kizftcup3ccjhemcimwog3ya1w1",

    "created_at":	"2017-12-24 06:38:45.115401 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.11541 +0000 UTC"

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

