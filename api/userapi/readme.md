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


    "email":	"vitae_iste@Browsetype.mil",

    "username":	"EricFox",

    "password":	"alnw4n9uh0p1ezn89wda",

    "password_confirm":	"ema39t388mribbepy70f",

    "tenant_id":	"tmruuhsxgdi7phaep0hd"

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


    "private_id":	"95wkfdahs7pyfy9bf6x0zjmwic0rzm",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 10:14:31.972344 +0000 UTC",

    "username":	"8Harper",

    "public_id":	"frutsfac2t6a1yztq0r9s50suum02w",

    "hash":	"wpti2opl6qqag8527059",

    "updated_at":	"2017-12-20 10:14:31.972352 +0000 UTC",

    "email":	"sequi_iure@Edgeify.gov",

    "tenant_id":	"l2os602c9akcb7k5rpzs"

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


    "public_id":	"6jvmuekh5nsowfzuyp2k55jtz8a2td",

    "private_id":	"0ras7ee48ykb70tvr9gdzwduk832qj",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 10:14:31.97296 +0000 UTC",

    "username":	"rWilson",

    "tenant_id":	"cw8ee1qyv5pk8mojgtzk",

    "hash":	"624o1rcwitm143u6oy4a",

    "updated_at":	"2017-12-20 10:14:31.972965 +0000 UTC",

    "email":	"voluptatem_ad@Linklinks.edu"

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


    "username":	"aut",

    "email":	"lWood@Chatterpoint.gov",

    "public_id":	"81kxzfc8f9mynfwt5j6ugo0yxhjrii",

    "hash":	"2cow0njq3wpff2do31j5",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 10:14:31.973423 +0000 UTC",

    "tenant_id":	"phn8qby1xu26dozhuskr",

    "private_id":	"rgvkdhsl52xit73aivka5450rv36lk",

    "updated_at":	"2017-12-20 10:14:31.973428 +0000 UTC"

},{


    "email":	"eum_eum_quis@Brightdog.org",

    "two_factor_auth":	true,

    "private_id":	"bbhat1omrd5gj330sb4sqcvtrdm870",

    "hash":	"xuki6z6caw6ekhagllbf",

    "created_at":	"2017-12-20 10:14:31.973861 +0000 UTC",

    "updated_at":	"2017-12-20 10:14:31.973865 +0000 UTC",

    "username":	"fGardner",

    "public_id":	"nfikks9lp9apq8cg697wbpk17f65gd",

    "tenant_id":	"kec71rr0dxmrzqust6ek"

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


    "email":	"CynthiaMartin@Kwinu.mil",

    "password":	"5i7mlbmrq4w7inuo9yvi",

    "password_confirm":	"y6nzx75o9ruddjotyy6f",

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

