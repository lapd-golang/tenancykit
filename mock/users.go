package mock

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tenancykit/db/mock/userdbbackendimpl"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/db/mock"
	"github.com/influx6/faux/context"
)

// UserBackend returns a mock implementation of the db.types.UserDBBackend interface.
func UserBackend() userdbbackendimpl.UserDBBackendImpl {
	var mocker userdbbackendimpl.UserDBBackendImpl

	db := make(map[string]tenancykit.User)

	mocker.CreateFunc = func(ctx context.Context, elem tenancykit.User) error {
		if _, exist := db[elem.PublicID]; exist {
			return errors.New("record already exists")
		}

		db[elem.PublicID] = elem
		return nil
	}

	mocker.DeleteFunc = func(ctx context.Context, publicID string) error {
		if _, exist := db[publicID]; !exist {
			return mock.ErrNotFound
		}
		delete(db, publicID)
		return nil
	}

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderby string) ([]tenancykit.User, error) {
		var records []tenancykit.User
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.User, int, error) {
		var records []tenancykit.User
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (tenancykit.User, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, mock.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (tenancykit.User, error) {
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

		return tenancykit.User{}, mock.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem tenancykit.User) error {
		if _, exist := db[publicID]; !exist {
			return mock.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
