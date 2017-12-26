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


    "name":	"Philip Spencer",

    "public_id":	"p6l6z27tsw23i8kabcslzek4aepvyy",

    "activities":	null,

    "created_at":	"2017-12-26 15:18:04.497569 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.49758 +0000 UTC"

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


    "name":	"Kelly Gomez",

    "public_id":	"1ng07n23tdlxhnamkjin6cbzboul3b",

    "activities":	null,

    "created_at":	"2017-12-26 15:18:04.498162 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.498167 +0000 UTC"

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


    "activities":	null,

    "created_at":	"2017-12-26 15:18:04.498561 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.498565 +0000 UTC",

    "name":	"Wayne Ryan",

    "public_id":	"6hlhkbdmojzjvndwymg6phlwg535vb"

},{


    "updated_at":	"2017-12-26 15:18:04.498889 +0000 UTC",

    "name":	"Roger Reed",

    "public_id":	"dxo6a45mwp52dhaj191nt9874xdjy5",

    "activities":	null,

    "created_at":	"2017-12-26 15:18:04.498884 +0000 UTC"

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


    "name":	"Mrs. Ms. Miss Carol Gonzales",

    "public_id":	"adihc4wil42m9xfr14kl5fbcbul7dj",

    "activities":	null,

    "created_at":	"2017-12-26 15:18:04.49933 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.499335 +0000 UTC"

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

