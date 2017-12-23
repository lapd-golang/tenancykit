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


    "email":	"RuthKennedy@Eabox.mil",

    "password":	"4e070yyv5yzr6v04vn4q",

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


    "user_id":	"2x8pw3y1zdtpte96gkb1",

    "public_id":	"53a2ap032ay27viogilh47u6u0t6ns",

    "token":	"frmblm42ng3ztu6dfbo7",

    "expires":	"2017-12-23 12:57:49.357167 +0000 UTC"

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


    "token":	"8ofiflvadvkw2fxeidjw",

    "expires":	"2017-12-23 12:57:49.357678 +0000 UTC",

    "user_id":	"vbo1bil4a6971lcxwuwn",

    "public_id":	"5iado1laj44biwht433f61zs5pqean"

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


    "expires":	"2017-12-23 12:57:49.35816 +0000 UTC",

    "user_id":	"726wmk5k05rr8lfzsavk",

    "public_id":	"1y1f2iu9c3h012fbpboycz9hkj4q0m",

    "token":	"u380mq42fwbc5n9tzbfr"

},{


    "user_id":	"iwmwwjc0d0zxk3f8hcgg",

    "public_id":	"zd44uqzkibshfnlmtyrvgn906eapbj",

    "token":	"kogr65ztbmy1hpqtquzb",

    "expires":	"2017-12-23 12:57:49.358697 +0000 UTC"

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


    "user_id":	"6eyduxqvk8wwj6nikszw",

    "public_id":	"ydp8t471s9xxxdd4ffhkos0j6108yu",

    "token":	"c126wmmcofwrwwkeen14",

    "expires":	"2017-12-23 12:57:49.359103 +0000 UTC"

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

