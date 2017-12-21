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


    "name":	"Thomas Lane",

    "email":	"AndreaThomas@Browseblab.mil"

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


    "created_at":	"2017-12-21 13:40:49.093831 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.093855 +0000 UTC",

    "name":	"Susan Howell I II III IV V MD DDS PhD DVM",

    "email":	"necessitatibus@Layo.name",

    "public_id":	"tzlzxs1cwyj872srgwasoni9ttyomj"

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


    "name":	"Kimberly Ferguson",

    "email":	"eum_est@Gigashots.name",

    "public_id":	"5ksugm2c200f7ctp1iwmt0gkjn0hov",

    "created_at":	"2017-12-21 13:40:49.094148 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.094153 +0000 UTC"

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


    "updated_at":	"2017-12-21 13:40:49.094458 +0000 UTC",

    "name":	"Mr. Dr. James Reynolds",

    "email":	"rerum_sed@Innojam.info",

    "public_id":	"fc07kggvqmmh78u5z54aaxdlx0vcsf",

    "created_at":	"2017-12-21 13:40:49.094453 +0000 UTC"

},{


    "name":	"Jean Banks",

    "email":	"tenetur@Devpoint.org",

    "public_id":	"6gde041k4llwiu804wtu6f6ll1h3g8",

    "created_at":	"2017-12-21 13:40:49.094689 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.094693 +0000 UTC"

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


    "email":	"ElizabethPeters@Wikizz.name",

    "public_id":	"4lhl6awkq2r55jnt18m9gdr5k0weht",

    "created_at":	"2017-12-21 13:40:49.094977 +0000 UTC",

    "updated_at":	"2017-12-21 13:40:49.094981 +0000 UTC",

    "name":	"Teresa Stewart"

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

