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


    "password":	"ccytf0i2x2lm8hda554v",

    "password_confirm":	"n98kkjb1upk6sjsvl29u",

    "twofactor_auth":	true,

    "tenant_id":	"vxaj1h4ahwppnrnpywk8",

    "email":	"xHamilton@Agivu.org",

    "username":	"NancyThomas"

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


    "tenant_id":	"crs6nclm7gmp4clwfrrr",

    "private_id":	"19uuc7c2iup2yupg7qqujxe4s39lfp",

    "hash":	"rdn7c4kijo754wfgdggq",

    "user_roles":	null,

    "two_factor_auth":	true,

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "email":	"AmandaRivera@Trunyx.name",

    "public_id":	"b1c0g3hbwx57aqy35e2kibz72554be",

    "created_at":	"2017-12-27T11:14:17+01:00",

    "username":	"vHarrison"

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


    "public_id":	"5fm9vnkqbzcrxgotjtjlxfcyj7rqos",

    "user_roles":	null,

    "two_factor_auth":	true,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "email":	"laborum@Kamba.name",

    "tenant_id":	"5f7yc2oc404bangb1z2i",

    "private_id":	"677x3dt6u39lune6g1yws3don5bjae",

    "hash":	"vcxxym9miyjciobmb6fv",

    "username":	"nihil_cumque_perferendis"

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


    "email":	"RebeccaReid@Geba.org",

    "public_id":	"yo8ydrpx3njfhyl0skhgxmljt8u6t6",

    "private_id":	"i7i75isfkoa7qrga2mgsrpumday6yt",

    "hash":	"adwq629a9qby4x6xmonb",

    "user_roles":	null,

    "username":	"8Jenkins",

    "tenant_id":	"k90752li8qknieg2gznx",

    "two_factor_auth":	true,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00"

},{


    "tenant_id":	"pnq9vyxhixlcruyi5rq1",

    "user_roles":	null,

    "two_factor_auth":	true,

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "username":	"JoeTurner",

    "email":	"KevinHarvey@Digitube.info",

    "public_id":	"ldyn69n95pmot982gx2cv203h9cb2w",

    "private_id":	"ljzbrpbqcpkqyh9xvzufl6aj9lz3ff",

    "hash":	"nstgu4b78uigwa2zoquu",

    "created_at":	"2017-12-27T11:14:17+01:00"

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


    "tenant_id":	"l66d4e3tjek0g4ziomir",

    "hash":	"rctex1bmcd63r8ayuuz9",

    "two_factor_auth":	true,

    "created_at":	"2017-12-27T11:14:17+01:00",

    "updated_at":	"2017-12-27T11:14:17+01:00",

    "username":	"inventore",

    "email":	"MarthaGomez@Quatz.com",

    "public_id":	"pfnguibb1z189dkpoh9f8klam9c0qo",

    "private_id":	"pcqwhpkat9b64h2cu4trcnr2wep5u4",

    "user_roles":	null

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

