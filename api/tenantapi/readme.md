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
    Create(context.Context, tenancykit.NewTenant) (tenancykit.Tenant, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Tenant/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the NewTenant type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
```

- Expected Request Body

```json
{


    "name":	"",

    "email":	""

}
```

- Expected Status Code

```
Failure: 500
Success: 201
```

- Expected Response Body

```json
{


    "name":	"",

    "email":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:11:27.098843 +0000 UTC",

    "updated_at":	"2017-12-17 19:11:27.098852 +0000 UTC"

}
```

## GET /Tenant/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the Tenant type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```json
{


    "email":	"",

    "name":	""

}
```

## GET /Tenant/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the Tenant type from the HTTP API.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```json
[{


    "name":	"",

    "email":	""

}]
```

## PUT /Tenant/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the Tenant type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record with the provided JSON request body.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
{


    "name":	"",

    "email":	"",

    "public_id":	"",

    "created_at":	"2017-12-17 19:11:27.101441 +0000 UTC",

    "updated_at":	"2017-12-17 19:11:27.101451 +0000 UTC"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```


- Expected Response Body

```json
```

## DELETE /Tenant/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the Tenant type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 204
```

- Expected Response Body

```json
```