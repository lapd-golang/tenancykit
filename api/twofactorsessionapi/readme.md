TwoFactorSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/twofactorsessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/twofactorsessionapi)

TwoFactorSession HTTP API is a auto generated http api for the struct `TwoFactorSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.TwoFactorSession, error)
    Update(context.Context, string, tenancykit.TwoFactorSession) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.TwoFactorSession, int, error)
    Create(context.Context, tenancykit.TwoFactorSession) (tenancykit.TwoFactorSession, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /TwoFactorSession/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the TwoFactorSession type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "user_id":	"375qibi7420fx8wq0o58",

    "public_id":	"6xpw3inkf6qkz413m50z7nju37ikwp",

    "bool":	true

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


    "bool":	true,

    "user_id":	"2j3kt3c5973opesdctpb",

    "public_id":	"f5xaeot6pl8zkdl04ior66r9nuw6w0"

}
```

## GET /TwoFactorSession/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the TwoFactorSession type from the HTTP API returning received result as a JSON
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


    "user_id":	"74boguthc5i2q2qap533",

    "public_id":	"xvldwht76f8lrlrivelab7zqj02t6k",

    "bool":	true

}
```

## GET /TwoFactorSession/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the TwoFactorSession type from the HTTP API.

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


    "user_id":	"2668v39mqebvrxysdoay",

    "public_id":	"m67kn8htxko31szys47loweavk2b0l",

    "bool":	true

},{


    "bool":	true,

    "user_id":	"syyi0ytojrw4vrnjckpk",

    "public_id":	"t9r18w8y4yib6hg5fwpxx907xa9wgg"

}]
```

## PUT /TwoFactorSession/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the TwoFactorSession type from the HTTP API returning received result as a JSON
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


    "public_id":	"94uikk0ju0tecglkqidjm54ecokhrj",

    "bool":	true,

    "user_id":	"5pg8ihhe5w8wasdtsilt"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /TwoFactorSession/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the TwoFactorSession type from the HTTP API returning received result as a JSON
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

