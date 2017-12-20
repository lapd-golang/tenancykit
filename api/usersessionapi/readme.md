UserSession HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/api/usersessionapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/api/usersessionapi)

UserSession HTTP API is a auto generated http api for the struct `UserSession`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (tenancykit.UserSession, error)
    Update(context.Context, string, tenancykit.UserSession) error
    GetAll(context.Context, string, string, int, int) ([]tenancykit.UserSession, int, error)
    Create(context.Context, tenancykit.CreateUserSession) (tenancykit.UserSession, error)
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


    "email":	"tFerguson@Skipstorm.edu",

    "password":	"q7ezleavqkq8wl68c0t3",

    "expiration":	nil

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


    "user_id":	"npq4kpht1xs66hxybqf1",

    "public_id":	"xxp52qmmnuycur08sampwzfe76djgc",

    "token":	"f1bgcp7s3j6cm9ns9ud6",

    "expires":	"2017-12-20 12:28:53.677854 +0000 UTC"

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


    "user_id":	"vmgmxb0dfkbmqtqexts1",

    "public_id":	"o26nl06blbl3vbj9g38uxhfu9s5ya1",

    "token":	"aigydplzs62bfhdr7rpv",

    "expires":	"2017-12-20 12:28:53.678401 +0000 UTC"

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


    "user_id":	"o298fersr91rorn9kzv6",

    "public_id":	"rvntmi9kuwk1miixk3eruz61olygpw",

    "token":	"387fjfk8xab7b2e8oz4x",

    "expires":	"2017-12-20 12:28:53.678866 +0000 UTC"

},{


    "user_id":	"cauv298tv0l3nryej3h3",

    "public_id":	"z8e6tucic8f2zhxb9t91w9jlnf4xqz",

    "token":	"szezz8mvhhchu3zk8jz1",

    "expires":	"2017-12-20 12:28:53.679389 +0000 UTC"

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


    "user_id":	"jfffcnzy086t4yf8zyxy",

    "public_id":	"0to49w742bfrmbplfoc21d4500ykdm",

    "token":	"kxvghs032kz8490peyb1",

    "expires":	"2017-12-20 12:28:53.680104 +0000 UTC"

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

