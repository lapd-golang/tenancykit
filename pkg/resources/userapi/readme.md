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


    "username":	"yMiller",

    "password":	"q425t2nzo4oww61wn79b",

    "password_confirm":	"bqnqfsazhnyk4mnbfxam",

    "twofactor_auth":	true,

    "tenant_id":	"eef8hnd9iu15vxazw5nz",

    "email":	"gBrown@Eare.com"

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


    "created_at":	"2017-12-24 06:38:45.344114 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.344123 +0000 UTC",

    "private_id":	"r28z2c2hyxy53ye4r2nxrip4l07trv",

    "hash":	"iz1xclc5cnr16jf5acq4",

    "two_factor_auth":	true,

    "username":	"3Adams",

    "email":	"ToddJones@Thoughtmix.gov",

    "public_id":	"f1h9wtxs0dxg4j0br01s1219btx3ol",

    "tenant_id":	"z468wv7ee3t0mdfgptg0"

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


    "updated_at":	"2017-12-24 06:38:45.344973 +0000 UTC",

    "username":	"zMitchell",

    "email":	"hBerry@Flipbug.gov",

    "tenant_id":	"ep43o409ww29s4s9er7u",

    "private_id":	"0idkum06615nlg7y1jewu44lb2ca70",

    "hash":	"sk6tb2qkfgqhrpbbe8ss",

    "two_factor_auth":	true,

    "public_id":	"8pysvn5qp0cwnqk1pzlyod2q3vp4gc",

    "created_at":	"2017-12-24 06:38:45.344965 +0000 UTC"

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


    "email":	"0Wheeler@Gabtune.net",

    "public_id":	"q4c8aiep2mw11kii1asqjefhemsvyh",

    "private_id":	"0iq4m2r27ymzir5e51dm5trakrv8p8",

    "hash":	"al48yr2486766fcqrl5b",

    "two_factor_auth":	true,

    "created_at":	"2017-12-24 06:38:45.345603 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.345609 +0000 UTC",

    "username":	"6Taylor",

    "tenant_id":	"vdlm4s1tvqs7kklwld40"

},{


    "tenant_id":	"pnqcvyip08js40qu6blh",

    "private_id":	"zzlvkbjexavx38vdqldwu6994ykcht",

    "hash":	"5mx97iqbwwna5kd7b9t6",

    "two_factor_auth":	true,

    "created_at":	"2017-12-24 06:38:45.346208 +0000 UTC",

    "public_id":	"mwrao3yed0zepu6xqa8nxq7i5e79l0",

    "email":	"ShawnFields@Kamba.name",

    "updated_at":	"2017-12-24 06:38:45.346214 +0000 UTC",

    "username":	"6Morgan"

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


    "email":	"voluptas@Meeveo.mil",

    "password":	"wtc3pxdwmtai9icivdki",

    "password_confirm":	"hquxsb3wcg1w7eazbze2"

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

