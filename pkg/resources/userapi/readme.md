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


    "tenant_id":	"xb94mtwnqt04ku7ucw7p",

    "email":	"hGordon@Dablist.net",

    "username":	"LouisStewart",

    "password":	"b0wb7xkr8grjywoqry1y",

    "password_confirm":	"26gfhn58wp7gzvefqu5o"

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


    "hash":	"4x11rjw2tcun68el27ob",

    "two_factor_auth":	true,

    "created_at":	"2017-12-21 15:13:51.283143 +0000 UTC",

    "username":	"numquam_unde",

    "email":	"et_placeat@Oyoyo.edu",

    "public_id":	"vanmdqqkv5gu7aexkr5rrwys6omfnp",

    "private_id":	"6z6weixdydcvwzs8jwym70y5vfxjtg",

    "tenant_id":	"k0fn71idkg84l2blwz11",

    "updated_at":	"2017-12-21 15:13:51.283153 +0000 UTC"

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


    "hash":	"bzfqe8er0saxgu03i5jq",

    "two_factor_auth":	true,

    "created_at":	"2017-12-21 15:13:51.283903 +0000 UTC",

    "public_id":	"1u0lndpe0490xdd3k7ieer5er67zll",

    "private_id":	"j37su5quqs6efjv3ld8od6aq42lgwh",

    "tenant_id":	"78bjc4w27nio7prs358h",

    "updated_at":	"2017-12-21 15:13:51.28391 +0000 UTC",

    "username":	"SeanStevens",

    "email":	"GaryGriffin@Eare.name"

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


    "hash":	"mnj9lq40xfwj2bq8tm8y",

    "created_at":	"2017-12-21 15:13:51.284529 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:51.284534 +0000 UTC",

    "email":	"MarieMartinez@Flashspan.com",

    "private_id":	"v572pljpcv0juv8p2ue6qmwzqt982y",

    "tenant_id":	"fkga4yaw42reboxywnmy",

    "two_factor_auth":	true,

    "username":	"BettyMontgomery",

    "public_id":	"qvlegyt12vxs1yhnbipdt93fscdnxy"

},{


    "username":	"consequuntur_maxime_perferendis",

    "tenant_id":	"6mjh2s8o7f2cq9jfvpas",

    "two_factor_auth":	true,

    "created_at":	"2017-12-21 15:13:51.28494 +0000 UTC",

    "email":	"3Bennett@Camimbo.name",

    "public_id":	"z4na9av7cg5d3nl3tgpd1xy5nz7j41",

    "private_id":	"4f1h8wmxa3bs4ks3iwwejeip4m97dd",

    "hash":	"q3tzhxj95n3lhogxxpma",

    "updated_at":	"2017-12-21 15:13:51.284944 +0000 UTC"

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


    "email":	"3Peters@Livefish.gov",

    "password":	"craptyzfc1wom9l4o11s",

    "password_confirm":	"tiqthdjeu9bp43qpgvdi"

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

