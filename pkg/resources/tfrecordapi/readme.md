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

    "domain":	"Omba.net",

    "max_length":	18

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


    "code_length":	8,

    "created_at":	"2017-12-26 15:18:04.61767 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.617679 +0000 UTC",

    "key":	"2kyzvmm9wim4qd796iwe",

    "domain":	"Kwideo.net",

    "user_id":	"wlc3kqv1y3pzhm3xvqx1",

    "totp":	"19t454ykgpfeivpqd0s3",

    "tenant_id":	"lkbbme2wu4xvilgqsyflutzs3ouz8t"

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


    "domain":	"Gabspot.mil",

    "user_id":	"hh3aajeonj13cf4dxhgx",

    "totp":	"h8rw7hatdtew95afhr9w",

    "tenant_id":	"51iejcf39xmka36530o0yadxmc6axg",

    "code_length":	16,

    "created_at":	"2017-12-26 15:18:04.618207 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.618212 +0000 UTC",

    "key":	"fqzhg2ezm6cd740i5z1s"

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


    "totp":	"9ua0t7cb12hri4h11j1e",

    "tenant_id":	"abqf2fx07ed1x7cv2fbc716qsgijct",

    "code_length":	9,

    "created_at":	"2017-12-26 15:18:04.618637 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.618642 +0000 UTC",

    "key":	"7s9ye017yb8bpa4enamh",

    "domain":	"Trupe.info",

    "user_id":	"bdbzysw66zww91ci0e52"

},{


    "user_id":	"1ynziet7kwkp4p4h3b20",

    "totp":	"lry4kczayi9lcoh181mg",

    "tenant_id":	"hb4rgqvvarm8s6t95phr0foxj93nko",

    "code_length":	4,

    "created_at":	"2017-12-26 15:18:04.619032 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.619037 +0000 UTC",

    "key":	"qnx8af06w4gw8bmquqhy",

    "domain":	"Blogpad.biz"

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


    "user_id":	"0bwq9w0qxu1c2pxadafl",

    "totp":	"frfoqth7xfw2ttb9tymj",

    "tenant_id":	"v8subzj7hdtf734m33hsqpti4cx6nd",

    "code_length":	7,

    "created_at":	"2017-12-26 15:18:04.619538 +0000 UTC",

    "updated_at":	"2017-12-26 15:18:04.619543 +0000 UTC",

    "key":	"ps5vnkb8s12liy1enzl6",

    "domain":	"Twitternation.biz"

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

