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


    "name":	"Rose Boyd",

    "email":	"eos_possimus_molestiae@Skivee.com"

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


    "name":	"Mildred Spencer",

    "email":	"kMitchell@Jatri.info",

    "public_id":	"s7jhfbq4gp7rwq0rlorgizx8cc72m1",

    "secret_id":	"0nb1rh9fmg9jruyfcetp",

    "created_at":	"2017-12-25 20:33:46.064464 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:46.064475 +0000 UTC"

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


    "name":	"Brandon Collins",

    "email":	"omnis_autem_est@Kwideo.mil",

    "public_id":	"7vn2btwb3rsizhergakbjjf2qcurb5",

    "secret_id":	"k55m9xj8zteujrgf5b76",

    "created_at":	"2017-12-25 20:33:46.065156 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:46.065165 +0000 UTC"

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


    "created_at":	"2017-12-25 20:33:46.065855 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:46.065865 +0000 UTC",

    "name":	"Teresa James",

    "email":	"RogerWells@Eidel.info",

    "public_id":	"gs12reztramgy37nvdhq4lujcevjqe",

    "secret_id":	"3ubp1oknoa6dqm3ops2c"

},{


    "updated_at":	"2017-12-25 20:33:46.066433 +0000 UTC",

    "name":	"Mr. Dr. Howard Day",

    "email":	"quidem@Photobug.mil",

    "public_id":	"d6tf9jdb0inthv0iq5ebe4nrhdjoql",

    "secret_id":	"tup7a6la2ufg89kbpm49",

    "created_at":	"2017-12-25 20:33:46.066425 +0000 UTC"

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


    "created_at":	"2017-12-25 20:33:46.067152 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:46.067161 +0000 UTC",

    "name":	"Aaron Williamson Jr. Sr. I II III IV V MD DDS PhD DVM",

    "email":	"eGray@Devpulse.name",

    "public_id":	"ymfootawuhx19uwugmzvedezra0q99",

    "secret_id":	"mnudai5yejiehec0ilkw"

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

