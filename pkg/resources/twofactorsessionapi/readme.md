TwoFactorSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/twofactorsessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/twofactorsessionapi)

TwoFactorSession HTTP API is a auto generated http api for the struct `TwoFactorSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.TwoFactorSession, error)
    Update(context.Context, string, pkg.TwoFactorSession) error
    GetAll(context.Context, string, string, int, int) ([]pkg.TwoFactorSession, int, error)
    Create(context.Context, pkg.TwoFactorSession) (pkg.TwoFactorSession, error)
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


    "user_id":	"3nn5a1o8q3hdxgh1l8ou",

    "public_id":	"mypid1u0hzd75kkp5bs932gui27kyd",

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


    "user_id":	"c7swe17o38zr0pjw9msp",

    "public_id":	"urh2xdhcl2mc7e14tlxb88ysfryosg",

    "bool":	true

}
```

## INFO /TwoFactorSession/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.TwoFactorSession.

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


    "user_id":	"luh4ssy8pv76rtqlfuq3",

    "public_id":	"h5uj2ux7e2mfbhxpis417iplhkgzk5",

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


    "user_id":	"w01ookf0sao7ivrwn1ps",

    "public_id":	"vqsdq5csl06cryiw1ceocprngg8nrd",

    "bool":	true

},{


    "bool":	true,

    "user_id":	"wo2gc7x9bkzz3h2xwhcc",

    "public_id":	"1mvwjwnon0klevl5xbn5nlymud0eq6"

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


    "public_id":	"whavs0ouiuw4ulyqz4l42xwm47evst",

    "bool":	true,

    "user_id":	"qxnxuzeq268xkhkmd2rh"

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

