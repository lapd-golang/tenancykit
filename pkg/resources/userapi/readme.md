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


    "twofactor_auth":	true,

    "tenant_id":	"0ptt8ptfc1hg4x4kbxbb",

    "email":	"MarieMiller@Edgeify.com",

    "username":	"MarieBurke",

    "password":	"fatc0icrd99wb4c70mti",

    "password_confirm":	"g21d1l6ecn8gux01q32o"

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


    "username":	"tempore",

    "public_id":	"7g06ubflozya077l2w462khk3c8r8b",

    "created_at":	"2017-12-24 18:16:09.286537 +0000 UTC",

    "email":	"AshleyMedina@Lazz.mil",

    "tenant_id":	"ffuduniyvs8ojcbnc8qr",

    "private_id":	"wbkoc3pub6i9dpgd1gvid9vtid5pye",

    "hash":	"hrgdz4agncxafcyqwk5l",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-24 18:16:09.286545 +0000 UTC"

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


    "public_id":	"5gjsrmj59shw6z3ez07ymkxoxqnyc3",

    "private_id":	"qb7b704iidmntqdcs24apspjyvn0zk",

    "two_factor_auth":	true,

    "created_at":	"2017-12-24 18:16:09.287025 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.28703 +0000 UTC",

    "username":	"JonathanWest",

    "email":	"JohnnyLittle@Teklist.mil",

    "tenant_id":	"9gk4y5ccnnh9mvu9lqd8",

    "hash":	"ych91twbsjrds569x7s9"

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


    "email":	"CarlosWatkins@Katz.name",

    "public_id":	"ci1r2opo36my27kojc21g4t2ugkwzm",

    "hash":	"gyoxxwi24j8b7q5839zx",

    "updated_at":	"2017-12-24 18:16:09.287633 +0000 UTC",

    "username":	"dolorem",

    "tenant_id":	"ydtegfrblg8ls0wnbxdq",

    "private_id":	"73yzh3pss4s6craqdswa3xks61z6ma",

    "two_factor_auth":	true,

    "created_at":	"2017-12-24 18:16:09.287627 +0000 UTC"

},{


    "email":	"LillianButler@Photojam.org",

    "created_at":	"2017-12-24 18:16:09.288047 +0000 UTC",

    "username":	"et_vero_perspiciatis",

    "public_id":	"p8zmhx29gxi00995xw6e56ggf65zlf",

    "tenant_id":	"13olr9gl536ffe30p2qk",

    "private_id":	"cnjxbm4s2d31uvfq23ozvcroxv42e7",

    "hash":	"9qlhquejyrs0zvyyqyvi",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-24 18:16:09.288052 +0000 UTC"

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


    "email":	"uMatthews@JumpXS.net",

    "password":	"pnus7c970udwtowli1r3",

    "password_confirm":	"l324g6cb1h5hn8webs5z"

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

