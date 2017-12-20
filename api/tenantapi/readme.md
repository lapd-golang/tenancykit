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


    "name":	"Brandon Ruiz Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"ChristopherJohnston@Trilith.net"

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


    "name":	"Joseph Powell",

    "email":	"in_dolor@Zava.biz",

    "public_id":	"lw6nwebustuarx7c3mochvttuut6ri",

    "created_at":	"2017-12-20 10:14:31.82501 +0000 UTC",

    "updated_at":	"2017-12-20 10:14:31.825032 +0000 UTC"

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


    "name":	"Willie Daniels",

    "email":	"ut_blanditiis@Leenti.org",

    "public_id":	"ky48zrmfktiz5p5fn8shaseinnd2rw",

    "created_at":	"2017-12-20 10:14:31.825316 +0000 UTC",

    "updated_at":	"2017-12-20 10:14:31.825321 +0000 UTC"

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


    "name":	"Jesse Moreno",

    "email":	"hKennedy@Demivee.info",

    "public_id":	"ig4zquz226o47hd84n7kq1nna8u2rr",

    "created_at":	"2017-12-20 10:14:31.825582 +0000 UTC",

    "updated_at":	"2017-12-20 10:14:31.825586 +0000 UTC"

},{


    "name":	"Theresa Oliver",

    "email":	"iFoster@Livetube.biz",

    "public_id":	"1dlkrh4n9l91q70d022pgqla8t7ifd",

    "created_at":	"2017-12-20 10:14:31.827033 +0000 UTC",

    "updated_at":	"2017-12-20 10:14:31.827043 +0000 UTC"

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


    "updated_at":	"2017-12-20 10:14:31.827948 +0000 UTC",

    "name":	"Ruth Matthews",

    "email":	"VictorStewart@Babbleset.info",

    "public_id":	"ofg4b28cshni11i3g8thf6u25dvz86",

    "created_at":	"2017-12-20 10:14:31.827942 +0000 UTC"

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

