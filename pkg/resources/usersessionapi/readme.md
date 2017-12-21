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


    "email":	"autem_non@Abatz.mil",

    "password":	"23pk1mktu7l0r2u31tj8",

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


    "token":	"vflvscxv23fh0onhoktx",

    "expires":	"2017-12-21 13:40:49.270925 +0000 UTC",

    "user_id":	"pu39dbdqygbft1pz4rmb",

    "public_id":	"ogj7lg00t6f5642r2ts4yp2w3c16fy"

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


    "token":	"q4113re28rc17ki2q0j0",

    "expires":	"2017-12-21 13:40:49.271923 +0000 UTC",

    "user_id":	"1yg68ff1hj3gyzslllq1",

    "public_id":	"h5fqncxjkh4463xhmebp7i1sxuclwi"

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


    "user_id":	"e5fyqeqyvzg15in1uct4",

    "public_id":	"l5ntd56e8bg4dtw2mgkkuv3seyalvp",

    "token":	"ocmx8lg6uj8gbq2fvp37",

    "expires":	"2017-12-21 13:40:49.272704 +0000 UTC"

},{


    "user_id":	"6v9b4j4g3uaxa9pg5fil",

    "public_id":	"ocz5gqqs26wzm6x8w45zv6570tr8i7",

    "token":	"4q65cfbonky5vgh68rbq",

    "expires":	"2017-12-21 13:40:49.273086 +0000 UTC"

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


    "user_id":	"vl6ncnqcjv3h1in71ruo",

    "public_id":	"t6xdf15jbd52sb1m9u02i6bqh8kcu7",

    "token":	"2zanyhjvdcpcy0tusohz",

    "expires":	"2017-12-21 13:40:49.273413 +0000 UTC"

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

