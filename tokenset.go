package tenancykit

import (
	"github.com/gokit/tenancykit/tokenset/tokens"
	"github.com/influx6/faux/context"
)

// TokenSet defines a interface contract which exposes a method to
// add a non-existing value for a giving targets set and querying
// if a giving target has a giving token value.
type TokenSet interface {
	Has(ctx context.Context, target string, token string) (bool, error)
	Add(ctx context.Context, target string, token string) (tokens.Token, error)
}
