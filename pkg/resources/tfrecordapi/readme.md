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


    "max_length":	9,

    "user":	null,

    "domain":	"Blogpad.info"

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


    "key":	"afr8vbvc799fdmhrsknu",

    "domain":	"Twitterbeat.info",

    "user_id":	"gdr8gzpn6qtgccav0d0w",

    "totp":	"h5gvkl3mnraapy4dmgv1",

    "tenant_id":	"vlirsmem1s8ba2m0xoor00eoo2tcuc",

    "code_length":	4,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00"

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


    "domain":	"Mita.info",

    "user_id":	"scf90z7jtilajmptirxa",

    "totp":	"ofed8ivnog72e6qdnztu",

    "tenant_id":	"1lw7zpphw4n9d12pgi7vfv71vn8r7n",

    "code_length":	7,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "key":	"2maz437m7t2kimxzpdmo"

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


    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "key":	"ct247p3t9u2kw61m44c2",

    "domain":	"Voonyx.info",

    "user_id":	"l4257iodp1a7mqqz14ad",

    "totp":	"i9w9r2eydot63b9scggl",

    "tenant_id":	"ujjw0w6k81f3k4ds2l38r73wknmb9k",

    "code_length":	14

},{


    "key":	"j82yagiofhw0fx23lrdq",

    "domain":	"Voonyx.net",

    "user_id":	"5xam3ro0pzvebngd50uw",

    "totp":	"h4xiqznb6x7swqo2k7xj",

    "tenant_id":	"r04pq2hfqkglyqn28i0e3qen41a40s",

    "code_length":	16,

    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00"

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


    "created_at":	"2017-12-26T18:57:10+01:00",

    "updated_at":	"2017-12-26T18:57:10+01:00",

    "key":	"7hvxqgp5xsgn7imo17bw",

    "domain":	"Agimba.org",

    "user_id":	"kwvmga07yvgnjxpoolzm",

    "totp":	"iwg53zyxv94k8iixaaym",

    "tenant_id":	"glwuruobl8ezh4k8vga9ltogcwxdpk",

    "code_length":	15

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

