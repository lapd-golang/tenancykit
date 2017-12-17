Token HTTP API 
===============================

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


    "value":	"",

    "public_id":	"",

    "target_id":	""

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


    "value":	"",

    "public_id":	"",

    "target_id":	""

}
```

## GET /Token/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the Token type from the HTTP API returning received result as a JSON
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


    "target_id":	"",

    "value":	"",

    "public_id":	""

}
```

## GET /Token/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the Token type from the HTTP API.

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


    "public_id":	"",

    "target_id":	"",

    "value":	""

}]
```

## PUT /Token/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the Token type from the HTTP API returning received result as a JSON
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


    "value":	"",

    "public_id":	"",

    "target_id":	""

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

## DELETE /Token/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the Token type from the HTTP API returning received result as a JSON
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