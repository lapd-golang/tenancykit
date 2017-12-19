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


    "name":	"Eric Larson",

    "email":	"qMiller@Twimbo.name"

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


    "name":	"Anthony Rogers Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"VictorWells@Brainbox.mil",

    "public_id":	"5k191z9e8xlkotsojo3axb78eqcaig",

    "created_at":	"2017-12-19 18:15:32.282065 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.282093 +0000 UTC"

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


    "email":	"ChrisReed@Reallinks.name",

    "public_id":	"whgi03nqydwfm2czfug6hnggx07i33",

    "created_at":	"2017-12-19 18:15:32.282606 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.282613 +0000 UTC",

    "name":	"Anna Stephens"

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


    "name":	"Susan Mitchell",

    "email":	"fCrawford@Browsezoom.org",

    "public_id":	"7xz5cl534fkquhek1ww84lw9tv6omb",

    "created_at":	"2017-12-19 18:15:32.28304 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.283045 +0000 UTC"

},{


    "name":	"Jacqueline Morrison I II III IV V MD DDS PhD DVM",

    "email":	"TinaWeaver@Voonix.info",

    "public_id":	"mfztk7l8omqq6p6981tiu66fwuhxz5",

    "created_at":	"2017-12-19 18:15:32.283312 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.283317 +0000 UTC"

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


    "created_at":	"2017-12-19 18:15:32.284082 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.284091 +0000 UTC",

    "name":	"Stephen Nelson",

    "email":	"5Hayes@Yakidoo.mil",

    "public_id":	"p4ykmcbcn2jz7aj9zxm7xtbpytaveu"

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

