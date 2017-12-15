User SQLDB API
===================================

User SQLDB API is a auto-generated CRUD implementation for the `User` in package `github.com/gokit/tenancykit/users`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Create

```go
Create(ctx context.Context, elem users.User) error
```

## GetByField

```go
GetByField(ctx context.Context, key string, value string)  (users.User,  error)
```

## Get

```go
Get(ctx context.Context, publicID string) (users.User, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]users.User, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem users.User) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```
