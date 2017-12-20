User HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/userapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/userapi)

User HTTP API is a auto generated http api for the struct `User`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.User, error)
    Update(context.Context, string, pkg.UpdateUser) error
    GetAll(context.Context, string, string, int, int) ([]pkg.User, int, error)
    Create(context.Context, pkg.CreateUser) (pkg.User, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /User/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the CreateUser type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "tenant_id":	"cqpfa9smisg3wfz80ixx",

    "email":	"2Barnes@Katz.biz",

    "username":	"JohnnyMills",

    "password":	"hm5x76pprwou1civdtbn",

    "password_confirm":	"shq4lm1xt9l3a8inp6vh"

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


    "email":	"rerum@Buzzshare.mil",

    "tenant_id":	"60iw30jpei7n5ibrcs39",

    "private_id":	"lsf7gt5iht13boop3kmo5vsknzc0qc",

    "hash":	"jf1d56cqwa2yqe07crvu",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 19:23:27.248052 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.248072 +0000 UTC",

    "username":	"oDunn",

    "public_id":	"rav52iz9ffjekqdvn7qnsyiz9g0ux1"

}
```

## GET /User/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the User type from the HTTP API returning received result as a JSON
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


    "public_id":	"iz3qwqcsns7xzttwqrmg1nl6eo14sb",

    "private_id":	"jhyjo238lvp6xlslo6cotpir9zqjx7",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 19:23:27.248631 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.248637 +0000 UTC",

    "username":	"fugit_beatae",

    "email":	"BenjaminCarter@Devify.info",

    "tenant_id":	"lv2wqsefbnj0viztgpgu",

    "hash":	"7ozjtpt9q7zotri2kb8p"

}
```

## GET /User/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the User type from the HTTP API.

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


    "username":	"PeterBishop",

    "tenant_id":	"y7g4v34pk6c65dv403qq",

    "created_at":	"2017-12-20 19:23:27.249179 +0000 UTC",

    "hash":	"nm5j9bqog8p5bo3a2cg5",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-20 19:23:27.249187 +0000 UTC",

    "email":	"aut@Devpoint.com",

    "public_id":	"zgxm50rqpdobjc2za1hm3u1meqqbc8",

    "private_id":	"8o5tq2x0hg87aqz44ouuyomsucdpr9"

},{


    "email":	"illo_iste@Blogspan.org",

    "public_id":	"knodtn53po04u643zs23s9y9v7fe68",

    "tenant_id":	"i7bv92ufbbxvm7z6y97s",

    "hash":	"osy1mmnp3ogg1s2f46oy",

    "username":	"quos_quibusdam_voluptas",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 19:23:27.249703 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.249708 +0000 UTC",

    "private_id":	"xt1wpkv257ed73ht3n08xqi3fdvegu"

}]
```

## PUT /User/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the User type from the HTTP API returning received result as a JSON
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


    "email":	"IreneSanders@Tagcat.name",

    "password":	"nluhjt05eo1vdqxqk4i2",

    "password_confirm":	"ebk5zkkd50bgd0kpv080",

    "is_password_update":	true

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /User/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the User type from the HTTP API returning received result as a JSON
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

