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


    "value":	"wsoa7qitnaby3imfpo31",

    "public_id":	"9uoj97s2uf2q64ap10gbpaoi039e1n",

    "target_id":	"ntk7cft0r2pyqsek43p7"

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


    "value":	"eqgjmvg0td5235f0dvgx",

    "public_id":	"kbntfrsog2t5ddv08a14v7yedqueau",

    "target_id":	"hw4vti95bxwn0nkltp94"

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


    "target_id":	"f1jf1k93jgyme3nqr0m7",

    "value":	"cw9s9cpehl15s0q3x2yt",

    "public_id":	"3lh2ufa3bwi6wbt8hb50zqpl8dhcgj"

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


    "value":	"c6zrcvw04fpe6eoo0oq4",

    "public_id":	"8qq5itd5firt5exnysobx1bnochpzc",

    "target_id":	"qfvaoukm5k36zmgejxka"

},{


    "value":	"m9sdkv82ivp8epkd2nzk",

    "public_id":	"gqpp54gpnwosdbywglhz0dnkcsma5r",

    "target_id":	"bs6zycpjeqvs110xz9vu"

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


    "target_id":	"rk9vchpau6jh1zmrkold",

    "value":	"a3lapgypyg26jkiqm3jr",

    "public_id":	"ehab9x3gamah46bngf0pmhnsdg69ej"

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

