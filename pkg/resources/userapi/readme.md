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


    "tenant_id":	"9mh9jyo45twl2yw2s50p",

    "email":	"7Graham@Vipe.edu",

    "username":	"DeborahJackson",

    "password":	"7nagr71c324tzi8om93c",

    "password_confirm":	"ezbdajblavv6czrc0nu6"

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


    "public_id":	"yf6g1ej1qkdcxsvpc6drh884mcioa6",

    "tenant_id":	"3i381th36adce0sm1r9b",

    "private_id":	"b4zrexcr4j7a3598vz1v0haebhkwor",

    "hash":	"s7qi7kn92di3dd96m90i",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-21 20:42:41.430941 +0000 UTC",

    "username":	"0Stone",

    "created_at":	"2017-12-21 20:42:41.430933 +0000 UTC",

    "email":	"dolorum_quo@Wordify.info"

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


    "tenant_id":	"vwmde7c82uo58nuz33en",

    "hash":	"l5v7bvu0dgpjp8642tvm",

    "created_at":	"2017-12-21 20:42:41.431638 +0000 UTC",

    "username":	"illum_mollitia",

    "public_id":	"ui1g30v1d2z99lydeonl7zkr3xq0cc",

    "private_id":	"d3t0vf78i7lbdyl7b97yruoa4bga2d",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-21 20:42:41.431646 +0000 UTC",

    "email":	"CharlesEdwards@Lazz.mil"

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


    "email":	"DonaldWatkins@Lazz.name",

    "public_id":	"29qfdb8gcdhtr78biisyqquz9ouesm",

    "hash":	"zu5kb4d6q8d1io5tscvn",

    "two_factor_auth":	true,

    "created_at":	"2017-12-21 20:42:41.432245 +0000 UTC",

    "username":	"yHarper",

    "tenant_id":	"hgactspddl4yqa59fu55",

    "private_id":	"hvkbout4q9v0yqus88bmw9nisttz8u",

    "updated_at":	"2017-12-21 20:42:41.43225 +0000 UTC"

},{


    "username":	"RandySims",

    "email":	"SarahJones@Dynabox.gov",

    "tenant_id":	"1otc0g3eud7j186fl8fl",

    "hash":	"lcx99un9e31v0gxd7bfb",

    "created_at":	"2017-12-21 20:42:41.432763 +0000 UTC",

    "updated_at":	"2017-12-21 20:42:41.432768 +0000 UTC",

    "public_id":	"vp56gaedhdw2v5jv9eqet8qcefihzr",

    "private_id":	"gz9cfg6vqhvg87gztg8p427rlgx8g5",

    "two_factor_auth":	true

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


    "email":	"3Black@Twitterlist.com",

    "password":	"b9ysw244esr78zlg51so",

    "password_confirm":	"joe5is4c6r4po8icrvbn"

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

