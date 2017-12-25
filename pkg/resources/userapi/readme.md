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


    "tenant_id":	"dkdv7f8igu8p0hfuk28h",

    "email":	"JerryLopez@Gevee.name",

    "username":	"optio",

    "password":	"5lfk1ft8kloemg2b7vjn",

    "password_confirm":	"9ozbm4qek5xtql3csnz5",

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


    "public_id":	"bei5r2lyrtvw6qengd9cywj4z6ywd8",

    "private_id":	"xfcwm0zsakczvgz4vq6axeux695lub",

    "hash":	"qam3mgfb7qc69jg5ghk1",

    "updated_at":	"2017-12-25 20:33:45.962063 +0000 UTC",

    "username":	"2Hamilton",

    "email":	"voluptates_eius_asperiores@Dynazzy.mil",

    "two_factor_auth":	true,

    "created_at":	"2017-12-25 20:33:45.962049 +0000 UTC",

    "tenant_id":	"lgy3ury17dsdrk19be6k",

    "user_roles":	null

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


    "public_id":	"63kheeteqdzmzpaxi8hxflt6dyre66",

    "private_id":	"y2ini2g2byiulohllu5eyzl7hp4yqb",

    "user_roles":	null,

    "two_factor_auth":	true,

    "created_at":	"2017-12-25 20:33:45.962943 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.962958 +0000 UTC",

    "username":	"AlanCunningham",

    "email":	"ChristinaHarper@Yata.mil",

    "tenant_id":	"arcp9bn0s09urw6l9x5k",

    "hash":	"bbndjcq5mop112qppw0m"

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


    "email":	"iusto_fuga_sed@Plambee.gov",

    "tenant_id":	"80droasrc4zqe2e8ctjd",

    "two_factor_auth":	true,

    "username":	"JoeLane",

    "public_id":	"j1e6vl0hlqapy68s2urjl0ck8upo4x",

    "private_id":	"u20w5wc6h3ou1ia9glmnynecpey4vn",

    "hash":	"acmuepv0dgcy25y9tedb",

    "user_roles":	null,

    "created_at":	"2017-12-25 20:33:45.963772 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.963784 +0000 UTC"

},{


    "tenant_id":	"7ogmd6whnoc1dzloaaap",

    "private_id":	"ul1enjt3s0spr9tw16t9fq3m1crje4",

    "hash":	"pxyfz89l8xfvrph9gcg8",

    "user_roles":	null,

    "updated_at":	"2017-12-25 20:33:45.964433 +0000 UTC",

    "username":	"esse",

    "public_id":	"r31n5myj2476pu348nz18cd4mswfpt",

    "created_at":	"2017-12-25 20:33:45.964423 +0000 UTC",

    "email":	"EugeneMeyer@Wordify.edu",

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


    "updated_at":	"2017-12-25 20:33:45.965218 +0000 UTC",

    "email":	"ea_ex@Voomm.edu",

    "public_id":	"194nkgo0r5kt8v5zgq5frlj9ly1hc5",

    "tenant_id":	"ucbbtdwcfl1xpqk31efe",

    "hash":	"mrrwclf7fjoym5qk822i",

    "user_roles":	null,

    "username":	"eum_ut",

    "private_id":	"92pqeyv52z2l20igy5a6dpoi1up7s5",

    "two_factor_auth":	true,

    "created_at":	"2017-12-25 20:33:45.965207 +0000 UTC"

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

