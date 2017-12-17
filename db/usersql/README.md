User SQLDB API
===================================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/db/usersql)](https://goreportcard.com/report/github.com/gokit/tenancykit/db/usersql)

User SQLDB API is a auto-generated CRUD implementation for the `User` in package `github.com/gokit/tenancykit`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Create

```go
Create(ctx context.Context, elem tenancykit.User) error
```

## GetByField

```go
GetByField(ctx context.Context, key string, value string)  (tenancykit.User,  error)
```

## Get

```go
Get(ctx context.Context, publicID string) (tenancykit.User, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]tenancykit.User, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem tenancykit.User) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```
