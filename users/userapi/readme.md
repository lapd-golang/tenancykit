User HTTP API 
===============================

User HTTP API is a auto generated http api for the struct `User`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (users.User, error)
    Update(context.Context, string, users.User) error
    GetAll(context.Context, string, string, int, int) ([]users.User, int, error)
    Create(context.Context, users.User) (users.User, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /User/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the User type which is delieved the 
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


    "two_factor_auth":	false,

    "created_at":	"2017-12-15 21:20:08.177069 +0000 UTC",

    "updated_at":	"2017-12-15 21:20:08.177092 +0000 UTC",

    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	""

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


    "updated_at":	"2017-12-15 21:20:08.177502 +0000 UTC",

    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-15 21:20:08.177496 +0000 UTC"

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


    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-15 21:20:08.17788 +0000 UTC",

    "updated_at":	"2017-12-15 21:20:08.177885 +0000 UTC",

    "username":	""

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


    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-15 21:20:08.17821 +0000 UTC",

    "updated_at":	"2017-12-15 21:20:08.178215 +0000 UTC"

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


    "username":	"",

    "public_id":	"",

    "private_id":	"",

    "hash":	"",

    "two_factor_auth":	false,

    "created_at":	"2017-12-15 21:20:08.178511 +0000 UTC",

    "updated_at":	"2017-12-15 21:20:08.178516 +0000 UTC"

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