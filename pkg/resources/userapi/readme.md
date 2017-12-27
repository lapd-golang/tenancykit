User HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/resources/userapi)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/resources/userapi)

User HTTP API is a auto generated http api for the struct `User`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (pkg.User, error)
    Update(context.Context, string, pkg.User) error
    GetAll(context.Context, string, string, int, int) ([]pkg.User, int, error)
    Create(context.Context, pkg.CreateUser) (pkg.User, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /User/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the CreateUser type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "username":	"ToddEllis",

    "password":	"wg2p6pbl2r1vsousqetu",

    "password_confirm":	"yc6yrirkmy2051jzyw06",

    "twofactor_auth":	true,

    "tenant_id":	"gyrh52ocb87vojssdao0",

    "email":	"voluptas@Avamba.mil"

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


    "updated_at":	"2017-12-27T17:22:59+01:00",

    "username":	"cMarshall",

    "email":	"eBarnes@Gabtune.com",

    "private_id":	"ypo37w23xhyhx4imz6u6ezqd7t1lbz",

    "user_roles":	null,

    "two_factor_auth":	true,

    "public_id":	"ebvellqlv3sioqhphdtzurhnmo3ol6",

    "tenant_id":	"6kmqgl0unkcwdwjaddnt",

    "hash":	"li87fk4j2rtgvsz7t8ci",

    "created_at":	"2017-12-27T17:22:59+01:00"

}
```

## INFO /User/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type pkg.User.

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

## GET /User/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the User type from the HTTP API returning received result as a JSON
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


    "email":	"JoanHolmes@Riffwire.gov",

    "public_id":	"zp4n7jgrjmcvkod0pbhihmtx993sbl",

    "hash":	"u1pj4lcs31x637ijw6sg",

    "two_factor_auth":	true,

    "created_at":	"2017-12-27T17:22:59+01:00",

    "updated_at":	"2017-12-27T17:22:59+01:00",

    "username":	"JasonKelley",

    "tenant_id":	"izwrugt3p5d2uh412uej",

    "private_id":	"ihy93ol364h84bfpwlkvqb8c95w1hu",

    "user_roles":	null

}
```

## GET /User/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the User type from the HTTP API.

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


    "user_roles":	null,

    "two_factor_auth":	true,

    "username":	"jCarpenter",

    "public_id":	"6y9hho0eeyzlt0cij81cmyn2m3pddz",

    "tenant_id":	"psza1bgw780lf35554ua",

    "hash":	"d9jvfifuloil07zpbik4",

    "email":	"SarahWashington@Mycat.net",

    "private_id":	"anoktp4epay6xrep4qz6jn18a949j0",

    "created_at":	"2017-12-27T17:22:59+01:00",

    "updated_at":	"2017-12-27T17:22:59+01:00"

},{


    "created_at":	"2017-12-27T17:22:59+01:00",

    "public_id":	"zxof8cdgpx7mfq6bes0wq77yngr6cd",

    "private_id":	"xws8njeonee9le30szqjykz8q5fnnz",

    "hash":	"sf1xuwwjdv6x39f0v50d",

    "user_roles":	null,

    "two_factor_auth":	true,

    "updated_at":	"2017-12-27T17:22:59+01:00",

    "username":	"quos_dolor",

    "email":	"gNelson@Jabbertype.info",

    "tenant_id":	"rabdtz7bno6euanuzh86"

}]
```

## PUT /User/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the User type from the HTTP API returning received result as a JSON
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


    "username":	"enim_aut_sunt",

    "tenant_id":	"1p9x171kzq46nqw6n9cr",

    "hash":	"l607fe4i0unw136kut8u",

    "user_roles":	null,

    "two_factor_auth":	true,

    "updated_at":	"2017-12-27T17:22:59+01:00",

    "email":	"voluptas_sapiente@Dabvine.edu",

    "public_id":	"yf5vfwcxzois8dmx4p2n4nl9l08zj3",

    "private_id":	"tcvnuq7ckp53l69obibejbgdidhle8",

    "created_at":	"2017-12-27T17:22:59+01:00"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /User/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the User type from the HTTP API returning received result as a JSON
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

