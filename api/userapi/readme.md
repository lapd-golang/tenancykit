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


    "tenant_id":	"w",

    "email":	"ShirleyBarnes@Aimbo.gov",

    "username":	"d",

    "password":	"w",

    "password_confirm":	"6"

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


    "two_factor_auth":	true,

    "updated_at":	"2017-12-19 14:01:02.739282 +0000 UTC",

    "email":	"1Lane@Jaxnation.gov",

    "public_id":	"8eg3tie365jgg8s",

    "tenant_id":	"e",

    "private_id":	"548i28n9f76wwz3",

    "hash":	"4",

    "username":	"b",

    "created_at":	"2017-12-19 14:01:02.739273 +0000 UTC"

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


    "two_factor_auth":	true,

    "username":	"2",

    "email":	"dArmstrong@Zoonoodle.info",

    "private_id":	"n5b6xzrwrxnxmw9",

    "hash":	"2",

    "public_id":	"5hf49c35rx6f7l2",

    "tenant_id":	"b",

    "created_at":	"2017-12-19 14:01:02.740153 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.74016 +0000 UTC"

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


    "public_id":	"d2vk05ja0z9cp8c",

    "private_id":	"lhw6q57vz5ccyji",

    "hash":	"n",

    "created_at":	"2017-12-19 14:01:02.74071 +0000 UTC",

    "updated_at":	"2017-12-19 14:01:02.740715 +0000 UTC",

    "username":	"w",

    "email":	"sint_praesentium@Chatterbridge.biz",

    "tenant_id":	"y",

    "two_factor_auth":	true

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


    "email":	"yFernandez@Camido.gov",

    "password":	"a",

    "password_confirm":	"e",

    "is_password_update":	true

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