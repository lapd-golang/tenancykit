package mock

import (
	"errors"
	"fmt"
	"strings"

	"context"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/mocks/userdbbackendimpl"
	"github.com/gokit/tenancykit/pkg/db/types"
)

// UserBackend returns a mock implementation of the db.types.UserDBBackend interface.
func UserBackend() types.UserDBBackend {
	var mocker userdbbackendimpl.UserDBBackendImpl

	db := make(map[string]pkg.User)
	mocker.CountFunc = func(ctx context.Context) (int, error) {
		return len(db), nil
	}

	mocker.CreateFunc = func(ctx context.Context, elem pkg.User) error {
		if _, exist := db[elem.PublicID]; exist {
			return errors.New("record already exists")
		}

		db[elem.PublicID] = elem
		return nil
	}

	mocker.DeleteFunc = func(ctx context.Context, publicID string) error {
		if _, exist := db[publicID]; !exist {
			return pkg.ErrNotFound
		}
		delete(db, publicID)
		return nil
	}

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderby string) ([]pkg.User, error) {
		var records []pkg.User
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.User, int, error) {
		var records []pkg.User
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (pkg.User, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, pkg.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (pkg.User, error) {
		val := fmt.Sprintf("%+s", value)
		for _, record := range db {
			switch strings.ToLower(key) {
			case "username":
				if record.Username == val {
					return record, nil
				}
			case "tenant_id":
				if record.TenantID == val {
					return record, nil
				}
			case "public_id":
				if record.PublicID == val {
					return record, nil
				}
			case "email":
				if record.Email == val {
					return record, nil
				}
			}
		}

		return pkg.User{}, pkg.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem pkg.User) error {
		if _, exist := db[publicID]; !exist {
			return pkg.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
