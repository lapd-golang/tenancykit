package db

import (
	"github.com/gokit/tenancykit/pkg/db/tenantmgo"
	"github.com/gokit/tenancykit/pkg/db/tfrecordmgo"
	"github.com/gokit/tenancykit/pkg/db/twofactorsessionmgo"
	"github.com/gokit/tenancykit/pkg/db/usermgo"
	"github.com/gokit/tenancykit/pkg/db/usersessionmgo"
	"github.com/ory/hydra/pkg"
)

// IsNotFoundError returns true/false if the giving err is a not found error
// from the db subpkg.
func IsNotFoundError(err error) bool {
	switch err {
	case tenantmgo.ErrNotFound:
		return true
	case usermgo.ErrNotFound:
		return true
	case usersessionmgo.ErrNotFound:
		return true
	case tfrecordmgo.ErrNotFound:
		return true
	case twofactorsessionmgo.ErrNotFound:
		return true
	case pkg.ErrNotFound:
		return true
	}
	return false
}
