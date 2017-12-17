User HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/userapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/userapi)

User HTTP API is a auto generated http api for the struct `User`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.User, error)
    Update(context.Context, string, tenancykit.UpdateUser) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.User, int, error)
    Create(context.Context, tenancykit.CreateUser) (tenancykit.User, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /User/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the CreateUser type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
```

- Expected Request Body

```json
{


    "password":	"",

    "password_confirm":	"",

    "tenant_id":	"",

    "email":	"",

    "username":	""

}
```

- Expected Status Code

```
Failure: 500
Success: 201
```

- Expected Response Body

```json
{


    "email":	"",

    "tenant_id":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-17 19:11:25.423986 +0000 UTC",

    "updated_at":	"2017-12-17 19:11:25.423995 +0000 UTC",

    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	""

}
```

## GET /User/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the User type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```json
{


    "tenant_id":	"",

    "email":	"",

    "username":	"",

    "password":	"",

    "password_confirm":	""

}
```

## GET /User/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the User type from the HTTP API.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```json
[{


    "password_confirm":	"",

    "tenant_id":	"",

    "email":	"",

    "username":	"",

    "password":	""

}]
```

## PUT /User/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the User type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record with the provided JSON request body.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
{


    "email":	"",

    "password":	"",

    "password_confirm":	""

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```


- Expected Response Body

```json
```

## DELETE /User/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the User type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Content Type: 

```http
    Content-Type: application/json
```

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```json
```

- Expected Status Code

```
Failure: 500
Success: 204
```

- Expected Response Body

```json
```