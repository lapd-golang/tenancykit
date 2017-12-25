package mock

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/mocks/activitydbbackendimpl"
	"github.com/gokit/tenancykit/pkg/db/types"
)

// ActivityBackend returns an implementation for types.ActivityDBBackend.
func ActivityBackend() types.ActivityDBBackend {
	var mocker activitydbbackendimpl.ActivityDBBackendImpl

	db := make(map[string]pkg.Activity)

	mocker.CountFunc = func(ctx context.Context) (int, error) {
		return len(db), nil
	}

	mocker.CreateFunc = func(ctx context.Context, elem pkg.Activity) error {
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

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderBy string) ([]pkg.Activity, error) {
		var records []pkg.Activity
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Activity, int, error) {
		var records []pkg.Activity
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (pkg.Activity, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, pkg.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (pkg.Activity, error) {
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
			}
		}

		return pkg.Activity{}, pkg.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem pkg.Activity) error {
		if _, exist := db[publicID]; !exist {
			return pkg.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
