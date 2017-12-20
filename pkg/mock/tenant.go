package mock

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/mocks"
	"github.com/gokit/tenancykit/pkg/db/mocks/tenantdbbackendimpl"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/influx6/faux/context"
)

// TenantDBBackend returns a mock implementation of the db.types.TenantDBBackend interface.
func TenantDBBackend() types.TenantDBBackend {
	var mocker tenantdbbackendimpl.TenantDBBackendImpl

	db := make(map[string]pkg.Tenant)

	mocker.CreateFunc = func(ctx context.Context, elem pkg.Tenant) error {
		if _, exist := db[elem.PublicID]; exist {
			return errors.New("record already exists")
		}

		db[elem.PublicID] = elem
		return nil
	}

	mocker.DeleteFunc = func(ctx context.Context, publicID string) error {
		if _, exist := db[publicID]; !exist {
			return mocks.ErrNotFound
		}
		delete(db, publicID)
		return nil
	}

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderBy string) ([]pkg.Tenant, error) {
		var records []pkg.Tenant
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Tenant, int, error) {
		var records []pkg.Tenant
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (pkg.Tenant, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, mocks.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (pkg.Tenant, error) {
		val := fmt.Sprintf("%+s", value)
		for _, record := range db {
			switch strings.ToLower(key) {
			case "name":
				if record.Name == val {
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

		return pkg.Tenant{}, mocks.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem pkg.Tenant) error {
		if _, exist := db[publicID]; !exist {
			return mocks.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
