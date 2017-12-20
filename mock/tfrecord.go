package mock

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tenancykit/db/types"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/db/mock"
	"github.com/gokit/tenancykit/db/mock/tfrecorddbbackendimpl"
	"github.com/influx6/faux/context"
)

// TFRecordBackend returns a mock implementation of the db.types.TFRecordDBBackend interface.
func TFRecordBackend() types.TFRecordDBBackend {
	var mocker tfrecorddbbackendimpl.TFRecordDBBackendImpl

	db := make(map[string]tenancykit.TFRecord)

	mocker.CreateFunc = func(ctx context.Context, elem tenancykit.TFRecord) error {
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

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderby string) ([]tenancykit.TFRecord, error) {
		var records []tenancykit.TFRecord
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.TFRecord, int, error) {
		var records []tenancykit.TFRecord
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (tenancykit.TFRecord, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, mock.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (tenancykit.TFRecord, error) {
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
			case "tenant_id":
				if record.TenantID == val {
					return record, nil
				}
			}
		}

		return tenancykit.TFRecord{}, mock.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem tenancykit.TFRecord) error {
		if _, exist := db[publicID]; !exist {
			return mock.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
