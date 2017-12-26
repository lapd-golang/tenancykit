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


    "tenant_id":	"k4snl8u2b0pckjpx9xrq",

    "email":	"4Alvarez@Twitterlist.org",

    "username":	"quia",

    "password":	"q0l387vncyyyeif6lv37",

    "password_confirm":	"mwwlutjnzolw82i6pg8p",

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


    "username":	"sCole",

    "hash":	"32jto2ld4lam4l116ntz",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-26 15:18:04.448631 +0000 UTC",

    "user_roles":	null,

    "created_at":	"2017-12-26 15:18:04.448619 +0000 UTC",

    "email":	"TerryBarnes@Edgeify.gov",

    "public_id":	"r95egel2lhulwfbfiio2yivimtqtvj",

    "tenant_id":	"jdyey76ib7q5p2xsp4ea",

    "private_id":	"lu2439tffliote3jcr3fbgiulbe7s9"

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


    "email":	"provident_est@Jayo.gov",

    "tenant_id":	"wff4yxqzgxw4tt9thw7e",

    "created_at":	"2017-12-26 15:18:04.449339 +0000 UTC",

    "user_roles":	null,

    "two_factor_auth":	true,

    "updated_at":	"2017-12-26 15:18:04.449347 +0000 UTC",

    "username":	"EdwardWood",

    "public_id":	"zk7xpnbyeckgxqen7puaovz6yfl20w",

    "private_id":	"y04zyhdyinfbis7hlk87lf583duha1",

    "hash":	"76r8p41yt86308disyns"

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


    "email":	"sunt_libero@Rhyloo.com",

    "tenant_id":	"v9qwaw9cpvh6y3lk6m7p",

    "private_id":	"a4y1qtrnx9r3lw6z5wpj9xapgf42yt",

    "user_roles":	null,

    "two_factor_auth":	true,

    "created_at":	"2017-12-26 15:18:04.449916 +0000 UTC",

    "username":	"RandyCampbell",

    "public_id":	"exu14qnbu9dgp7sev5kk274j0nh59b",

    "hash":	"o13xeph366hsbgu4wllh",

    "updated_at":	"2017-12-26 15:18:04.449924 +0000 UTC"

},{


    "two_factor_auth":	true,

    "created_at":	"2017-12-26 15:18:04.450527 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.450539 +0000 UTC",

    "username":	"MichelleHudson",

    "email":	"JuliaHart@Reallinks.name",

    "tenant_id":	"37vcs0e8zmed7xi4knfu",

    "private_id":	"lm55d742begyss9oei204ewisbj95c",

    "public_id":	"akcm7rzce6u905xpbhfbkq6zb75nzm",

    "hash":	"ui73fd4cdybywnhpxtpc",

    "user_roles":	null

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


    "updated_at":	"2017-12-26 15:18:04.451109 +0000 UTC",

    "tenant_id":	"ww7lyaprrzp1n7f5wgsg",

    "private_id":	"mepya1lbrxmuj256ke6wvzmo8jeegy",

    "hash":	"wzke0q990q9nyde6va2a",

    "user_roles":	null,

    "created_at":	"2017-12-26 15:18:04.451103 +0000 UTC",

    "username":	"cWilson",

    "email":	"nesciunt_aut@Wikido.name",

    "public_id":	"413vz8toqjmb13wvaua00ntt66yvjm",

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

