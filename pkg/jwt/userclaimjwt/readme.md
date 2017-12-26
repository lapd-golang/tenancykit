UserClaim HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/jwt/userclaimjwt)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/jwt/userclaimjwt)

UserClaim HTTP API is a auto generated http api for the struct `UserClaim`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type IdentityBackend interface {
	Count(context.Context) (int, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (Identity, error)
	Update(context.Context, string, Identity) error
	Revoke(context.Context, string) error
	Refresh(context.Context, string) (JWTAuth, error)
	GetAll(context.Context, string, string, int, int) ([]Identity, int, error)
	Grant(context.Context, pkg.CreateUserSession) (JWTAuth, error)
	Authenticate(context.Context, string) (pkg.UserClaim, error)
}
```

It exposes the following methods for each endpoint:

## Grant { POST /UserClaim/ }
### Method: `func (api *JWTAPI) Grant(ctx *httputil.Context) error`

Request login authentication using pkgCreateUserSession type which is delivered as
JSON to the API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "email":	"DeniseCruz@Voonix.mil",

    "password":	"khkse994c9tp1qgpo8ce",

    "expiration":	null

}
```

- Expected Status Code

```
Success: 200
```

- Expected Response Body

```http
    Content-Type: application/json
```

```json
{
	"expires": 323231,
	"token_type": "Bearer",
	"refresh_expires": 323231,
	"refresh_token":"7fd15938c823cf58e78019bea2af142f9449696a",
	"access_token":"7fd15938c823cf58e78019bea2af142f9449696awe434343d33434343343wcsdvdfvdffvefefvdddfvdvd",
}
```

## POST /UserClaim/revoke
### Method: `func (api *JWTAPI) Revoke(ctx *httputil.Context) error`

Revokes a given JWT refresh token.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{"refresh_token":"7fd15938c823cf58e78019bea2af142f9449696a"}
```

- Expected Status Code

```
Failure:
- 500
- 404

Success: 200
```


## POST /UserClaim/refresh
### Method: `func (api *JWTAPI) Refresh(ctx *httputil.Context) error`

Refreshes a given JWT access token by sending non-expired refreshed token as json string.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{"refresh_token":"7fd15938c823cf58e78019bea2af142f9449696a"}
```

- Expected Status Code

```
Success: 200
Failure:
	- 500
	- 404
```

- Expected Response Body

```http
Content-Type: application/json
```

```json
{
	"expires": 323231,
	"token_type": "Bearer",
	"refresh_expires": 323231,
	"refresh_token":"7fd15938c823cf58e78019bea2af142f9449696a",
	"access_token":"7fd15938c823cf58e78019bea2af142f9449696awe434343d33434343343wcsdvdfvdffvefefvdddfvdvd",
}
```

## GET /UserClaim/authenticate
### Method: `func (api *JWTAPI) Authenticate(ctx *httputil.Context) error`

Authenticate provides a method for authenticating users identity by providing a Authorization Header.

- Expected Request Header

```http
    Authorization: Bearer 7fd15938c823cf58e78019bea2af142f9449696a
```

- Expected Status Code

```
Success: 200
Failure:
	- 500
	- 404
```

## POST /UserClaim/authenticateJSON
### Method: `func (api *JWTAPI) AuthenticateJSON(ctx *httputil.Context) error`

AuthenticateJSON provides a method for authenticating users identity by providing a json payload,
containing access token and type.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{"access_token":"7fd15938c823cf58e78019bea2af142f9449696a", "type":"Bearer"}
```

- Expected Status Code

```
Success: 200
Failure:
	- 500
	- 404
```


## INFO /UserClaim/
### Method: `func (api *JWTAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.UserClaim.

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

## GET /UserClaim/:public_id
### Method: `func (api *JWTAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the UserClaim type from the HTTP API returning received result as a JSON
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

```go
type Identity struct{
	PublicID        string        `json:"public_id"`
	RefreshToken    string        `json:"refresh_token"`
	Expires         time.Time     `json:"expires"`
	TargetID        string        `json:"target_id"`
	LastLogin       time.Time     `json:"last_login"`
	IssuedAt        time.Time     `json:"last_login"`
	RefreshInterval time.Duration `json:"refresh_interval"`
	Data pkg.UserClaim `json:"data"`
}
```

## GET /UserClaim/
### Method: `func (api *JWTAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the UserClaim type from the HTTP API.

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
[Identity{}, Identity{}]
```

## PUT /UserClaim/:public_id
### Method: `func (api *JWTAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the UserClaim type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record with the provided JSON request body.

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```http
    Content-Type: application/json
```

```go
type Identity struct{
	PublicID        string        `json:"public_id"`
	RefreshToken    string        `json:"refresh_token"`
	Expires         time.Time     `json:"expires"`
	TargetID        string        `json:"target_id"`
	LastLogin       time.Time     `json:"last_login"`
	IssuedAt        time.Time     `json:"last_login"`
	RefreshInterval time.Duration `json:"refresh_interval"`
	Data pkg.UserClaim `json:"data"`
}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /UserClaim/:public_id
### Method: `func (api *JWTAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the UserClaim type from the HTTP API returning received result as a JSON
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

