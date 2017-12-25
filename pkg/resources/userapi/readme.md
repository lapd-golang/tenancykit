User HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/userapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/userapi)

User HTTP API is a auto generated http api for the struct `User`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.User, error)
    Update(context.Context, string, pkg.User) error
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


    "tenant_id":	"5uth9dpedw99wlqoo0ti",

    "email":	"StevenPhillips@Abatz.edu",

    "username":	"at_sed_minus",

    "password":	"8jpfff27iyk1j8d74hq2",

    "password_confirm":	"0ws8usxpigcvuyxmd51n",

    "twofactor_auth":	true

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


    "created_at":	"2017-12-25 16:32:45.880277 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.880312 +0000 UTC",

    "username":	"WalterPatterson",

    "public_id":	"un2vl7s4dmhgosty5o9z89t6w2zho9",

    "tenant_id":	"q2t7ee5x97wgtc03pu38",

    "hash":	"qgcyig3obxui1ycgj49z",

    "email":	"IreneTucker@Realbridge.mil",

    "private_id":	"j453ukn6oorie9oq0bd6tbw36tvfcf",

    "user_roles":	null,

    "two_factor_auth":	true

}
```

## INFO /User/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.User.

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


    "created_at":	"2017-12-25 16:32:45.881088 +0000 UTC",

    "username":	"dolor_rerum_sit",

    "hash":	"jhj57omwuunadf9xtrzf",

    "user_roles":	null,

    "private_id":	"owzgqc13bftu95cx630gbv6zzclpn8",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-25 16:32:45.881097 +0000 UTC",

    "email":	"bParker@Gabcube.com",

    "public_id":	"veyma831l2qwjit7id7k2g8jf1941v",

    "tenant_id":	"b7fj7exsoo5d6ic5lzdc"

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


    "username":	"tGarcia",

    "public_id":	"dr0riwf2hpylulm3fyvh6tt0ls2mbt",

    "tenant_id":	"f43sevxaqwolm5k91wpz",

    "private_id":	"207ppbma7yifjotpm2emfaszofmsst",

    "hash":	"vsvauwbd4t5ippfh1gm1",

    "created_at":	"2017-12-25 16:32:45.88184 +0000 UTC",

    "email":	"EmilySanders@Divape.edu",

    "user_roles":	null,

    "two_factor_auth":	true,

    "updated_at":	"2017-12-25 16:32:45.881854 +0000 UTC"

},{


    "username":	"yCarter",

    "email":	"JonathanGomez@Zoonder.name",

    "public_id":	"h7ttlfg82nujzkxbsqf33qev2z8dmd",

    "tenant_id":	"bhglgzfidby35uhqa8au",

    "private_id":	"917baj32txyzwfbx0ztt58auup89q9",

    "hash":	"40gq1gp7oy1tkrxognlf",

    "user_roles":	null,

    "two_factor_auth":	true,

    "created_at":	"2017-12-25 16:32:45.882783 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.882798 +0000 UTC"

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


    "private_id":	"q4wjp4c0g67601q7arwro7zhlcm4r4",

    "hash":	"e49q4m0ezcvaznk9w1ng",

    "user_roles":	null,

    "created_at":	"2017-12-25 16:32:45.883921 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:45.88394 +0000 UTC",

    "username":	"LarryAllen",

    "email":	"ScottRamos@Mymm.info",

    "tenant_id":	"dq7qi63v3ozfey6z4ghk",

    "public_id":	"1duzp44o3czo94aktfg0xdl80wocw6",

    "two_factor_auth":	true

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

