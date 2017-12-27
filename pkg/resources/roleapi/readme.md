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


    "name":	"Jason Allen",

    "public_id":	"ejaprbc3ofqcnfigcnf72qo3cql5dz",

    "activities":	null,

    "created_at":	"2017-12-27T17:22:58+01:00",

    "updated_at":	"2017-12-27T17:22:58+01:00"

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


    "updated_at":	"2017-12-27T17:22:58+01:00",

    "name":	"Frances Hanson",

    "public_id":	"hdxmjcojv5tr0dcno96ct80x18x9us",

    "activities":	null,

    "created_at":	"2017-12-27T17:22:58+01:00"

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


    "public_id":	"cb2kqyw1ueu2wp1o1ddr09tw9oydoe",

    "activities":	null,

    "created_at":	"2017-12-27T17:22:58+01:00",

    "updated_at":	"2017-12-27T17:22:58+01:00",

    "name":	"Frank Bowman"

},{


    "name":	"Kathryn Willis",

    "public_id":	"1v8smg99xsry6y44t5u36fqeq1ka8i",

    "activities":	null,

    "created_at":	"2017-12-27T17:22:58+01:00",

    "updated_at":	"2017-12-27T17:22:58+01:00"

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


    "public_id":	"hy8pygzo1c92fy356ylca0noxxnept",

    "activities":	null,

    "created_at":	"2017-12-27T17:22:58+01:00",

    "updated_at":	"2017-12-27T17:22:58+01:00",

    "name":	"Betty Romero"

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

