UserSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/usersessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/usersessionapi)

UserSession HTTP API is a auto generated http api for the struct `UserSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.UserSession, error)
    Update(context.Context, string, pkg.UserSession) error
    GetAll(context.Context, string, string, int, int) ([]pkg.UserSession, int, error)
    Create(context.Context, pkg.CreateUserSession) (pkg.UserSession, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /UserSession/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the CreateUserSession type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "email":	"DorothyRamirez@Meemm.mil",

    "password":	"y2trt6wtmqf9pd97qfil",

    "expiration":	null

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


    "user_id":	"1j364r0llgwu5lqeja30",

    "public_id":	"4zzapsu12gked2h5rlrvr23qp88yep",

    "token":	"qfyjzjqw76a6s0lbt8gd",

    "expires":	"2017-12-21 15:13:51.098507 +0000 UTC"

}
```

## INFO /UserSession/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.UserSession.

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

## GET /UserSession/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the UserSession type from the HTTP API returning received result as a JSON
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


    "expires":	"2017-12-21 15:13:51.099083 +0000 UTC",

    "user_id":	"zesbws1kgoru0l7pito4",

    "public_id":	"8w5gml6a37qqhvmy2r37jotgfphd2s",

    "token":	"48sa6kbv1vh86wn4gl4a"

}
```

## GET /UserSession/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the UserSession type from the HTTP API.

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


    "user_id":	"8pgbosrv8buowcu46t6k",

    "public_id":	"u8ub75me3rtljgk25ul4kigjvohx16",

    "token":	"ptnjrrllofppkqqihfm1",

    "expires":	"2017-12-21 15:13:51.099883 +0000 UTC"

},{


    "user_id":	"kgst642luuaiu9ctvzfw",

    "public_id":	"8nkyolu130yo8grfz5x0ecjg8apg8c",

    "token":	"5cgbsl58j74mb95y7avn",

    "expires":	"2017-12-21 15:13:51.100284 +0000 UTC"

}]
```

## PUT /UserSession/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the UserSession type from the HTTP API returning received result as a JSON
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


    "user_id":	"d9ge9nyxq39voaodqlbx",

    "public_id":	"ah4ce2hi2jrx4cbeun4foyce6fpymm",

    "token":	"tgey3ngpx6kn3u8cidid",

    "expires":	"2017-12-21 15:13:51.10064 +0000 UTC"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /UserSession/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the UserSession type from the HTTP API returning received result as a JSON
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

