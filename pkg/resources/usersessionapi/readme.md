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


    "email":	"quae_sit@Trilia.mil",

    "password":	"06wthf4m6jxer1mmrn45",

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


    "user_id":	"0kmbrk7v91fy5jn387wb",

    "public_id":	"dmxrsrhzqa73qckgoorb390xzcs093",

    "token":	"q9t7m6pwvojpw6ydfisc",

    "expires":	"2017-12-24 06:38:45.292966 +0000 UTC"

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


    "user_id":	"pf0pqsxex8gek57rlpni",

    "public_id":	"t3ocejyefcz7v3sfozjvvb4y1s3rzf",

    "token":	"75517di4vn1nsn2dflyn",

    "expires":	"2017-12-24 06:38:45.293495 +0000 UTC"

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


    "user_id":	"fzv7ygvfhnmm5zs35dua",

    "public_id":	"uijr12wi2gii9x96af6jquyulwu3l8",

    "token":	"kosgca8r440t9bfyw470",

    "expires":	"2017-12-24 06:38:45.293891 +0000 UTC"

},{


    "user_id":	"cjee1x6t0aal15ele05o",

    "public_id":	"2wyufeos8gwlp8niby91qumj8ar45l",

    "token":	"8kcdllk91nklbmyst9o6",

    "expires":	"2017-12-24 06:38:45.294172 +0000 UTC"

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


    "user_id":	"p2ubexy2jhffyju526lh",

    "public_id":	"ruqg7ydrotr46ekv81urahpqi14du9",

    "token":	"ytn6tnjc7d7q345243rd",

    "expires":	"2017-12-24 06:38:45.294486 +0000 UTC"

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

