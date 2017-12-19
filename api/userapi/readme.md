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

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "tenant_id":	"lxchsy6egtyvh5weoegd",

    "email":	"bRoberts@Gigabox.org",

    "username":	"CatherineGutierrez",

    "password":	"trltdttqhcq089mlhh2u",

    "password_confirm":	"eky7d8owe9tlobj9tif8"

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


    "public_id":	"mfhez29dxtl5hw1w7t9d5nd14bsnvh",

    "hash":	"90yfz1a8djzlp4hwppw8",

    "created_at":	"2017-12-19 18:15:32.547067 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.54708 +0000 UTC",

    "username":	"7Perry",

    "email":	"PaulaTucker@Jaxbean.net",

    "tenant_id":	"xwgc6u3z5ctzsw8kqr47",

    "private_id":	"478whexu9y0oribpw72vkc4bwwoxyl",

    "two_factor_auth":	true

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


    "username":	"recusandae_ullam",

    "email":	"libero_ut@Quire.org",

    "private_id":	"mu02ast718l2vm3saweokpiladhlcn",

    "two_factor_auth":	true,

    "created_at":	"2017-12-19 18:15:32.563277 +0000 UTC",

    "public_id":	"15g9ynvbyp5ofv073w3a9snl6hy1bd",

    "tenant_id":	"rl9sujfqyuiueefd1v07",

    "hash":	"4ww47ytpecj6igowiegn",

    "updated_at":	"2017-12-19 18:15:32.563289 +0000 UTC"

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


    "email":	"gAustin@Zoomzone.edu",

    "tenant_id":	"mqmwyj8cv65a99juvlx9",

    "hash":	"qu0bpo2g379ixh0bo2co",

    "username":	"EricHarper",

    "private_id":	"dxso54qszulenoxs4s9f6bqnn00wes",

    "two_factor_auth":	true,

    "created_at":	"2017-12-19 18:15:32.565097 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.565108 +0000 UTC",

    "public_id":	"s1c9pnedtt7zki68t6cx9vxx0uq8q2"

},{


    "username":	"EdwardHunt",

    "email":	"cGonzales@Yodo.biz",

    "public_id":	"a08p46biseilolkeqy9omsaz6vfi3p",

    "hash":	"2hfcgamtrxixoqoafo9l",

    "updated_at":	"2017-12-19 18:15:32.565707 +0000 UTC",

    "tenant_id":	"ftpj1k2qweorb4xwgmo3",

    "private_id":	"otfn5yjwsxdnei8lakbskuf7cufrka",

    "two_factor_auth":	true,

    "created_at":	"2017-12-19 18:15:32.565698 +0000 UTC"

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


    "password":	"4lei2ywwwktg4tuu3697",

    "password_confirm":	"ys997bb49t2wqryp05ht",

    "is_password_update":	true,

    "email":	"lStevens@Eayo.info"

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

