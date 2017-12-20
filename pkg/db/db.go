package db

import (
	"github.com/gokit/tenancykit/pkg/db/mocks"
	"github.com/gokit/tenancykit/pkg/db/tenantmgo"
	"github.com/gokit/tenancykit/pkg/db/tenantsql"
	"github.com/gokit/tenancykit/pkg/db/tfrecordmgo"
	"github.com/gokit/tenancykit/pkg/db/tfrecordsql"
	"github.com/gokit/tenancykit/pkg/db/twofactorsessionmgo"
	"github.com/gokit/tenancykit/pkg/db/twofactorsessionsql"
	"github.com/gokit/tenancykit/pkg/db/usermgo"
	"github.com/gokit/tenancykit/pkg/db/usersessionmgo"
	"github.com/gokit/tenancykit/pkg/db/usersessionsql"
	"github.com/gokit/tenancykit/pkg/db/usersql"
)

// IsNotFoundError returns true/false if the giving err is a not found error
// from the db subpkg.
func IsNotFoundError(err error) bool {
	switch err {
	case tenantmgo.ErrNotFound:
		return true
	case tenantsql.ErrNotFound:
		return true
	case usermgo.ErrNotFound:
		return true
	case usersql.ErrNotFound:
		return true
	case usersessionmgo.ErrNotFound:
		return true
	case usersessionsql.ErrNotFound:
		return true
	case tfrecordmgo.ErrNotFound:
		return true
	case tfrecordsql.ErrNotFound:
		return true
	case twofactorsessionmgo.ErrNotFound:
		return true
	case twofactorsessionsql.ErrNotFound:
		return true
	case mocks.ErrNotFound:
		return true
	}
	return false
}
