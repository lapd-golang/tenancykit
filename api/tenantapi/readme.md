Tenant HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/tenantapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/tenantapi)

Tenant HTTP API is a auto generated http api for the struct `Tenant`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.Tenant, error)
    Update(context.Context, string, tenancykit.Tenant) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.Tenant, int, error)
    Create(context.Context, tenancykit.CreateTenant) (tenancykit.Tenant, error)
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


    "name":	"Douglas Fields",

    "email":	"pSimpson@Vipe.name"

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


    "created_at":	"2017-12-20 08:05:05.645478 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.645487 +0000 UTC",

    "name":	"Tammy Lee I II III IV V MD DDS PhD DVM",

    "email":	"3Wagner@Linklinks.net",

    "public_id":	"7s2zoe5vm23h59spx7d1xc9k8z17k0"

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


    "public_id":	"q05q3cdsboa1i7slua7rz2j7v2rckm",

    "created_at":	"2017-12-20 08:05:05.645906 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.645911 +0000 UTC",

    "name":	"Paul Rodriguez",

    "email":	"6Long@Skipstorm.biz"

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


    "name":	"Timothy Butler",

    "email":	"RaymondMartinez@Tagchat.com",

    "public_id":	"ywobemk8lmcdolmpeqj0zac5xaz9y1",

    "created_at":	"2017-12-20 08:05:05.646465 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.64647 +0000 UTC"

},{


    "updated_at":	"2017-12-20 08:05:05.647124 +0000 UTC",

    "name":	"Carol Allen",

    "email":	"reiciendis@Kazio.com",

    "public_id":	"bdayx417adzzqfpjwf71h8nd27u9y6",

    "created_at":	"2017-12-20 08:05:05.647116 +0000 UTC"

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


    "name":	"Roger Jenkins",

    "email":	"LarryRobertson@Rooxo.org",

    "public_id":	"vgcbw5umzjmcd537tzvnrlej7pw8et",

    "created_at":	"2017-12-20 08:05:05.648208 +0000 UTC",

    "updated_at":	"2017-12-20 08:05:05.648218 +0000 UTC"

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

