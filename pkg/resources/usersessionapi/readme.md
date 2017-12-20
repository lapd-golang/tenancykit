UserSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/usersessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/usersessionapi)

UserSession HTTP API is a auto generated http api for the struct `UserSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.UserSession, error)
    Update(context.Context, string, pkg.UserSession) error
    GetAll(context.Context, string, string, int, int) ([]pkg.UserSession, int, error)
    Create(context.Context, pkg.CreateUserSession) (pkg.UserSession, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /UserSession/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the CreateUserSession type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "email":	"ElizabethHicks@Agimba.org",

    "password":	"yb4uihnc2bnkjlynzge6",

    "expiration":	null

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


    "expires":	"2017-12-20 19:23:27.468368 +0000 UTC",

    "user_id":	"hdfg14giqvos1m68g4x4",

    "public_id":	"lqudcnxu1n93aflz06x5afsm9iy705",

    "token":	"nqf22nlke26pg87exz7t"

}
```

## GET /UserSession/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the UserSession type from the HTTP API returning received result as a JSON
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


    "user_id":	"3848r6kh8nzj0wqcvzh7",

    "public_id":	"smvbeht0wxmx046ka6keuuyclpderj",

    "token":	"pdxign056qxl27nndefz",

    "expires":	"2017-12-20 19:23:27.468779 +0000 UTC"

}
```

## GET /UserSession/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the UserSession type from the HTTP API.

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


    "user_id":	"jf9mvtfx7b1sumj5rx1q",

    "public_id":	"pgvrxo2g0zfjyc2b6x7s21hs4kj7qw",

    "token":	"kldcxd6hhgbd4hv02u1v",

    "expires":	"2017-12-20 19:23:27.469536 +0000 UTC"

},{


    "user_id":	"e64vq6mrvlj1m5etxdqh",

    "public_id":	"epzpu1yhai8jxg55xuz3dxzzd3o2ie",

    "token":	"def5xjh4tsqqa4k8crrd",

    "expires":	"2017-12-20 19:23:27.47034 +0000 UTC"

}]
```

## PUT /UserSession/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the UserSession type from the HTTP API returning received result as a JSON
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


    "user_id":	"0p75qjf1msu3t3obpuzp",

    "public_id":	"f2hbcso2825htadjrxejhawlzd6vtp",

    "token":	"qz7djagej9ehc76ijsup",

    "expires":	"2017-12-20 19:23:27.471169 +0000 UTC"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /UserSession/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the UserSession type from the HTTP API returning received result as a JSON
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

