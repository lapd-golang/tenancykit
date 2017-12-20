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


    "user":	nil,

    "domain":	"Zoonoodle.net",

    "max_length":	18,

    "tenant":	nil

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


    "public_id":	"wj8pn4hja1rdo3d6qp9igtuajxsldg",

    "created_at":	"2017-12-20 12:28:53.725712 +0000 UTC",

    "user_id":	"lgjkzzagz9j6de7o4sau",

    "totp":	"a3kti781zt7jc84jiavg",

    "tenant_id":	"cpus6hw54w3y75t4hrrs",

    "code_length":	8,

    "updated_at":	"2017-12-20 12:28:53.725722 +0000 UTC",

    "key":	"zdywdjl0368lkdyx7924",

    "domain":	"Mita.gov"

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


    "public_id":	"wokkhn2v2z5spkalmjdst3yhitofpu",

    "tenant_id":	"amkdeta5lt1ld20pw28i",

    "code_length":	16,

    "created_at":	"2017-12-20 12:28:53.727327 +0000 UTC",

    "key":	"qntn0lzp20w8pfhxjtvh",

    "totp":	"3eqdynt8itmnhjfjeftj",

    "updated_at":	"2017-12-20 12:28:53.727337 +0000 UTC",

    "domain":	"Fiveclub.org",

    "user_id":	"23679r6rke2ilxhj8zqn"

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


    "updated_at":	"2017-12-20 12:28:53.72789 +0000 UTC",

    "user_id":	"flcq023upv8hydgnrbwi",

    "totp":	"mr6zp74kk2j9euhgym5p",

    "public_id":	"giud7l4blolpoccdp5bpye0pj1r3j6",

    "tenant_id":	"6ctvq8l2p8czl80odgqd",

    "code_length":	9,

    "created_at":	"2017-12-20 12:28:53.727882 +0000 UTC",

    "key":	"qvr1ixpzk6hlizt3m1yl",

    "domain":	"Photolist.edu"

},{


    "key":	"cpo4o53ssl5lhpwmqlzr",

    "totp":	"tx0tx9pj5za3jx6x2rhp",

    "public_id":	"d2hhcsgufgljmkw8oifwie5htsm1ev",

    "created_at":	"2017-12-20 12:28:53.728336 +0000 UTC",

    "updated_at":	"2017-12-20 12:28:53.728341 +0000 UTC",

    "domain":	"Livetube.name",

    "user_id":	"r9wzvb0zwfsmzs95437e",

    "tenant_id":	"mfgowd6jjyuwukquhnya",

    "code_length":	4

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


    "domain":	"Eayo.mil",

    "user_id":	"gxdpm13n39i4mxybsmic",

    "totp":	"n2ye0ergju3yw56n5rdb",

    "public_id":	"5o72rwq0m4zoc8amh9odrc3mu1x134",

    "code_length":	7,

    "created_at":	"2017-12-20 12:28:53.728804 +0000 UTC",

    "key":	"a9qsl6b4c70ajoqh6wtt",

    "tenant_id":	"4v4g6gkf75n4zk55m0zv",

    "updated_at":	"2017-12-20 12:28:53.728809 +0000 UTC"

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

