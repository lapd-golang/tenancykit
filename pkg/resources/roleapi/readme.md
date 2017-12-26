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


    "updated_at":	"2017-12-26T18:57:09+01:00",

    "name":	"Joseph Peters",

    "public_id":	"58ya451z58dvo7ledqaw1sf4k3tth6",

    "activities":	null,

    "created_at":	"2017-12-26T18:57:09+01:00"

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


    "name":	"Gloria Welch",

    "public_id":	"avj0aoj433uxxm52u6v57t4gzap621",

    "activities":	null,

    "created_at":	"2017-12-26T18:57:09+01:00",

    "updated_at":	"2017-12-26T18:57:09+01:00"

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


    "updated_at":	"2017-12-26T18:57:09+01:00",

    "name":	"Denise Sims",

    "public_id":	"q1dczczsr78ure2w2plznhwkkk1boq",

    "activities":	null,

    "created_at":	"2017-12-26T18:57:09+01:00"

},{


    "name":	"Jane Bryant",

    "public_id":	"kvptfroksigi6e8tfx8s4z8t2q9djj",

    "activities":	null,

    "created_at":	"2017-12-26T18:57:09+01:00",

    "updated_at":	"2017-12-26T18:57:09+01:00"

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


    "updated_at":	"2017-12-26T18:57:09+01:00",

    "name":	"Matthew Burton Jr. Sr. I II III IV V MD DDS PhD DVM",

    "public_id":	"2tsaj4kpzyvf7tlywksc6gpr5a563u",

    "activities":	null,

    "created_at":	"2017-12-26T18:57:09+01:00"

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

