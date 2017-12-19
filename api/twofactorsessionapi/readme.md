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


    "user_id":	"sfdv06uwe9b80peej465",

    "public_id":	"yyw5o8lt3ww45qmvtx76g68hr0okiy",

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


    "user_id":	"byqlt9tfz1awf3d5mqcr",

    "public_id":	"8kpphtfd3jhvtpaopxt8mlmz1mw7nv",

    "bool":	true

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


    "public_id":	"xaqlp5vkuhluwkozxsy5t8wkp02pas",

    "bool":	true,

    "user_id":	"kvjk36otko63amcwxq0s"

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


    "user_id":	"5nhvjtwo1urt88f8mf3s",

    "public_id":	"eb32lpnxczmu2jazdwg8tirs9ybjh6",

    "bool":	true

},{


    "user_id":	"fdlpt7gsc78bb7sw5rzd",

    "public_id":	"4fylo2cklxfrvkum125ljmdf8h7t57",

    "bool":	true

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


    "user_id":	"ouguc7vrsxstd7xqgmkx",

    "public_id":	"ej2my2d4fyxfbwfpvfip63rvuoif0d",

    "bool":	true

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

