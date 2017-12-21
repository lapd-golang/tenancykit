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


    "username":	"EricBanks",

    "password":	"7pt1xcnf6fdm12o1bdkf",

    "password_confirm":	"i7e532f4ubr5zye6ji7f",

    "tenant_id":	"3t4bgcykpt783umatv8x",

    "email":	"hWilliamson@Zoovu.biz"

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


    "created_at":	"2017-12-21 13:40:49.317101 +0000 UTC",

    "tenant_id":	"8sfus2bf0k6u5mblolz9",

    "private_id":	"9axdtuslfzn63l3lwy1jp3czmzdtjv",

    "hash":	"x9fysbty6o6mubyca9cl",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-21 13:40:49.31711 +0000 UTC",

    "username":	"pRivera",

    "email":	"9Woods@Trupe.biz",

    "public_id":	"usyjumpoef991zpppml66ddt75psem"

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


    "hash":	"fx4ebwsuz8ygfey36s5q",

    "two_factor_auth":	true,

    "created_at":	"2017-12-21 13:40:49.317598 +0000 UTC",

    "username":	"SaraStephens",

    "public_id":	"spi79qozwj7icr3g2q1whak1geswp3",

    "tenant_id":	"1ej3cdkpznr2rzkciuwp",

    "private_id":	"krf9qb4sfh5f19z01e1tvxd98m7vzc",

    "email":	"zCole@Edgetag.com",

    "updated_at":	"2017-12-21 13:40:49.317603 +0000 UTC"

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


    "username":	"nostrum_id",

    "email":	"BrendaHamilton@Yacero.org",

    "public_id":	"s5ymepzhe1oyg7uctpi8e3cyesvq2g",

    "hash":	"7js1trua868t21esfxup",

    "two_factor_auth":	true,

    "tenant_id":	"thwc68ot885qit53q1e6",

    "private_id":	"wty1oq5xmy12ltiv3eg58rppfh40x0",

    "created_at":	"2017-12-21 13:40:49.318076 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.318081 +0000 UTC"

},{


    "created_at":	"2017-12-21 13:40:49.318457 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.318462 +0000 UTC",

    "email":	"zReynolds@Jetpulse.mil",

    "tenant_id":	"bd7dgrde3wrdhse1sl16",

    "two_factor_auth":	true,

    "hash":	"ckyxczf3cl2muxn56ou7",

    "username":	"est",

    "public_id":	"97o8nvvwqikul7b0tugs6nqgqpsx3w",

    "private_id":	"qwulcllbog7th45c4426hitftcuxxs"

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


    "email":	"ea_ratione@Flashdog.mil",

    "password":	"s13hm79uknb20959uvqa",

    "password_confirm":	"eko15vzdzk3z2dl5yzsx"

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

