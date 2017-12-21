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


    "name":	"Cheryl Jackson",

    "email":	"bRichards@Tambee.info"

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


    "name":	"Susan Mitchell",

    "email":	"dolore_eligendi@Quimba.info",

    "public_id":	"xdkuiuuic2a99522dsbxrqb9zhbj2w",

    "created_at":	"2017-12-21 15:13:50.998233 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:50.998259 +0000 UTC"

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


    "name":	"Joseph Reed",

    "email":	"nHarrison@Linktype.biz",

    "public_id":	"jb269uvofrf5acicn2hcemfjcceoaq",

    "created_at":	"2017-12-21 15:13:50.998749 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:50.998755 +0000 UTC"

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


    "name":	"Amanda Wheeler",

    "email":	"velit@Rhynoodle.gov",

    "public_id":	"kcqc6eorce56gf8qaf9idw0k78nkpy",

    "created_at":	"2017-12-21 15:13:50.999058 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:50.999063 +0000 UTC"

},{


    "name":	"Mr. Dr. Eric Hanson",

    "email":	"BillyHarrison@Gigabox.mil",

    "public_id":	"8tf18a0etafyyzq9ruvawzxanjxglh",

    "created_at":	"2017-12-21 15:13:50.999318 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:50.999324 +0000 UTC"

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


    "public_id":	"klsza7hbql006c61tmqzxk2wobzyx1",

    "created_at":	"2017-12-21 15:13:50.999629 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:50.999634 +0000 UTC",

    "name":	"Jean Anderson",

    "email":	"6Richardson@Gabtype.net"

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

