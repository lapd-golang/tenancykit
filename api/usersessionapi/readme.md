UserSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/usersessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/usersessionapi)

UserSession HTTP API is a auto generated http api for the struct `UserSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.UserSession, error)
    Update(context.Context, string, tenancykit.UserSession) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.UserSession, int, error)
    Create(context.Context, tenancykit.CreateUserSession) (tenancykit.UserSession, error)
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


    "email":	"molestiae_consequatur_maiores@Youbridge.info",

    "password":	"rpczvrcxzy1uegaat5fy",

    "expiration":	nil

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


    "user_id":	"l2vw2muejq0pv6tt5slj",

    "public_id":	"gosu5khg9stejv0g8r75lumpc9ifyq",

    "token":	"3uewl4jujqf2ip59qn82",

    "expires":	"2017-12-19 18:15:32.454413 +0000 UTC"

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


    "user_id":	"aru3yn7ahwt9tnkqmxbm",

    "public_id":	"piyleiknglyhtvqm2s1a0hf9z3gwjg",

    "token":	"6zxkr36d47c27dw6y8ie",

    "expires":	"2017-12-19 18:15:32.454934 +0000 UTC"

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


    "user_id":	"xemudko3jhbsel46ut7f",

    "public_id":	"oggu6xblq2rnklzbs4ulp8ho5q4xn7",

    "token":	"w0agsqs5hax33138lkft",

    "expires":	"2017-12-19 18:15:32.45544 +0000 UTC"

},{


    "expires":	"2017-12-19 18:15:32.455919 +0000 UTC",

    "user_id":	"ckrybvzz7t0peeiiky8t",

    "public_id":	"hnedmuzbeovn92c2tn246ui6dieodw",

    "token":	"ewlvz5rx4oi0m6c8qogt"

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


    "token":	"i0heh2abgrpdwunm0045",

    "expires":	"2017-12-19 18:15:32.456521 +0000 UTC",

    "user_id":	"jns61f8x9m7sz18i4ig5",

    "public_id":	"kv48cbvajal2lxf5salorxxll8musl"

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

