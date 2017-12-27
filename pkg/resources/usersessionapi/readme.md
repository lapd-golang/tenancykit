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


    "expiration":	null,

    "email":	"explicabo_qui_molestiae@Skivee.name",

    "password":	"6lntyxquy7fbctitltbw"

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


    "public_id":	"g0cfx3xwmvgivrylbczbab4okpitl3",

    "token":	"aqes2nqihvf73st855ss",

    "expires":	"2017-12-27T11:14:17+01:00",

    "user_id":	"p061c6o3vkk2oe4uy7t1"

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


    "user_id":	"ap8au14ypayz1c00vn73",

    "public_id":	"cidoyhxtk1bjkpby469ryphym2elfz",

    "token":	"b53t6pgxtfairupr2sq8",

    "expires":	"2017-12-27T11:14:17+01:00"

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


    "token":	"hbx0uuunkdekttyuhne1",

    "expires":	"2017-12-27T11:14:17+01:00",

    "user_id":	"gz7vtz0833zmhrfk66sd",

    "public_id":	"ihk3jn8e7f7xssxvng372sxf7sezc7"

},{


    "user_id":	"19gfp81e3ji9afabwgwa",

    "public_id":	"sh6zx2kmj54wa5lbgomae81cylucj6",

    "token":	"e40h8mkiml89qvqrtlou",

    "expires":	"2017-12-27T11:14:17+01:00"

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


    "public_id":	"bezljq5rv0il6mcfsp0ngb30kej5qo",

    "token":	"ealt922u1tca0zrf3f6e",

    "expires":	"2017-12-27T11:14:17+01:00",

    "user_id":	"s7p2pbvmuzywqs674g29"

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

