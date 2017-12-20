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


    "email":	"oHarper@Tagopia.biz",

    "password":	"x6db21cywourj2zv2n7n",

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


    "user_id":	"hhgy677lsuobd586497r",

    "public_id":	"gbyu9xlb4dles5xedor40ocdsyjm6s",

    "token":	"gzilc5vpsd1ykmvvodk0",

    "expires":	"2017-12-20 08:05:05.532164 +0000 UTC"

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


    "expires":	"2017-12-20 08:05:05.532641 +0000 UTC",

    "user_id":	"6y7iuu3dtvd9nq8sa7mj",

    "public_id":	"0wdaw9eqxp2wj69puonet4dtfnhvl2",

    "token":	"txuubd1gpxzsp1fgl1lc"

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


    "token":	"z8adqreq23nb8g52sj69",

    "expires":	"2017-12-20 08:05:05.533133 +0000 UTC",

    "user_id":	"pv01n92k43gi856r43q9",

    "public_id":	"yttrqthmo3tejglu98tcld19bc9z1x"

},{


    "public_id":	"wxgb7n1nanlvz9k7kh9nnip7djrqk4",

    "token":	"7lo599sm3s4nsu3c9ye4",

    "expires":	"2017-12-20 08:05:05.533549 +0000 UTC",

    "user_id":	"dgoknxtk7y2what3kgvg"

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


    "user_id":	"ib9r8vdn1vewe456i7si",

    "public_id":	"h2ugzljpp0h5n92nsky9m4msjdl3t0",

    "token":	"uq483u1xwddj7uea6icw",

    "expires":	"2017-12-20 08:05:05.534011 +0000 UTC"

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

