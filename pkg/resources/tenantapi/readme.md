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


    "name":	"Susan Simmons",

    "email":	"eius_voluptatem_in@Quaxo.org"

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


    "email":	"totam_accusantium_ad@Oloo.org",

    "public_id":	"nuv331afnzuqj5tunpv4qiepbfgjky",

    "secret_id":	"upk7pjv2v7o4upxmhmad",

    "created_at":	"2017-12-26 15:18:04.546708 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.546717 +0000 UTC",

    "name":	"Jean Jackson"

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


    "name":	"Andrew Coleman",

    "email":	"ullam_voluptatem@Voomm.edu",

    "public_id":	"rypp4ret2v74lei1y0b6meiqyocf1i",

    "secret_id":	"1p3q3hnsd7o69h3ug1mo",

    "created_at":	"2017-12-26 15:18:04.547293 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.5473 +0000 UTC"

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


    "name":	"Elizabeth Hicks I II III IV V MD DDS PhD DVM",

    "email":	"GeraldLane@Flashpoint.org",

    "public_id":	"2zve7djjcax8czgpnxr5o9aspd3ilf",

    "secret_id":	"v878iyqfoziot23goaob",

    "created_at":	"2017-12-26 15:18:04.547772 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.547778 +0000 UTC"

},{


    "name":	"Debra Sanders",

    "email":	"pGreene@Skivee.edu",

    "public_id":	"ma0yj8mvcpaayskyvgwarq575wn2a4",

    "secret_id":	"evbkrsxve8x99evgy7f3",

    "created_at":	"2017-12-26 15:18:04.54819 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.548196 +0000 UTC"

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


    "name":	"Jesse Kelley Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"vel@Meedoo.org",

    "public_id":	"ohu1oxuzo0kggty9asd198jmzc0mbs",

    "secret_id":	"a2ulcglhoo8z484w7eds",

    "created_at":	"2017-12-26 15:18:04.548659 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.548664 +0000 UTC"

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

