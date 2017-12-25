package mock

import (
	"errors"
	"fmt"
	"strings"

	"context"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/mocks/tfrecorddbbackendimpl"
	"github.com/gokit/tenancykit/pkg/db/types"
)

// TFRecordBackend returns a mock implementation of the db.types.TFRecordDBBackend interface.
func TFRecordBackend() types.TFRecordDBBackend {
	var mocker tfrecorddbbackendimpl.TFRecordDBBackendImpl

	db := make(map[string]pkg.TFRecord)

	mocker.CountFunc = func(ctx context.Context) (int, error) {
		return len(db), nil
	}
	mocker.CreateFunc = func(ctx context.Context, elem pkg.TFRecord) error {
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

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderby string) ([]pkg.TFRecord, error) {
		var records []pkg.TFRecord
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.TFRecord, int, error) {
		var records []pkg.TFRecord
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (pkg.TFRecord, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, pkg.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (pkg.TFRecord, error) {
		val := fmt.Sprintf("%+s", value)
		for _, record := range db {
			switch strings.ToLower(key) {
			case "domain":
				if record.Domain == val {
					return record, nil
				}
			case "public_id":
				if record.PublicID == val {
					return record, nil
				}
			case "user_id":
				if record.UserID == val {
					return record, nil
				}
			}
		}

		return pkg.TFRecord{}, pkg.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem pkg.TFRecord) error {
		if _, exist := db[publicID]; !exist {
			return pkg.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
