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


    "name":	"Benjamin Fernandez",

    "email":	"zHarper@Chatterbridge.net"

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


    "created_at":	"2017-12-25 16:32:46.147718 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.147726 +0000 UTC",

    "name":	"Sharon Morgan",

    "email":	"RogerMorgan@Meeveo.net",

    "public_id":	"eje6wif7rk1m7trdvutdk04t4xzoyi",

    "secret_id":	"8wdmzoxfncp71vu7qlek"

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


    "name":	"Nicholas Black",

    "email":	"incidunt@Voonix.gov",

    "public_id":	"oq8d3wvup0576dtwzpeafe8ch4x5ik",

    "secret_id":	"mi0d7w2ch2vqj7m5q7gv",

    "created_at":	"2017-12-25 16:32:46.148301 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.148307 +0000 UTC"

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


    "email":	"NicoleWarren@Bubblebox.biz",

    "public_id":	"45pibaeux5kege9n7zq7p0bpkqw0ln",

    "secret_id":	"ogtfkzpqa4r33fzgedog",

    "created_at":	"2017-12-25 16:32:46.148772 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.148778 +0000 UTC",

    "name":	"Steven Wallace"

},{


    "name":	"Phillip Foster Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"7Baker@Browsezoom.org",

    "public_id":	"5vj4gon43pqk3mnttn94cgj15x8y7s",

    "secret_id":	"npw9xw7kzu2kxib076kz",

    "created_at":	"2017-12-25 16:32:46.149185 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.14919 +0000 UTC"

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


    "email":	"JerryFoster@Meetz.net",

    "public_id":	"l8g0xxpd9ly5lr9pg6kk66kcmxkyk7",

    "secret_id":	"o5cfnrxa5fn6afbnhpet",

    "created_at":	"2017-12-25 16:32:46.149547 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.149552 +0000 UTC",

    "name":	"Eric Lynch"

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

