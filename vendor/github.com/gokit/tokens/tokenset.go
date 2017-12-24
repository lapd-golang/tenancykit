package tokens

import (
	"context"
)

//go:generate sqlkit -generate.dest=./db generate
//go:generate mgokit -generate.dest=./db generate
//go:generate httpkit -generate.dest=./api generate
//go:generate mockkit -generate.dest=./mock generate
//go:generate mockkit -generate.dest=./mock -generate.target=./db/types generate

// TokenSet defines a interface contract which exposes a method to
// add a non-existing value for a giving targets set and querying
// if a giving target has a giving token value.
// @implement
type TokenSet interface {
	Remove(ctx context.Context, target string, token string) error
	Has(ctx context.Context, target string, token string) (bool, error)
	Add(ctx context.Context, target string, token string) (Token, error)
}
