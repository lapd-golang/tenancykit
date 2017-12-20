package backend

import (
	"github.com/gokit/tenancykit/tokenset/tokens"
	"github.com/gokit/tenancykit/tokenset/tokens/types"
	"github.com/influx6/faux/context"
)

// TokenBackend implements tokenapi.TokenBackend.
type TokenBackend struct {
	types.TokenDBBackend
}

// Create attempts to save provided token into db and returning token once done.
// It implements tokenapi.TokenBackend interface.
func (t TokenBackend) Create(ctx context.Context, elem tokens.Token) (tokens.Token, error) {
	return elem, t.TokenDBBackend.Create(ctx, elem)
}
