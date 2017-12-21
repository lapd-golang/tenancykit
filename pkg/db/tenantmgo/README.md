Tenant MongoDB API
===================================
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/db/tenantmgo)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/db/tenantmgo)

Tenant MongoDB API is a auto-generated CRUD implementation for the `Tenant` in package `github.com/gokit/tenancykit/pkg`.

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
Create(ctx context.Context, elem pkg.Tenant) error
```

## Get

```go
Get(ctx context.Context, publicID string) (pkg.Tenant, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]pkg.Tenant, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem pkg.Tenant) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```