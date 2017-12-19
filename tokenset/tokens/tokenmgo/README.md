Token MongoDB API
===================================
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit/tokenset/tokens/tokenmgo)](https://goreportcard.com/report/github.com/gokit/tenancykit/tokenset/tokens/tokenmgo)

Token MongoDB API is a auto-generated CRUD implementation for the `Token` in package `github.com/gokit/tenancykit/tokenset/tokens`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(col *mgo.Collection) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Create

```go
Create(ctx context.Context, elem tokens.Token) error
```

## Get

```go
Get(ctx context.Context, publicID string) (tokens.Token, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]tokens.Token, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem tokens.Token) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```