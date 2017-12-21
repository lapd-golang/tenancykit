package tenancykit

import (
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/resources/tokenapi"
	"github.com/influx6/faux/metrics"
)

// TwoFactorAPI exposes the full wrapper over the http api with underline db
// for the database linked in.
type TokenAPI struct {
	tokenapi.TokenHTTP
}

// NewTokenAPI returns a new instance of TwoFactorAPI.
func NewTokenAPI(m metrics.Metrics, db types.TokenDBBackend) TokenAPI {
	return TokenAPI{
		TokenHTTP: tokenapi.New(m, backends.TokenBackend{
			TokenDBBackend: db,
		}),
	}
}
