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


    "name":	"Sean Stevens",

    "email":	"zGarrett@Jabbertype.org"

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


    "updated_at":	"2017-12-27T17:22:58+01:00",

    "name":	"Gerald Collins",

    "email":	"RyanPerez@Shuffledrive.com",

    "public_id":	"cybqqtdst64r5jupw555b3gi97ixhu",

    "secret_id":	"ibt2ted4433npo24k6fb",

    "created_at":	"2017-12-27T17:22:58+01:00"

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


    "name":	"Karen Peters",

    "email":	"et@Zoombox.com",

    "public_id":	"2op53pd9eze9w47z80p7vm1lej5udm",

    "secret_id":	"cpc52lzx2z8b5fnvfgtj",

    "created_at":	"2017-12-27T17:22:58+01:00",

    "updated_at":	"2017-12-27T17:22:58+01:00"

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


    "updated_at":	"2017-12-27T17:22:58+01:00",

    "name":	"Phillip Long Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"2Mendoza@Yoveo.edu",

    "public_id":	"7cplpgu511mop7fr7osn679wmuxw3m",

    "secret_id":	"8rpsqt6hlzcu420gt5bu",

    "created_at":	"2017-12-27T17:22:58+01:00"

},{


    "name":	"Joyce Daniels",

    "email":	"dolor_provident_corrupti@Npath.net",

    "public_id":	"uz8tiscxu9zqf6shf9k8w80fxngzts",

    "secret_id":	"fy84ohbin20vdsgd2wik",

    "created_at":	"2017-12-27T17:22:58+01:00",

    "updated_at":	"2017-12-27T17:22:58+01:00"

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


    "created_at":	"2017-12-27T17:22:58+01:00",

    "updated_at":	"2017-12-27T17:22:58+01:00",

    "name":	"Terry Williamson",

    "email":	"eGreen@Shufflebeat.net",

    "public_id":	"j4mtx1lf83wsow92ec6f2qk1bt7ext",

    "secret_id":	"owjyp2g4gvvnrjrfjg28"

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

