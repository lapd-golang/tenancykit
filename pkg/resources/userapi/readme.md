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


    "email":	"JustinYoung@Realpoint.org",

    "username":	"placeat",

    "password":	"413de0iteiijteeseccy",

    "password_confirm":	"8ze8c6axrp1wccs0bhow",

    "twofactor_auth":	true,

    "tenant_id":	"l6xcpgar2l0hna5v22gd"

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


    "tenant_id":	"84qfvegm9xet3op457e6",

    "private_id":	"1y9dw1lk3gcbriia5lwxz8jhfodazb",

    "hash":	"2qpjlizv15371ykuc1or",

    "two_factor_auth":	true,

    "email":	"error_consequatur_necessitatibus@Kwilith.biz",

    "public_id":	"bx8rqtwjn5by4839d1505el2nuw97k",

    "updated_at":	"2017-12-23 12:57:49.408855 +0000 UTC",

    "username":	"et",

    "created_at":	"2017-12-23 12:57:49.408847 +0000 UTC"

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


    "hash":	"0rwe40paz3gqbvauogn5",

    "two_factor_auth":	true,

    "username":	"4Graham",

    "public_id":	"rndvalh2ekrcp7nwrc7gjapbthzzv9",

    "private_id":	"8tuc8wzxa5pyecnopksq91zlgjuxrh",

    "created_at":	"2017-12-23 12:57:49.40945 +0000 UTC",

    "updated_at":	"2017-12-23 12:57:49.409456 +0000 UTC",

    "email":	"2Burton@Browsetype.name",

    "tenant_id":	"m9znyqjlqu1njvsn8mt7"

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


    "username":	"FrankWillis",

    "tenant_id":	"dd6sp8e287z29guxf4zc",

    "private_id":	"15s1si8ff3jiy441f889xml5r0xs2i",

    "hash":	"1s1b51240jrhg3ggg17i",

    "created_at":	"2017-12-23 12:57:49.409958 +0000 UTC",

    "email":	"tBailey@Avamm.gov",

    "public_id":	"xnyxv0iba6acb6lohxwoiuhtmd8po9",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-23 12:57:49.409963 +0000 UTC"

},{


    "public_id":	"xxmjmafdthwii7onhzsfdxe7o06q5w",

    "private_id":	"ldkli273sgqbievxpcauypguu4r0vl",

    "hash":	"kuqhukjuepzip698el6t",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-23 12:57:49.410411 +0000 UTC",

    "username":	"xBrown",

    "email":	"molestias_adipisci_quo@Skimia.info",

    "tenant_id":	"190msz0oxjjxol88m2od",

    "created_at":	"2017-12-23 12:57:49.410406 +0000 UTC"

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


    "email":	"AlbertHarrison@Aimbo.info",

    "password":	"uv8z47aikb35jpae2cq4",

    "password_confirm":	"neln3bqjhdc3ya16sgf6"

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

