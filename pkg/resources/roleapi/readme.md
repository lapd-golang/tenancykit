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


    "activities":	null,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "name":	"Patricia Meyer",

    "public_id":	"84u0his2njnjnyz5944u54w3maj1zx"

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


    "public_id":	"6m4yqinv45zq5ctfqsc8a8kzaoe0o0",

    "activities":	null,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "name":	"William Hayes"

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


    "updated_at":	"2017-12-27T11:14:17+01:00",

    "name":	"Mrs. Ms. Miss Gloria Andrews",

    "public_id":	"clx5el9snwqrmxdcu60r8yjeqt9qii",

    "activities":	null,

    "created_at":	"2017-12-27T11:14:17+01:00"

},{


    "name":	"Margaret Alvarez",

    "public_id":	"qxodl1vuqnicnf0l63zx9nl26bjlc4",

    "activities":	null,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

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


    "name":	"Bonnie Evans",

    "public_id":	"8zv7g77bnvpuz3w730umt97nzkvmjc",

    "activities":	null,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

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

