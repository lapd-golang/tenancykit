Token HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/tokenapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/tokenapi)

Token HTTP API is a auto generated http api for the struct `Token`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.Token, error)
    Update(context.Context, string, pkg.Token) error
    GetAll(context.Context, string, string, int, int) ([]pkg.Token, int, error)
    Create(context.Context, pkg.Token) (pkg.Token, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Token/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the Token type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "value":	"wti05jgjw8xji8paritq",

    "public_id":	"pt0kgbwhudehhbrh9i1xex2qtq3but",

    "target_id":	"d7d9e3kenuflt0n4ld22"

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


    "public_id":	"p6tl0ajk6tsk1a5rngkubntv8oare4",

    "target_id":	"mjrovobshec6jsjw7std",

    "value":	"whklxm6utp3sowmxjg5r"

}
```

## GET /Token/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the Token type from the HTTP API returning received result as a JSON
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


    "value":	"tmlorizsy1fmo73ufkfb",

    "public_id":	"ybtl1zjlbvoyvopesij5rieczrfou9",

    "target_id":	"x032toflbbm1pwu9wp13"

}
```

## GET /Token/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the Token type from the HTTP API.

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


    "value":	"naie38vq1zn5q1aoc7y0",

    "public_id":	"8dy58xdjeefiv6n4p5qzbwlrir4i98",

    "target_id":	"tr4kx7u6c9qxykccbfwc"

},{


    "public_id":	"izpojx83fbur5686phajieo23e0ja0",

    "target_id":	"22j3s4va5a6lq1novbef",

    "value":	"56z0l5c1iy5a9kglqpwo"

}]
```

## PUT /Token/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the Token type from the HTTP API returning received result as a JSON
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


    "value":	"zwu8obmtr2ym7603k3s0",

    "public_id":	"oxo6vykrizfk3vh7dn8g7jc9l9mqg9",

    "target_id":	"0c1bn3h1zbh0uwyrr43i"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /Token/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the Token type from the HTTP API returning received result as a JSON
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

