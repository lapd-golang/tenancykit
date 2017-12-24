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


    "tenant_id":	"b0mqh0iancgnusk30rrq",

    "email":	"molestias@Miboo.biz",

    "username":	"quasi",

    "password":	"8dj3blclf6ousemu2xmo",

    "password_confirm":	"3cxgow3xchbpejzymzgc",

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


    "created_at":	"2017-12-24 17:21:03.799079 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.79909 +0000 UTC",

    "public_id":	"p1fa2x83q4xyyj56r6replxut4yrpm",

    "private_id":	"sxi2kyav0i9ojuzyioezlvk95kn89u",

    "two_factor_auth":	true,

    "hash":	"2hly3echyujz2rud9akk",

    "username":	"iure_excepturi_molestiae",

    "email":	"voluptates_sed_praesentium@Zoomzone.biz",

    "tenant_id":	"20fjjnxol32y8nftjpx4"

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


    "email":	"JacquelineMarshall@Quire.name",

    "hash":	"7n24jivo4qm7tvwhappe",

    "two_factor_auth":	true,

    "username":	"MarilynHarper",

    "public_id":	"c5b7n9prugt2r4rzhur1zi0mdxi7d4",

    "tenant_id":	"e849ek9bq0wpziwioa5d",

    "private_id":	"6ptqjpqtxqaqct2h1xnqpknoqokvqy",

    "created_at":	"2017-12-24 17:21:03.800403 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.800414 +0000 UTC"

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


    "email":	"tenetur@Pixoboo.gov",

    "private_id":	"sfvk4ic2qykelumwb1tz7jxrl3fjdn",

    "two_factor_auth":	true,

    "created_at":	"2017-12-24 17:21:03.801599 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.801607 +0000 UTC",

    "username":	"ad_rerum",

    "public_id":	"g98i9xytnypvj9jvb16gtvivk2eyz1",

    "tenant_id":	"vxdyipgwbdjm44cvpk6d",

    "hash":	"ins0s4kw0fejapk0zkiv"

},{


    "username":	"placeat_incidunt",

    "tenant_id":	"hlmzjrmo8rtmfp893na0",

    "private_id":	"f3ugbu96jjruork9gbtijuaz5cqlrg",

    "created_at":	"2017-12-24 17:21:03.802881 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.802896 +0000 UTC",

    "email":	"oRoss@Omba.mil",

    "public_id":	"yn68nr9b8r0yfutcmx1nco9mblj7qx",

    "hash":	"o2yak0pm88ay5qut5xgu",

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


    "email":	"JackMills@Photojam.gov",

    "password":	"y7j53iz8h0b9e8p14dq5",

    "password_confirm":	"sak7elhl0kwn7647t34v"

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

