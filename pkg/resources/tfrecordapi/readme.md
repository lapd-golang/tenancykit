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

    "domain":	"Jaxbean.org"

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


    "created_at":	"2017-12-25 20:33:45.666845 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.666857 +0000 UTC",

    "key":	"vr9g1cql8pxehwzixw34",

    "domain":	"Photofeed.gov",

    "user_id":	"9v7fa37o3666vovjjs7y",

    "totp":	"i3t57yn9zzdu72nth838",

    "tenant_id":	"7q7opt3how9d6kbrha0ept1vdk5et5",

    "code_length":	11

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


    "code_length":	1,

    "created_at":	"2017-12-25 20:33:45.667846 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.667862 +0000 UTC",

    "key":	"n1yggumdx6a5iz24v7iy",

    "domain":	"Trudoo.mil",

    "user_id":	"1m0kzu4pq7cq9n8qlla3",

    "totp":	"2svdjava6a20wqbvkdjo",

    "tenant_id":	"0f9yvvmqazzjtbls54omhqily2vr33"

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


    "code_length":	11,

    "created_at":	"2017-12-25 20:33:45.668582 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.668592 +0000 UTC",

    "key":	"0psuy1hpmry3by3chfi6",

    "domain":	"Yodoo.net",

    "user_id":	"f4tufe5llny5601691ye",

    "totp":	"v09r0ccacne3x648le9q",

    "tenant_id":	"3r46xlidmprrn6vc3wvdk5mr9oh9am"

},{


    "code_length":	17,

    "created_at":	"2017-12-25 20:33:45.669186 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.669195 +0000 UTC",

    "key":	"olk3gc0zn5eif0g37p65",

    "domain":	"Trudeo.edu",

    "user_id":	"syqi49cg932ghb1thwzo",

    "totp":	"2o2pudrmu1xant57vsgm",

    "tenant_id":	"h3xbii1emyu4c5lvuzcnh8i5lcbd00"

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


    "totp":	"ifq0jivpb4tdtenazda1",

    "tenant_id":	"gwtzthigb2obb7gyqlr3dkqipph9cx",

    "code_length":	0,

    "created_at":	"2017-12-25 20:33:45.670002 +0000 UTC",

    "updated_at":	"2017-12-25 20:33:45.67002 +0000 UTC",

    "key":	"p9int7v9nu2uwbtu363k",

    "domain":	"Avavee.com",

    "user_id":	"eszddjxytuwblsgs8fia"

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

