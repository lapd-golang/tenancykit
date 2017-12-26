UserSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/usersessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/usersessionapi)

UserSession HTTP API is a auto generated http api for the struct `UserSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.UserSession, error)
    Update(context.Context, string, pkg.UserSession) error
    GetAll(context.Context, string, string, int, int) ([]pkg.UserSession, int, error)
    Create(context.Context, pkg.CreateUserSession) (pkg.UserSession, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /UserSession/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the CreateUserSession type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "email":	"1Coleman@Babbleopia.gov",

    "password":	"cshvuw9aax2l91hcaqz5",

    "expiration":	null

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


    "token":	"9w4fnjd0gh8h8yznphdu",

    "expires":	"2017-12-26T18:57:10+01:00",

    "user_id":	"fatqw2gi73z3f6l0qlht",

    "public_id":	"ocj21639v4s1d5huzsvwu8c3gr56a1"

}
```

## INFO /UserSession/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.UserSession.

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

## GET /UserSession/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the UserSession type from the HTTP API returning received result as a JSON
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


    "user_id":	"w85fh4krizo1l5pio8i8",

    "public_id":	"3n6uij5wwebhpvxly1a76clma9bomj",

    "token":	"szjhvenhf68jkz4ejfhf",

    "expires":	"2017-12-26T18:57:10+01:00"

}
```

## GET /UserSession/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the UserSession type from the HTTP API.

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


    "user_id":	"0t1bic1w92f67k0qs7my",

    "public_id":	"1s2xlny4c76i91d31lfhddgvfgzcfg",

    "token":	"2ikh8nuzespzysmfa5po",

    "expires":	"2017-12-26T18:57:10+01:00"

},{


    "user_id":	"l9c3hgaun0f7qzipkgx3",

    "public_id":	"03y9ytjag8ix7w5tgwt8elkgry7fkt",

    "token":	"ccisp0nctrem4rvostk1",

    "expires":	"2017-12-26T18:57:10+01:00"

}]
```

## PUT /UserSession/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the UserSession type from the HTTP API returning received result as a JSON
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


    "user_id":	"e07oaviv7bf1i8hdo3mc",

    "public_id":	"hemo1v0drt2rs1zjrrqc66kjexq7jj",

    "token":	"cxnzfaz1kjksarupbmc2",

    "expires":	"2017-12-26T18:57:10+01:00"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /UserSession/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the UserSession type from the HTTP API returning received result as a JSON
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

