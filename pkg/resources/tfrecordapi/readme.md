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


    "tenant":	null,

    "user":	null,

    "domain":	"Digitube.org",

    "max_length":	10

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


    "key":	"un11to4u5x0qkjniltiv",

    "user_id":	"amevuot9mm8bn2dtmdjt",

    "totp":	"p623bgn44sgjwcwsoka7",

    "tenant_id":	"x1c7we0ov1cn5l41fdrj",

    "updated_at":	"2017-12-24 17:21:03.642525 +0000 UTC",

    "domain":	"Shuffletag.edu",

    "public_id":	"kylrib30otuzm2cidsa1gqvakvxf17",

    "code_length":	11,

    "created_at":	"2017-12-24 17:21:03.642516 +0000 UTC"

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


    "key":	"ojecf4em7hgtpe749mpl",

    "domain":	"Gabtune.biz",

    "user_id":	"xs6ne3s8frkir9pbado6",

    "public_id":	"rj7r8lvmenxjfr0ujj30pr6x0urbxk",

    "tenant_id":	"yi0uduln001ssvqld9xv",

    "created_at":	"2017-12-24 17:21:03.643235 +0000 UTC",

    "totp":	"obbgt1oinncay40rbr4p",

    "code_length":	1,

    "updated_at":	"2017-12-24 17:21:03.643241 +0000 UTC"

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


    "user_id":	"ev13u0ytpk0cshcd16vd",

    "public_id":	"oxupb4wpyrsv48iqu1sxa4mbcglq08",

    "tenant_id":	"q61clwgc1zmrtmohj35o",

    "code_length":	11,

    "created_at":	"2017-12-24 17:21:03.643809 +0000 UTC",

    "key":	"ka3v4zqefd3ycdpi2ihb",

    "domain":	"Skaboo.edu",

    "totp":	"4nhz53aa33x8nr02aay8",

    "updated_at":	"2017-12-24 17:21:03.643815 +0000 UTC"

},{


    "key":	"tut3sbp58yoncujzc1m7",

    "user_id":	"paxatuab1bjaj0shkinv",

    "totp":	"tiq79m4x7tndce55j749",

    "public_id":	"vtoozls9ntjy1aien7xzdbh4mvhzew",

    "domain":	"Abata.org",

    "tenant_id":	"050wgdr86v6ux50yte90",

    "code_length":	17,

    "created_at":	"2017-12-24 17:21:03.64424 +0000 UTC",

    "updated_at":	"2017-12-24 17:21:03.644244 +0000 UTC"

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


    "created_at":	"2017-12-24 17:21:03.644777 +0000 UTC",

    "domain":	"Mynte.name",

    "public_id":	"vym3rnirlire4rqb68sa6kltp30wqh",

    "tenant_id":	"rx5mvmn0mmlo7i52555o",

    "code_length":	0,

    "key":	"lo7dtzp9b9dska4l8yev",

    "user_id":	"i92true0minywtsq9j0b",

    "totp":	"3le3ms5dqgytwewn8iyt",

    "updated_at":	"2017-12-24 17:21:03.644782 +0000 UTC"

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

