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

    "domain":	"Youfeed.edu"

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


    "totp":	"mv4l9inzx80k18myu53z",

    "tenant_id":	"is2mip6xmxjbk40a0vltar8186e0ym",

    "code_length":	11,

    "created_at":	"2017-12-27T17:22:59+01:00",

    "updated_at":	"2017-12-27T17:22:59+01:00",

    "key":	"3wr6w09j6r0cvl0ozbf3",

    "domain":	"Meetz.gov",

    "user_id":	"k4fcarozeyg9130mnb8o"

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


    "tenant_id":	"nq695clx30ilnljqtghyo5we637vzo",

    "code_length":	1,

    "created_at":	"2017-12-27T17:22:59+01:00",

    "updated_at":	"2017-12-27T17:22:59+01:00",

    "key":	"pr8o5ms3y2kifo67b4l0",

    "domain":	"Blogspan.net",

    "user_id":	"w84szx8vuoye0d2op6bj",

    "totp":	"nby4cx3alu1rx8pi32sj"

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


    "key":	"nwrvjbotgf7qq1zbp7vi",

    "domain":	"Skipfire.gov",

    "user_id":	"0y3wb0ighv1g4stnirf3",

    "totp":	"6c7t2f4ab8qduvc7lv1m",

    "tenant_id":	"d0r3p9cap61431lkvfynjqmggs6chg",

    "code_length":	11,

    "created_at":	"2017-12-27T17:22:59+01:00",

    "updated_at":	"2017-12-27T17:22:59+01:00"

},{


    "updated_at":	"2017-12-27T17:22:59+01:00",

    "key":	"f8webos7qa842j7ygiss",

    "domain":	"Meevee.biz",

    "user_id":	"21k32jwtse3rcftc6jc4",

    "totp":	"o72j9ukqkhvxtpmc6q7f",

    "tenant_id":	"vmcrm3xrb4nu5apfopjjslcvubunzs",

    "code_length":	17,

    "created_at":	"2017-12-27T17:22:59+01:00"

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


    "domain":	"Zoomlounge.com",

    "user_id":	"oci0m68np1hfr497szg1",

    "totp":	"diw4mr83ckry3r54kd0y",

    "tenant_id":	"08zygisy78xrylvwca2at5w48ergwv",

    "code_length":	0,

    "created_at":	"2017-12-27T17:22:59+01:00",

    "updated_at":	"2017-12-27T17:22:59+01:00",

    "key":	"hmx9ccxfca85sh8k23sm"

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

