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


    "user_id":	"id3mrvdnpm8rhsspqkcj",

    "public_id":	"t5w6ru7u5se3xa73xt3fbahmetrhko",

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

    "user_id":	"lfv2loas81tx98rifn0f",

    "public_id":	"r4azd2015q9vjoitiqg0gh8yki139c"

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


    "bool":	true,

    "user_id":	"vtp58307ujamjg9b0nla",

    "public_id":	"y1w0fozf77wuiolvo1f22k3132tfcp"

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


    "user_id":	"bycbgik3txveparu8gmt",

    "public_id":	"ed3sd90fc5s8f3lpeva5c2hrqlq61y",

    "bool":	true

},{


    "user_id":	"ce9003vfbxmyd5an8lah",

    "public_id":	"g6bji36dq8kbbqa7h9ggbesckh84r2",

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


    "user_id":	"mt2t5zv62uup83nm76gb",

    "public_id":	"v0fpxw16sfd34clzvl6h44hx4rvk36",

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

