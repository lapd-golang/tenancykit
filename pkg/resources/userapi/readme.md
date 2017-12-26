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


    "username":	"dolorem_ut_sunt",

    "password":	"erias1hha5exqgkrh57n",

    "password_confirm":	"4ahoa4oshqfmjeftj4ni",

    "twofactor_auth":	true,

    "tenant_id":	"x9ftjlwuuf134uu1d2jw",

    "email":	"bPerkins@Jabbersphere.com"

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


    "user_roles":	null,

    "two_factor_auth":	true,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "tenant_id":	"6ew6259dirl42u0rf4e6",

    "private_id":	"35jj0fi25h6hp0gukbqr4a5rgxilg8",

    "hash":	"slk88g849itre2qfbpad",

    "username":	"sRuiz",

    "email":	"7Simmons@Flipopia.net",

    "public_id":	"p6y7ltdtbkajh7b7golzywz39iwnob"

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


    "email":	"qWillis@Yombu.edu",

    "public_id":	"uepnz5qsipjnhifazeqthb91ijp16g",

    "tenant_id":	"tkjctz3gfjky1h8gdqks",

    "user_roles":	null,

    "two_factor_auth":	true,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "username":	"est_dolor",

    "private_id":	"d6ahcn5z9pua5vdzg95ayrio6qfy1z",

    "hash":	"4u1g9nh1m3sa2nu6h0ym"

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


    "user_roles":	null,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "username":	"CarolynHarrison",

    "tenant_id":	"ucr56hfsogyojipw54a4",

    "private_id":	"v64wsalail5iwdgq2yleyufh8o7j1u",

    "two_factor_auth":	true,

    "email":	"CarlCrawford@Voolia.biz",

    "public_id":	"bq1eza5rne7u3t13ol8ec3saqog2cw",

    "hash":	"dquirze8gyamtq9xhb7m"

},{


    "username":	"9Palmer",

    "hash":	"88rt043xv3ipdt4ttg8h",

    "user_roles":	null,

    "two_factor_auth":	true,

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "email":	"hStanley@Tagcat.edu",

    "public_id":	"w6qklxpknykwqbbjrmt3xpsh11q9jf",

    "tenant_id":	"fdwl9wvbep8b5dxf4jud",

    "private_id":	"fetq76q8ejlmxtzrvbbk9fslegtbra",

    "created_at":	"2017-12-26T18:57:10+01:00"

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


    "public_id":	"thaegone10pamboiwfu0bbdd8ftwij",

    "tenant_id":	"ah1pjnmcg1m01rmraykk",

    "hash":	"2s4gf6q5rqu8gjwsdqwd",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "created_at":	"2017-12-26T18:57:10+01:00",

    "username":	"nAustin",

    "email":	"JudithWilliams@Dynabox.mil",

    "private_id":	"1ifig78ae4bom560cs2jo9x4cyfs9l",

    "user_roles":	null,

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

