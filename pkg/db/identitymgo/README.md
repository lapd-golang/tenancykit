Identity MongoDB API
===================================
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/db/identitymgo)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/db/identitymgo)

Identity MongoDB API is a auto-generated CRUD implementation for the `Identity` in package `github.com/gokit/tenancykit/pkg/jwt/userclaimjwt`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(col *mgo.Collection) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Count

```go
Count(ctx context.Context) (int, error)
```

## Create

```go
Create(ctx context.Context, elem userclaimjwt.Identity) error
```

## Get

```go
Get(ctx context.Context, publicID string) (userclaimjwt.Identity, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]userclaimjwt.Identity, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem userclaimjwt.Identity) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```