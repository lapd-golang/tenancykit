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


    "name":	"Russell Phillips",

    "email":	"ElizabethRichards@Quaxo.gov"

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


    "public_id":	"cddrbz25hjdwor0qm83sd5ibvbiyap",

    "secret_id":	"eyajsvgulfk2l8jwnonx",

    "created_at":	"2017-12-24 17:21:03.572951 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.572974 +0000 UTC",

    "name":	"Roy Jones",

    "email":	"BrendaPerez@Voolith.gov"

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


    "name":	"Shirley Edwards",

    "email":	"uKennedy@Wordify.com",

    "public_id":	"09jioi1ppfihokritc7wllubwe9n51",

    "secret_id":	"e7ql186pf9mgpod194uy",

    "created_at":	"2017-12-24 17:21:03.573305 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.573309 +0000 UTC"

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


    "updated_at":	"2017-12-24 17:21:03.573626 +0000 UTC",

    "name":	"Terry Matthews",

    "email":	"cAdams@Zooxo.com",

    "public_id":	"uds5p2yjwc4eibaaz2xc4x54mm3kcz",

    "secret_id":	"pof2qfjhbarober2vv1z",

    "created_at":	"2017-12-24 17:21:03.57362 +0000 UTC"

},{


    "created_at":	"2017-12-24 17:21:03.573992 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.573996 +0000 UTC",

    "name":	"Patricia Taylor",

    "email":	"qui@Layo.org",

    "public_id":	"p4np7cjqhbip9uac3du20hqedpmurj",

    "secret_id":	"76tiknlxrmw3am6ypfr2"

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


    "name":	"Aaron Mitchell",

    "email":	"SteveBurton@Mudo.net",

    "public_id":	"15069bq0hgi85lh05pnbwj4avt5ly6",

    "secret_id":	"yk40et3bu2rj5575s6s3",

    "created_at":	"2017-12-24 17:21:03.574324 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.574329 +0000 UTC"

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

