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
    Create(context.Context, pkg.Role) (pkg.Role, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Role/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the Role type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "name":	"Samuel Matthews",

    "public_id":	"3me4k6d91kyqbf7t7q95vvtgbxurlk",

    "activities":	null,

    "created_at":	"2017-12-25 16:32:46.040976 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.040995 +0000 UTC"

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


    "name":	"Mr. Dr. Benjamin Butler",

    "public_id":	"bjkcxrt7nz2bhtuqafia3ahvggxm4s",

    "activities":	null,

    "created_at":	"2017-12-25 16:32:46.041879 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.04189 +0000 UTC"

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


    "name":	"Christopher Ramirez",

    "public_id":	"bwfn9b7uomajy2s7kf1q3z502y22lj",

    "activities":	null,

    "created_at":	"2017-12-25 16:32:46.043358 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.043369 +0000 UTC"

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


    "created_at":	"2017-12-25 16:32:46.043978 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.04399 +0000 UTC",

    "name":	"Patrick Boyd",

    "public_id":	"rokgvth5d4s1llesmgy4gke8uihn8u",

    "activities":	null

},{


    "name":	"Mrs. Ms. Miss Joan Bradley",

    "public_id":	"ffsqzqdawkrysp7d5x4rhft3i0iasv",

    "activities":	null,

    "created_at":	"2017-12-25 16:32:46.044437 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.044444 +0000 UTC"

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


    "name":	"Christopher Hernandez",

    "public_id":	"uky3r5rmz6wrdnnqmbvwt7x4f7m7g7",

    "activities":	null,

    "created_at":	"2017-12-25 16:32:46.044992 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.045 +0000 UTC"

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

