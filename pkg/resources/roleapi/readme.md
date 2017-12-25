Role HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/roleapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/roleapi)

Role HTTP API is a auto generated http api for the struct `Role`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.Role, error)
    Update(context.Context, string, pkg.Role) error
    GetAll(context.Context, string, string, int, int) ([]pkg.Role, int, error)
    Create(context.Context, pkg.RoleName) (pkg.Role, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Role/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the RoleName type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
pkg.RoleName
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


    "updated_at":	"2017-12-25 20:33:45.560295 +0000 UTC",

    "name":	"Jonathan Jackson",

    "public_id":	"2j60swm3xybfmgyh8fnjxzjtojt4pg",

    "activities":	null,

    "created_at":	"2017-12-25 20:33:45.560273 +0000 UTC"

}
```

## INFO /Role/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.Role.

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

## GET /Role/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the Role type from the HTTP API returning received result as a JSON
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


    "name":	"Cheryl Simmons",

    "public_id":	"vjovtpw36n3gc9xt8v57zt1hhso8lt",

    "activities":	null,

    "created_at":	"2017-12-25 20:33:45.5612 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.56121 +0000 UTC"

}
```

## GET /Role/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the Role type from the HTTP API.

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


    "name":	"Robert Medina",

    "public_id":	"g563rr4q0yb6yni5z8ucaz31drbbfa",

    "activities":	null,

    "created_at":	"2017-12-25 20:33:45.561825 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.561835 +0000 UTC"

},{


    "public_id":	"ok4pegvu6sn1yqpm6kpx4haq2nevsi",

    "activities":	null,

    "created_at":	"2017-12-25 20:33:45.562341 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.562349 +0000 UTC",

    "name":	"Frances Gibson I II III IV V MD DDS PhD DVM"

}]
```

## PUT /Role/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the Role type from the HTTP API returning received result as a JSON
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


    "name":	"Lawrence Gordon",

    "public_id":	"6phs7nlsosec5ttc9twp71infcz0y7",

    "activities":	null,

    "created_at":	"2017-12-25 20:33:45.562993 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.563005 +0000 UTC"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /Role/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the Role type from the HTTP API returning received result as a JSON
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

