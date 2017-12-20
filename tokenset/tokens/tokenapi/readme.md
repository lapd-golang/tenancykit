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


    "value":	"q5m8fl02pat00euiyi4o",

    "public_id":	"1lv9snpx46hehijyy9cnc5i5estjwo",

    "target_id":	"apz0jy00o1d2qqmtfl2f"

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


    "value":	"dmuwfkymgnja3d1dmwe8",

    "public_id":	"n3emgvkvwcofffu11xqwyfwzb71lte",

    "target_id":	"vufzyv735y00joppc9qp"

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


    "value":	"3k2jpzsar0ojzagm1bqz",

    "public_id":	"37b66qkf9hp3xahrlg57x6i4kd3h1j",

    "target_id":	"06dz2dcojgtjslk1lq5m"

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


    "value":	"5xkmflt59zlvxqzhkau2",

    "public_id":	"zb98om4hrfyfbrv8vj2kdnx8cqmka0",

    "target_id":	"h7y9a345fz6jx5d63uhy"

},{


    "value":	"3c9lcog4bselezsj86fx",

    "public_id":	"a62nxhmk563sw2x9ul9xzxq06ei58m",

    "target_id":	"c6axrf5znpc5x6r68tww"

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


    "value":	"lvac7xu5p8spi2qnsiag",

    "public_id":	"sweto2bw8ihzcgxfecn3f3l08cws1c",

    "target_id":	"oejxy9x7ohp1925ag8zg"

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

