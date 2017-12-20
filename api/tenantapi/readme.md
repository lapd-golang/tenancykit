Tenant HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/tenantapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/tenantapi)

Tenant HTTP API is a auto generated http api for the struct `Tenant`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.Tenant, error)
    Update(context.Context, string, tenancykit.Tenant) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.Tenant, int, error)
    Create(context.Context, tenancykit.CreateTenant) (tenancykit.Tenant, error)
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


    "name":	"Victor Berry Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"tStone@Browsetype.mil"

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


    "name":	"Ashley Olson",

    "email":	"MildredDay@Topiclounge.org",

    "public_id":	"t97hffyik0xvn1k7crjmt20gbbnez6",

    "created_at":	"2017-12-20 12:28:53.530486 +0000 UTC",

    "updated_at":	"2017-12-20 12:28:53.530542 +0000 UTC"

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


    "created_at":	"2017-12-20 12:28:53.531464 +0000 UTC",

    "updated_at":	"2017-12-20 12:28:53.531472 +0000 UTC",

    "name":	"Kenneth Peterson",

    "email":	"yWagner@Vimbo.com",

    "public_id":	"v1dcqhze6zp266573uk8oxqra3np0a"

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


    "created_at":	"2017-12-20 12:28:53.532588 +0000 UTC",

    "updated_at":	"2017-12-20 12:28:53.532597 +0000 UTC",

    "name":	"Amy Gutierrez",

    "email":	"AnnaJones@Jamia.name",

    "public_id":	"tqxmwo79eu62r2bbqcdj8vdfj1qf88"

},{


    "created_at":	"2017-12-20 12:28:53.533485 +0000 UTC",

    "updated_at":	"2017-12-20 12:28:53.533493 +0000 UTC",

    "name":	"Barbara Andrews",

    "email":	"BettyLong@Gevee.net",

    "public_id":	"rkrnhpfsow8rh1fnvg5njb14rapemh"

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


    "email":	"AngelaGonzalez@Jamia.info",

    "public_id":	"qtppvzmmz6i3iklai3yly1pj3d8kdk",

    "created_at":	"2017-12-20 12:28:53.535709 +0000 UTC",

    "updated_at":	"2017-12-20 12:28:53.53572 +0000 UTC",

    "name":	"Joe Morrison"

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

