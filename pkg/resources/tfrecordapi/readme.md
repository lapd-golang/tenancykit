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


    "max_length":	18,

    "tenant":	null,

    "user":	null,

    "domain":	"Trilith.name"

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


    "key":	"mef9hf40eabdaled8no4",

    "user_id":	"kze3zg0c6q0q73wcoq2a",

    "public_id":	"7tx89ltkfiltg2m7ez7rdcnzyjmlah",

    "tenant_id":	"aru88h9ewug9mhzsa8ej",

    "code_length":	8,

    "created_at":	"2017-12-21 15:13:51.232143 +0000 UTC",

    "domain":	"Rhyzio.com",

    "totp":	"2c0x5r20yurlyrtbzi1k",

    "updated_at":	"2017-12-21 15:13:51.232155 +0000 UTC"

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


    "domain":	"Blogtag.net",

    "created_at":	"2017-12-21 15:13:51.233104 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:51.233122 +0000 UTC",

    "key":	"8o4iunfc69mt916z2szl",

    "user_id":	"vdct9kxfmo30vofdqhd3",

    "totp":	"yba2p26aaitus1j4uewn",

    "public_id":	"5p9slum3ml56vhq7b8fbgdrlz8vfrp",

    "tenant_id":	"81mc3x1lbkciqog7p0yv",

    "code_length":	16

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


    "key":	"dpcyq4gep6wq7m405x2h",

    "public_id":	"69k3e4g46cdtbw6xk3ws22hp7q6fyq",

    "code_length":	9,

    "updated_at":	"2017-12-21 15:13:51.233913 +0000 UTC",

    "domain":	"Dynabox.mil",

    "user_id":	"4cdgxn78gbtebyk5alwa",

    "totp":	"bsuf8e1lmogr87ozrywz",

    "tenant_id":	"iy88c2prftbeh34bbnmg",

    "created_at":	"2017-12-21 15:13:51.233906 +0000 UTC"

},{


    "key":	"t3yv4q35rt9gwgw0kgvf",

    "domain":	"Lajo.info",

    "totp":	"fci9h8tztu3axw2xhc2o",

    "public_id":	"26l3owf91l9gr8y58ywiohagj31pdg",

    "code_length":	4,

    "user_id":	"7dpf2vwf6w1lokswy9cc",

    "tenant_id":	"sdxpj8mypyuku05xu8le",

    "created_at":	"2017-12-21 15:13:51.234349 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:51.234357 +0000 UTC"

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


    "domain":	"JumpXS.info",

    "user_id":	"g851vqsvlv9mg3wiaq0s",

    "totp":	"05il73r9vt9xowlnd7pc",

    "created_at":	"2017-12-21 15:13:51.235054 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:51.235059 +0000 UTC",

    "key":	"6fx8fpmwp9cnr00ffab1",

    "public_id":	"8trrrnqvblmkstq4zrb0lf2w4kcvdf",

    "tenant_id":	"mwqrwaci649wjrnfrbxw",

    "code_length":	7

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

