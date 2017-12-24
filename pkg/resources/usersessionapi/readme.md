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


    "email":	"SarahCastillo@Twitterbridge.name",

    "password":	"u719aw3xl6w5iy6bu5rb",

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


    "token":	"rr75dg9e7q6pb7j7cyp9",

    "expires":	"2017-12-24 17:21:03.746563 +0000 UTC",

    "user_id":	"kittejxr93pkwp4ahhbm",

    "public_id":	"ws9vvm3tmb5mrdrthoyju5o3h7jv6g"

}
```

## INFO /UserSession/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.UserSession.

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


    "expires":	"2017-12-24 17:21:03.747584 +0000 UTC",

    "user_id":	"o8rfbrx5bn8tqzauivks",

    "public_id":	"3qtp0qyik5fczhjysoo4xqi6z5v3pe",

    "token":	"hzy3f8opmujunvu2dl3q"

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


    "user_id":	"11mnt1tzhjfy0g3zh3aa",

    "public_id":	"nry2243dmunukt1wj3obv2bnrz4j9x",

    "token":	"vw62q2o7126bdamxt53c",

    "expires":	"2017-12-24 17:21:03.747992 +0000 UTC"

},{


    "user_id":	"lize1tr1ylvg9qsbnsks",

    "public_id":	"pl8d9mduwux0z4vyn7c3wall6g20yz",

    "token":	"42a7yjajh7h26iro1qv9",

    "expires":	"2017-12-24 17:21:03.748377 +0000 UTC"

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


    "token":	"5f9uitxisx28kds0m4l9",

    "expires":	"2017-12-24 17:21:03.74886 +0000 UTC",

    "user_id":	"le0b9dkybp73sqv72cv7",

    "public_id":	"vgjjtudya6palxizrpajkjygo701xg"

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

