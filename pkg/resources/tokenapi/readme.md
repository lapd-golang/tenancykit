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


    "value":	"v0b6898q9hi38q1jrrk5",

    "public_id":	"c8hiqc5awvvcwfkvdzdboybjakxqs6",

    "target_id":	"fnr4jbjgmw1dq1k4o9wq"

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


    "public_id":	"6w1gnvriikk7o6arq90w961jude8fx",

    "target_id":	"0kyycwi67o46r4obd3z1",

    "value":	"p652q1iky86m7ozy4qmb"

}
```

## INFO /Token/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.Token.

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


    "value":	"qyqygo4gjyuoxdgpzj4m",

    "public_id":	"h5203ug026fjrarl9x9o9a54phbsmx",

    "target_id":	"8e4noex15bqu9u6zvfsh"

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


    "value":	"tu4djgdnzpgb9vfmm7px",

    "public_id":	"151my45qxkzkktz68m8etoacbfswj2",

    "target_id":	"wiy0q5dfyh996buzu5xh"

},{


    "value":	"qgiksprkran4r577x8e2",

    "public_id":	"n51acsw6j74lnok1sftm6cvjbwmofp",

    "target_id":	"zrsh52ake9xp4krcvxs4"

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


    "value":	"uyyuxrg1xqawojza4t8h",

    "public_id":	"19392agev211uyklgjl30g1gqau42e",

    "target_id":	"yvehndssms9g25u0uh1i"

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

