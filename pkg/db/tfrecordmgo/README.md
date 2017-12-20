TFRecord MongoDB API
===================================
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/db/tfrecordmgo)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/db/tfrecordmgo)

TFRecord MongoDB API is a auto-generated CRUD implementation for the `TFRecord` in package `github.com/gokit/tenancykit/pkg`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(col *mgo.Collection) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Create

```go
Create(ctx context.Context, elem pkg.TFRecord) error
```

## Get

```go
Get(ctx context.Context, publicID string) (pkg.TFRecord, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]pkg.TFRecord, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem pkg.TFRecord) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```