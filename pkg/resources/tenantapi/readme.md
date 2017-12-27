Tenant HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/tenantapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/tenantapi)

Tenant HTTP API is a auto generated http api for the struct `Tenant`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.Tenant, error)
    Update(context.Context, string, pkg.Tenant) error
    GetAll(context.Context, string, string, int, int) ([]pkg.Tenant, int, error)
    Create(context.Context, pkg.CreateTenant) (pkg.Tenant, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /Tenant/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the CreateTenant type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "name":	"Mr. Dr. Alan Burns",

    "email":	"laudantium@Gabtune.com"

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


    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "name":	"Aaron Bradley",

    "email":	"nisi@Trudoo.biz",

    "public_id":	"pkeh6zbh26u54fpq5d0zqxwtv94mkq",

    "secret_id":	"zk689qui0izg2wzfpfuj"

}
```

## INFO /Tenant/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.Tenant.

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

## GET /Tenant/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the Tenant type from the HTTP API returning received result as a JSON
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


    "name":	"Joe Hayes",

    "email":	"EmilyMorales@Kamba.edu",

    "public_id":	"ewodolm65nxrrss6uvrzapiyjq9khd",

    "secret_id":	"wy3g7brau4uh9bdpyhr6",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

}
```

## GET /Tenant/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the Tenant type from the HTTP API.

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


    "name":	"Carlos Phillips",

    "email":	"DeborahJohnston@Youspan.info",

    "public_id":	"vq98azo1diz2rb9mw2adxh0ofcnweu",

    "secret_id":	"u2u1vsp3nq3srpajnnso",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

},{


    "updated_at":	"2017-12-27T11:14:17+01:00",

    "name":	"Marie Moore",

    "email":	"qGreen@Tagtune.org",

    "public_id":	"0lim1w0jhylnfjavuz4f4sx2jldnxd",

    "secret_id":	"2svjpb13vhdniew4e88k",

    "created_at":	"2017-12-27T11:14:17+01:00"

}]
```

## PUT /Tenant/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the Tenant type from the HTTP API returning received result as a JSON
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


    "email":	"TimothyMoore@Skivee.gov",

    "public_id":	"65wgwz0ioh2x84es4q50nc35ujno8d",

    "secret_id":	"5oses6rm0994ksr2h1vc",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "name":	"Mrs. Ms. Miss Nicole Griffin"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /Tenant/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the Tenant type from the HTTP API returning received result as a JSON
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

