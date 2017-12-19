Token HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/tokenset/tokens/tokenapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/tokenset/tokens/tokenapi)

Token HTTP API is a auto generated http api for the struct `Token`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tokens.Token, error)
    Update(context.Context, string, tokens.Token) error
    GetAll(context.Context, string, string, int, int) ([]tokens.Token, int, error)
    Create(context.Context, tokens.Token) (tokens.Token, error)
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


    "value":	"vhvs4j5o4pbs4wb3da7d",

    "public_id":	"fg2wvoro8n5lsjd6lak06a26jrdivi",

    "target_id":	"5gmz3i3uy67vym60bq2g"

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


    "public_id":	"kdpvgnan4qbryrbhmudseeoxo0oybi",

    "target_id":	"j82vru21043rov29v0qr",

    "value":	"jz9spkki1bdrxsay87bp"

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


    "public_id":	"6gdpoduxburyxp9odz59o2wvdqk0iw",

    "target_id":	"evddqkxo9jwao4mj9i05",

    "value":	"p1ur7squz9gqv4bygmwf"

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


    "value":	"guxwxhr578ggufulzap2",

    "public_id":	"8isgpsezu4pksqwg9kza042yz7ccyx",

    "target_id":	"jtml2ey5blvf278umawv"

},{


    "value":	"li19pyzeqnqi186y7eeb",

    "public_id":	"89arusq3o3mem8iki6ysxgoqn952ur",

    "target_id":	"hnkic1lt92x8uj77n7zy"

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


    "target_id":	"jcou5wyenv0s7j4brqxq",

    "value":	"4ka4ny3mtwuv6zejjgbw",

    "public_id":	"431sq7ng5shwpwtdzwiw2en196nh4h"

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

