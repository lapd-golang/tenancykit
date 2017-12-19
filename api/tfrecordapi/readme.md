TFRecord HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/tfrecordapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/tfrecordapi)

TFRecord HTTP API is a auto generated http api for the struct `TFRecord`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.TFRecord, error)
    Update(context.Context, string, tenancykit.TFRecord) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.TFRecord, int, error)
    Create(context.Context, tenancykit.NewTF) (tenancykit.TFRecord, error)
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


    "domain":	"Mybuzz.name",

    "max_length":	10,

    "tenant":	nil,

    "user":	nil

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


    "key":	"7lkwet7uo0r3quztro9c",

    "code_length":	11,

    "created_at":	"2017-12-19 18:15:32.336901 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.336913 +0000 UTC",

    "domain":	"Realfire.gov",

    "user_id":	"9f68xnhepo63f2sl9hmr",

    "totp":	"v4u1yf07s7kbb68abons",

    "public_id":	"z1ussbkxh8blz5oth37rp5d6fowtfr",

    "tenant_id":	"hhq2xlohxdv6lhssiqso"

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


    "domain":	"Oyope.name",

    "code_length":	1,

    "created_at":	"2017-12-19 18:15:32.337905 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.337917 +0000 UTC",

    "key":	"hvl3uqxnehwt2cvhitec",

    "user_id":	"ttkb6wlv6fajujpyubq4",

    "totp":	"5t97jceyqterdulvsysf",

    "public_id":	"927wacpo9jmca7hskq8njcdnpgeh2q",

    "tenant_id":	"mvqva8d6ny6ick4s2akl"

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


    "key":	"xo3tybeieubbvgz0uezb",

    "user_id":	"rq2mlml53ifsamz5692i",

    "public_id":	"l8s964xy3wvvw6cxi90vdgwecfyww2",

    "tenant_id":	"nxy1x9r9vbihqnxr0u6r",

    "updated_at":	"2017-12-19 18:15:32.33881 +0000 UTC",

    "domain":	"Fliptune.name",

    "totp":	"lgzshbz47m85o3ymj9qa",

    "code_length":	11,

    "created_at":	"2017-12-19 18:15:32.338799 +0000 UTC"

},{


    "key":	"79eqnsemy6la6g6zqjiq",

    "user_id":	"7r245ui97nalqzp2kw6r",

    "totp":	"0evffzoq52dm5mouz1iy",

    "tenant_id":	"snxe5v2rrtbtw14z5mnp",

    "code_length":	17,

    "created_at":	"2017-12-19 18:15:32.339375 +0000 UTC",

    "updated_at":	"2017-12-19 18:15:32.339385 +0000 UTC",

    "domain":	"Browsezoom.name",

    "public_id":	"3rob5ti2a55e66il6xm9b0ztcxmhy5"

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


    "user_id":	"t93r4fw8ihsjimjg9l12",

    "public_id":	"8l29pqy6et3gz3b976cube108if82v",

    "code_length":	0,

    "created_at":	"2017-12-19 18:15:32.340096 +0000 UTC",

    "domain":	"Zoomdog.com",

    "totp":	"pgxyooa9qeos5j9uuzg5",

    "tenant_id":	"dr6wnfk4h44mvyu463gn",

    "updated_at":	"2017-12-19 18:15:32.340105 +0000 UTC",

    "key":	"henj49awpvjk1dgottgc"

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

