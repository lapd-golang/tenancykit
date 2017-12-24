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


    "max_length":	10,

    "tenant":	null,

    "user":	null,

    "domain":	"Skimia.biz"

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


    "public_id":	"6mzpcqc67dr51l0s2ljpj68biu2jlm",

    "code_length":	11,

    "created_at":	"2017-12-24 06:38:45.185956 +0000 UTC",

    "key":	"3mvmnh3ukvgtamt7idtr",

    "user_id":	"evlak8b5tvofd92ic0ry",

    "totp":	"quubhsvuo30utxkca1xo",

    "domain":	"Wordware.com",

    "tenant_id":	"2w3hr2xqhizj98gkd6b2",

    "updated_at":	"2017-12-24 06:38:45.185965 +0000 UTC"

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


    "key":	"ft0eb59egk20i5815tuu",

    "tenant_id":	"86bg7y38nliol3hqpled",

    "created_at":	"2017-12-24 06:38:45.186644 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.18665 +0000 UTC",

    "domain":	"Livetube.org",

    "user_id":	"55p8pz9b7j85g7169kp5",

    "totp":	"lj7v6lzbqmaozhll1kt4",

    "public_id":	"scrcdxvdep943nt4hefjk4twlisboj",

    "code_length":	1

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


    "domain":	"Buzzster.info",

    "public_id":	"idimqbq6o77ao2q3pp1y5ws6gt6lhz",

    "key":	"dc1znteqfxnldfmtxed3",

    "user_id":	"azr5s08fni3k0i9rkryf",

    "totp":	"ve0j8muby38al2gtpxxm",

    "tenant_id":	"pqmcicwih60fm87b48rl",

    "code_length":	11,

    "created_at":	"2017-12-24 06:38:45.187225 +0000 UTC",

    "updated_at":	"2017-12-24 06:38:45.18723 +0000 UTC"

},{


    "updated_at":	"2017-12-24 06:38:45.187656 +0000 UTC",

    "key":	"cje5ndqm14r7rngfds60",

    "domain":	"Oba.biz",

    "public_id":	"1vq92jwm6lxisk1lyicqg9tdhtgcsk",

    "code_length":	17,

    "created_at":	"2017-12-24 06:38:45.187651 +0000 UTC",

    "user_id":	"f0cyb6czn08rld2n38v7",

    "totp":	"nx5hqdaz91uud9bq5aiv",

    "tenant_id":	"o138hzbc0n1m2eh3j05p"

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


    "user_id":	"jk1vwxrnrg03pv71h1xs",

    "public_id":	"egsjn3wol4185m4j0xvo3rmsw8shjb",

    "created_at":	"2017-12-24 06:38:45.188141 +0000 UTC",

    "key":	"hetgwjvfj41hzjudrl1b",

    "domain":	"Brightdog.gov",

    "code_length":	0,

    "updated_at":	"2017-12-24 06:38:45.188146 +0000 UTC",

    "totp":	"2qv6ka4mj8sem4qlcq4c",

    "tenant_id":	"wnm9gnshoq6afw7fccys"

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

