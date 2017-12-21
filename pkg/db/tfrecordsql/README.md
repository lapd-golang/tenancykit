TFRecord SQLDB API
===================================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/pkg/db/tfrecordsql)](https://goreportcard.com/report/github.com/gokit/tenancykit/pkg/db/tfrecordsql)

TFRecord SQLDB API is a auto-generated CRUD implementation for the `TFRecord` in package `github.com/gokit/tenancykit/pkg`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Count

```go
Count(ctx context.Context) (int, error)
```

## Create

```go
Create(ctx context.Context, elem pkg.TFRecord) error
```

## GetByField

```go
GetByField(ctx context.Context, key string, value string)  (pkg.TFRecord,  error)
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
