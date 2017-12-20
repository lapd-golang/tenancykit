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


    "tenant_id":	"cuartpnolylql5qynrnm",

    "email":	"ratione_esse@Eamia.edu",

    "username":	"voluptatem_quidem_accusamus",

    "password":	"353oxxsbyev18u1nljjo",

    "password_confirm":	"8joc61n3ulq0q5k7b0kj"

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


    "email":	"nisi_libero_laboriosam@Npath.biz",

    "created_at":	"2017-12-20 08:05:05.313612 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.313637 +0000 UTC",

    "username":	"TinaPowell",

    "public_id":	"gvx4gdb9jr1k9kplx3i0v7cx9pyh0f",

    "tenant_id":	"ysrsruj0nxpo5n8xrdmc",

    "private_id":	"xqy9mf9q6e3xmkw0ipnzhawv36udwi",

    "hash":	"5g13sn28ox3x1t0vxvh4",

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


    "email":	"EvelynRose@Camido.com",

    "hash":	"jrarhofgw7kwliqll7n8",

    "two_factor_auth":	true,

    "created_at":	"2017-12-20 08:05:05.315026 +0000 UTC",

    "username":	"RoseHughes",

    "public_id":	"24jc3offd7fofw4g6gftddrzvigziy",

    "tenant_id":	"bb0lojt3thmfs9hamkjd",

    "private_id":	"491l138rvo1yulvj31iwij1bcjncqx",

    "updated_at":	"2017-12-20 08:05:05.315037 +0000 UTC"

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


    "email":	"quas_saepe_enim@Twitterlist.name",

    "public_id":	"tg32mr4f94n55dj1w7v62c3hl9bvns",

    "two_factor_auth":	true,

    "hash":	"2qh73nz5dd6dkb6trcd2",

    "created_at":	"2017-12-20 08:05:05.316402 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.316412 +0000 UTC",

    "username":	"enim",

    "tenant_id":	"2yj1euso3e0zuxcn5j6k",

    "private_id":	"7svylouvus6dxduoxaoo9c4xeo6n1w"

},{


    "two_factor_auth":	true,

    "updated_at":	"2017-12-20 08:05:05.317897 +0000 UTC",

    "username":	"debitis_laudantium",

    "email":	"DebraHayes@Jaxworks.biz",

    "private_id":	"83oke47u3mla3n3czrtb91eu5s45qn",

    "hash":	"kzmz6ajjt33ij44ayhsf",

    "created_at":	"2017-12-20 08:05:05.317887 +0000 UTC",

    "public_id":	"5ov17u95mjfhwob96s97ql6kabau4d",

    "tenant_id":	"a44tg33znwe7civ4ow7t"

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


    "email":	"1Ward@Browsetype.info",

    "password":	"1m2w145a9jap4109icmi",

    "password_confirm":	"zt4bk0hh0bqzmw64m1rr",

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

