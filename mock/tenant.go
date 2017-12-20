package mock

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tenancykit/db/types"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/db/mock"
	"github.com/gokit/tenancykit/db/mock/tenantdbbackendimpl"
	"github.com/influx6/faux/context"
)

// TenantDBBackend returns a mock implementation of the db.types.TenantDBBackend interface.
func TenantDBBackend() types.TenantDBBackend {
	var mocker tenantdbbackendimpl.TenantDBBackendImpl

	db := make(map[string]tenancykit.Tenant)

	mocker.CreateFunc = func(ctx context.Context, elem tenancykit.Tenant) error {
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

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderBy string) ([]tenancykit.Tenant, error) {
		var records []tenancykit.Tenant
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.Tenant, int, error) {
		var records []tenancykit.Tenant
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (tenancykit.Tenant, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, mock.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (tenancykit.Tenant, error) {
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

		return tenancykit.Tenant{}, mock.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem tenancykit.Tenant) error {
		if _, exist := db[publicID]; !exist {
			return mock.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
