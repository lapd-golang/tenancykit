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


    "name":	"Debra King",

    "email":	"similique@Aimbo.info"

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


    "name":	"Mrs. Ms. Miss Deborah Bailey",

    "email":	"FredMartin@Jaxnation.info",

    "public_id":	"89h208hxlcyo15pnoi11x8bwliirs6",

    "created_at":	"2017-12-20 19:23:27.324332 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.324341 +0000 UTC"

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


    "name":	"Evelyn Harris",

    "email":	"quis_cumque@Shufflester.net",

    "public_id":	"fh72jvve94wzdtddk2drd4fnu28ye6",

    "created_at":	"2017-12-20 19:23:27.324673 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.324677 +0000 UTC"

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


    "name":	"Donna Olson",

    "email":	"1Butler@Trunyx.biz",

    "public_id":	"y1eej6fb32i92017hkpr1ocyin6x0q",

    "created_at":	"2017-12-20 19:23:27.325109 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.325113 +0000 UTC"

},{


    "name":	"Louise Thomas",

    "email":	"RobertTucker@Wikizz.com",

    "public_id":	"8orhgph7do58sbu5r28524q09h25ut",

    "created_at":	"2017-12-20 19:23:27.325498 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.325503 +0000 UTC"

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


    "public_id":	"6axyjdq4anwrgu8ne7s6mfjemq9avc",

    "created_at":	"2017-12-20 19:23:27.325846 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.325851 +0000 UTC",

    "name":	"Michelle Allen",

    "email":	"eos_dignissimos@Feedfish.biz"

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

