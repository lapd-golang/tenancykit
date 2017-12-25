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


    "max_length":	14,

    "user":	null,

    "domain":	"Mynte.info"

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


    "domain":	"Divanoodle.mil",

    "user_id":	"10m8fvw7nnmm7gmoj7l6",

    "totp":	"e2k6cib6cvy8rw27kbiz",

    "tenant_id":	"mry2i0ulc7svk6o4ywklp0bdji52n0",

    "code_length":	16,

    "created_at":	"2017-12-25 16:32:46.216623 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.216632 +0000 UTC",

    "key":	"pwmfhwwlxeuhd8hchwfm"

}
```

## INFO /TFRecord/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.TFRecord.

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


    "tenant_id":	"r5htmdu6wasa6y2wumm0ph6azqpfzg",

    "code_length":	15,

    "created_at":	"2017-12-25 16:32:46.217184 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.21719 +0000 UTC",

    "key":	"1fm8g32rg8jgjm6fvy73",

    "domain":	"Yozio.mil",

    "user_id":	"j7l0fuiewhui5vq9mwjv",

    "totp":	"3zbgssaigzb39ucmevkc"

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


    "updated_at":	"2017-12-25 16:32:46.217647 +0000 UTC",

    "key":	"pz2mq4elh08x53u8in0x",

    "domain":	"Dabtype.name",

    "user_id":	"2029h0ofjwxjop3qa2u1",

    "totp":	"4r3msmsopvh8rrrqcjna",

    "tenant_id":	"t37075kcsza9h0bubw2dsrxxpm6kds",

    "code_length":	13,

    "created_at":	"2017-12-25 16:32:46.217642 +0000 UTC"

},{


    "totp":	"skooo22r5btmtkvlqcmi",

    "tenant_id":	"clpp9kmufcc373q2hmscxgeg6xwuul",

    "code_length":	8,

    "created_at":	"2017-12-25 16:32:46.217973 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.217977 +0000 UTC",

    "key":	"4gcjwnq2ba1rvhjo02cq",

    "domain":	"Edgetag.edu",

    "user_id":	"02ltip1gebmepogya46g"

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


    "user_id":	"5bzi0ggx5zjmhxp71mcd",

    "totp":	"pa0jjwo9wkr9szaeaeo5",

    "tenant_id":	"ta6khrphrmi3z6dr6agzn472u5y5qz",

    "code_length":	11,

    "created_at":	"2017-12-25 16:32:46.218575 +0000 UTC",

    "updated_at":	"2017-12-25 16:32:46.218581 +0000 UTC",

    "key":	"wue8ie3gfnw91gtxeod2",

    "domain":	"Kimia.edu"

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

