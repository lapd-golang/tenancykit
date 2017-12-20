TFRecord HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/tfrecordapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/tfrecordapi)

TFRecord HTTP API is a auto generated http api for the struct `TFRecord`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.TFRecord, error)
    Update(context.Context, string, pkg.TFRecord) error
    GetAll(context.Context, string, string, int, int) ([]pkg.TFRecord, int, error)
    Create(context.Context, pkg.NewTF) (pkg.TFRecord, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /TFRecord/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the NewTF type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "user":	null,

    "domain":	"JumpXS.edu",

    "max_length":	8,

    "tenant":	null

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


    "user_id":	"v57rn85dmgli4muyo5fo",

    "totp":	"i18oeovykz9khmkdcql5",

    "public_id":	"8qekrdkbryb8pp26uvlw52b759ix3s",

    "tenant_id":	"7jauf6nks6raoakln4d0",

    "key":	"8o2vfx2pxce16ddf43t4",

    "domain":	"Skibox.edu",

    "code_length":	11,

    "created_at":	"2017-12-20 19:23:27.398534 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.398542 +0000 UTC"

}
```

## GET /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the TFRecord type from the HTTP API returning received result as a JSON
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


    "updated_at":	"2017-12-20 19:23:27.399075 +0000 UTC",

    "key":	"i3wn5a0r11ejpxt1xlfy",

    "user_id":	"btqhh6663sx09c3xa9r8",

    "tenant_id":	"y3z7qws1trraeq5i3n6u",

    "created_at":	"2017-12-20 19:23:27.39907 +0000 UTC",

    "domain":	"Wordtune.mil",

    "totp":	"mi9b60c5wklve2w5soet",

    "public_id":	"atsvexvjbvf0ky6mve2bddmgttktva",

    "code_length":	10

}
```

## GET /TFRecord/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the TFRecord type from the HTTP API.

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


    "created_at":	"2017-12-20 19:23:27.399591 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.399595 +0000 UTC",

    "user_id":	"1n1w91d7iak4ljuz4fmt",

    "tenant_id":	"o23ecnpi7wwzry5omyjh",

    "totp":	"udkv42cjrlektcdnsaxl",

    "public_id":	"s54n2wxbfkgqdhh1x0bgtwizj5syrs",

    "code_length":	11,

    "key":	"ltvy0219yvmb1y5nm0yj",

    "domain":	"Cogibox.biz"

},{


    "key":	"il76v72gntq2mg2v9lsg",

    "domain":	"Viva.name",

    "totp":	"sogtxesh7hk7183x66rf",

    "tenant_id":	"ygipsvfm9p5thdtf326u",

    "code_length":	13,

    "user_id":	"1bxv5dd9op42w6kwa238",

    "public_id":	"6yti5obx2klfxpcx4ajtih7m1ynnbc",

    "created_at":	"2017-12-20 19:23:27.400592 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.4006 +0000 UTC"

}]
```

## PUT /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the TFRecord type from the HTTP API returning received result as a JSON
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


    "user_id":	"upzz245egg3q0wqc8lqs",

    "public_id":	"b3a2bstha8kzxpv34l8jsd2zx0fehn",

    "tenant_id":	"ni0y68jty8n5bxzmrdcl",

    "code_length":	16,

    "created_at":	"2017-12-20 19:23:27.401818 +0000 UTC",

    "updated_at":	"2017-12-20 19:23:27.401826 +0000 UTC",

    "domain":	"Skajo.org",

    "totp":	"s2pnmzzzrrr5smmeanwj",

    "key":	"o4ucywfsotvc9ntv1umd"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /TFRecord/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the TFRecord type from the HTTP API returning received result as a JSON
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

