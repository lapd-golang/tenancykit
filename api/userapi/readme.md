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


    "email":	"EdwardSimmons@Pixoboo.edu",

    "username":	"quaerat",

    "password":	"ukpyzrgz72eqxif9uwh1",

    "password_confirm":	"b8d0bw5rdvols9ljvome",

    "tenant_id":	"lr7gma9y5gg9rvlxp60o"

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


    "email":	"GregoryEdwards@Yadel.org",

    "public_id":	"cwqdm3343n7jx87vgv2ipt1u6exodj",

    "tenant_id":	"fsxvbz2glwk0tqgcno6l",

    "private_id":	"f2fv5c45ogbnycwz80mc3s8ufyczmf",

    "hash":	"nu3gkrl4bflgg6c00832",

    "created_at":	"2017-12-20 12:28:53.773386 +0000 UTC",

    "updated_at":	"2017-12-20 12:28:53.773396 +0000 UTC",

    "username":	"CarolynEvans",

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


    "public_id":	"cyaru6xhpzrgj62szcz0bor82fb8pg",

    "tenant_id":	"g6vzeg0ij5infecj3vls",

    "hash":	"gxizw4ifzuds4zbpdotx",

    "created_at":	"2017-12-20 12:28:53.77489 +0000 UTC",

    "updated_at":	"2017-12-20 12:28:53.7749 +0000 UTC",

    "username":	"aut_non",

    "email":	"iMorris@Zoomdog.edu",

    "private_id":	"7mwjb6onc14hvz8fmrvzyxqxpgrxkp",

    "two_factor_auth":	true

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


    "username":	"7Jacobs",

    "email":	"0Reed@Fivebridge.net",

    "public_id":	"7e9rd25ke19b7n9ov64nr91nilyiuk",

    "tenant_id":	"9aaustjuq6xttweo7eui",

    "private_id":	"01gggswa0edzjmwa3as4ywrit8c9hb",

    "hash":	"amh8s5qbbfkovrhljokf",

    "created_at":	"2017-12-20 12:28:53.775842 +0000 UTC",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-20 12:28:53.775848 +0000 UTC"

},{


    "public_id":	"ytxplh7wgl0crxly0f25yz4puei7kw",

    "private_id":	"de7f2h5w9hjd9pchu3p9pz9t9dbcgj",

    "hash":	"kg7agx8wkt7vtfn70a88",

    "created_at":	"2017-12-20 12:28:53.77629 +0000 UTC",

    "username":	"ScottWoods",

    "email":	"sunt_voluptatem_sint@Shuffledrive.name",

    "tenant_id":	"9b0e07iq9n2kc2n39hd2",

    "two_factor_auth":	true,

    "updated_at":	"2017-12-20 12:28:53.776295 +0000 UTC"

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


    "email":	"aut@Mydo.edu",

    "password":	"94fjbyqvoi89gdhq59qi",

    "password_confirm":	"4wo46m6jtlrdcc5ab6up",

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

