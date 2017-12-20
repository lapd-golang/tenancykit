UserSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/usersessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/usersessionapi)

UserSession HTTP API is a auto generated http api for the struct `UserSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.UserSession, error)
    Update(context.Context, string, tenancykit.UserSession) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.UserSession, int, error)
    Create(context.Context, tenancykit.CreateUserSession) (tenancykit.UserSession, error)
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


    "email":	"minus@Roomm.name",

    "password":	"qnaqph36z8ecxkpj9tjc",

    "expiration":	nil

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


    "expires":	"2017-12-20 10:14:31.925051 +0000 UTC",

    "user_id":	"kdlg5ohibficdm0emp64",

    "public_id":	"g3u9cqd38fdmsrxqhi0voyepm34ktj",

    "token":	"xlj8e2g3wvrfrj3qfkoi"

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


    "token":	"kxi69dtpxo656zbdpypz",

    "expires":	"2017-12-20 10:14:31.925434 +0000 UTC",

    "user_id":	"32kx3kwzz7oridkowxxq",

    "public_id":	"lsi8wnhramtxaf4kpzxrw5svn37ono"

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


    "public_id":	"xv4a3o2i40e63dfrik991cy0qejg3e",

    "token":	"ptqjtu1xle25aj0ih2dy",

    "expires":	"2017-12-20 10:14:31.925812 +0000 UTC",

    "user_id":	"bbdp1pcjgiip8s5px98r"

},{


    "token":	"bb0dkec7g1kuzkaa2gbi",

    "expires":	"2017-12-20 10:14:31.926106 +0000 UTC",

    "user_id":	"sp1g3wkhbx41hvrn80f1",

    "public_id":	"kk88n8ygecy31uewlvazuihlvfsqny"

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


    "user_id":	"6uv70cf28vdeu0dabvog",

    "public_id":	"rodvcwb1o54t4ss6zwjxqfkd2caed8",

    "token":	"4rhh3rf35isocowpjqke",

    "expires":	"2017-12-20 10:14:31.926418 +0000 UTC"

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

