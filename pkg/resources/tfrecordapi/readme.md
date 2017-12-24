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

    "user":	null,

    "domain":	"Blogspan.net"

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


    "key":	"wili77tudp5mkv59b2lk",

    "domain":	"Voonder.mil",

    "user_id":	"sng40b6iacf3exwajtp9",

    "totp":	"rovi8b7ul1zugtzfsf63",

    "tenant_id":	"t2zg489wwuh4b9gm04bydrcir2au1y",

    "code_length":	11,

    "created_at":	"2017-12-24 18:16:09.132143 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.132152 +0000 UTC"

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


    "key":	"jk0jdds7kmv8xpy06rn2",

    "domain":	"Yombu.name",

    "user_id":	"vmyp5bs9m9wq42mfwd11",

    "totp":	"nptk3su4umcub01ka6nz",

    "tenant_id":	"p7usnn2lytdng1m22h129f8uhqsj79",

    "code_length":	1,

    "created_at":	"2017-12-24 18:16:09.132792 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.132798 +0000 UTC"

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


    "user_id":	"xq1dl6obm7hl2o7by0d1",

    "totp":	"sbpxxug8czv421dtkefs",

    "tenant_id":	"ullzho6sl2s1h7npyy1yeb8vaqxzue",

    "code_length":	11,

    "created_at":	"2017-12-24 18:16:09.133273 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.133281 +0000 UTC",

    "key":	"r8lvldbi55tsc74za4ib",

    "domain":	"Jaxspan.name"

},{


    "user_id":	"o972surzdan767yp4d6g",

    "totp":	"gcde2jtcw6kv96zrz9pz",

    "tenant_id":	"uckjvoc0wv8o5maoxhep9jr5rcj2wg",

    "code_length":	17,

    "created_at":	"2017-12-24 18:16:09.133711 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.133716 +0000 UTC",

    "key":	"xgchgsrd4p69510qcjkw",

    "domain":	"Feedspan.edu"

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


    "key":	"ayhhogfz60rpq1dm52oc",

    "domain":	"Skidoo.com",

    "user_id":	"44uk3d1v8veadlv7ko4w",

    "totp":	"wmi0atsxc895ncbpxamy",

    "tenant_id":	"xus2bgx3q38tpv8kqqoz1lx5pwzlue",

    "code_length":	0,

    "created_at":	"2017-12-24 18:16:09.134437 +0000 UTC",

    "updated_at":	"2017-12-24 18:16:09.134443 +0000 UTC"

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

