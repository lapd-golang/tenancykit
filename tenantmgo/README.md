Tenant MongoDB API
===================================

Tenant MongoDB API is a auto-generated CRUD implementation for the `Tenant` in package `github.com/gokit/tenancykit`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(col *mgo.Collection) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Create

```go
Create(ctx context.Context, elem tenancykit.Tenant) error
```

## Get

```go
Get(ctx context.Context, publicID string) (tenancykit.Tenant, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]tenancykit.Tenant, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem tenancykit.Tenant) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```