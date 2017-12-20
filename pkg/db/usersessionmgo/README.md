UserSession MongoDB API
===================================
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/db/usersessionmgo)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/db/usersessionmgo)

UserSession MongoDB API is a auto-generated CRUD implementation for the `UserSession` in package `github.com/gokit/tenancykit/pkg`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(col *mgo.Collection) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Create

```go
Create(ctx context.Context, elem pkg.UserSession) error
```

## Get

```go
Get(ctx context.Context, publicID string) (pkg.UserSession, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]pkg.UserSession, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem pkg.UserSession) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```