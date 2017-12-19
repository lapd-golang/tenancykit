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


    "email":	"fRichardson@Wikizz.gov",

    "username":	"s",

    "password":	"7",

    "password_confirm":	"s",

    "tenant_id":	"k"

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


    "email":	"voluptatibus_aliquam_sit@Oloo.net",

    "tenant_id":	"h",

    "created_at":	2017-12-19 13:42:58.853141 +0000 UTC,

    "username":	"q",

    "public_id":	"ue8bd6ehw79viy3",

    "private_id":	"4cf6n109jqmbhr5",

    "hash":	"w",

    "two_factor_auth":	true,

    "updated_at":	2017-12-19 13:42:58.85315 +0000 UTC

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


    "created_at":	2017-12-19 13:42:58.85442 +0000 UTC,

    "email":	"MatthewDuncan@Rhycero.info",

    "public_id":	"amtsj2vc3pzqknh",

    "tenant_id":	"a",

    "private_id":	"zfmogpd2wqow1i1",

    "hash":	"4",

    "two_factor_auth":	true,

    "updated_at":	2017-12-19 13:42:58.854425 +0000 UTC,

    "username":	"s"

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


    "updated_at":	2017-12-19 13:42:58.854973 +0000 UTC,

    "username":	"8",

    "email":	"doloremque_corporis@Podcat.name",

    "public_id":	"e2wassth6rseolw",

    "tenant_id":	"t",

    "private_id":	"prgj2286cudhin1",

    "hash":	"s",

    "created_at":	2017-12-19 13:42:58.854968 +0000 UTC,

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


    "email":	"VirginiaMcdonald@Centimia.info",

    "password":	"v",

    "password_confirm":	"q",

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