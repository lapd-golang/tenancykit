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


    "tenant_id":	"ohl55aalba3zm6qxg8sv",

    "email":	"lHanson@Yakidoo.edu",

    "username":	"DeniseElliott",

    "password":	"42fkqnuhkfk43mrwh5a6",

    "password_confirm":	"hmzzk7qhu8pu323smjp5"

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


    "hash":	"zbndk0uzotuzjk8l250o",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 15:06:24.819844 +0000 UTC",

    "private_id":	"8zhpxwwl6y0brc7w1qx6oxq8fcqqxk",

    "updated_at":	"2017-12-20 15:06:24.819853 +0000 UTC",

    "username":	"iure_adipisci_tempore",

    "email":	"JosephFox@Flashdog.info",

    "public_id":	"ylvoy0uls7djg25zf153bq67avrplp",

    "tenant_id":	"7bc85ucbugu6blg1cnxw"

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


    "username":	"oCastillo",

    "tenant_id":	"rcwai7789mgsd4tjwm27",

    "private_id":	"2tj4f567wl3xk6ncua56jj5xcm2rqv",

    "created_at":	"2017-12-20 15:06:24.820397 +0000 UTC",

    "email":	"iure@Tagcat.info",

    "public_id":	"y0t1oi8ibxs36u5wh5xzywd65ikiqt",

    "hash":	"5q84hz2wop1nz5x6o1kw",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-20 15:06:24.820403 +0000 UTC"

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


    "created_at":	"2017-12-20 15:06:24.820901 +0000 UTC",

    "username":	"lGarrett",

    "email":	"SarahFlores@Yabox.info",

    "public_id":	"o3t5hauk7kdakwoga8xl28rw73d8c2",

    "private_id":	"k29n7n4bq7g901p1pvjkgc5dwrtess",

    "hash":	"kxkqpd9rct8chf9tscms",

    "tenant_id":	"1i2gzv7ufvv2ihgjngod",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-20 15:06:24.820907 +0000 UTC"

},{


    "two_factor_auth":	true,

    "updated_at":	"2017-12-20 15:06:24.821435 +0000 UTC",

    "username":	"veritatis",

    "email":	"9Sullivan@Voolia.org",

    "private_id":	"9xbmvwza4kki635y85s85rpdjgfduk",

    "created_at":	"2017-12-20 15:06:24.82143 +0000 UTC",

    "public_id":	"frt9099r0hljoxwk794av08iltj71y",

    "tenant_id":	"2qzfjpkyj6d8h6muah4q",

    "hash":	"oanxhw5uk1ilkdrhl14u"

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


    "email":	"cWest@Zoomdog.gov",

    "password":	"opzwl7zqgdy33ij7ypgo",

    "password_confirm":	"fkwtpxspxxgdokj0f7qj",

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

