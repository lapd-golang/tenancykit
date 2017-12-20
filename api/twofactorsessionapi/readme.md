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


    "user_id":	"8rwzbruwra8xuss45g68",

    "public_id":	"04zrc9y4seeq80farfzc62kszc2ypl",

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


    "user_id":	"3mjeglakkf578822vuul",

    "public_id":	"2pc6fdm1c3ohrs04lgxksixsu2a0tg",

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


    "user_id":	"if5wldy8szr2u1gvanv2",

    "public_id":	"33k5y45ml1e9ln3a337zlv99iqayfk",

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


    "user_id":	"ndkezahbi21b54xm33rn",

    "public_id":	"k6bxm0q2gbbpzcq8eubkllzz7hviie",

    "bool":	true

},{


    "user_id":	"yrymw0v6sa50zr0kt1oo",

    "public_id":	"vetsi6f6ap5ena3ekg3gk7su9lbeqd",

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


    "user_id":	"rccxf03bs8dz1qsjy82t",

    "public_id":	"ox5jopkr8q2bslzdzywvm7v84gvooo",

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

