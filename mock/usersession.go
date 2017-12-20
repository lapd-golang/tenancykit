package mock

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tenancykit/db/types"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/db/mock"
	"github.com/gokit/tenancykit/db/mock/usersessiondbbackendimpl"
	"github.com/influx6/faux/context"
)

// UserSessionBackend returns a mock implementation of the db.types.UserSessionDBBackend interface.
func UserSessionBackend() types.UserSessionDBBackend {
	var mocker usersessiondbbackendimpl.UserSessionDBBackendImpl

	db := make(map[string]tenancykit.UserSession)

	mocker.CreateFunc = func(ctx context.Context, elem tenancykit.UserSession) error {
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

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderby string) ([]tenancykit.UserSession, error) {
		var records []tenancykit.UserSession
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tenancykit.UserSession, int, error) {
		var records []tenancykit.UserSession
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (tenancykit.UserSession, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, mock.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (tenancykit.UserSession, error) {
		val := fmt.Sprintf("%+s", value)
		for _, record := range db {
			switch strings.ToLower(key) {
			case "user_id":
				if record.UserID == val {
					return record, nil
				}
			case "public_id":
				if record.PublicID == val {
					return record, nil
				}
			case "token":
				if record.Token == val {
					return record, nil
				}
			}
		}

		return tenancykit.UserSession{}, mock.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem tenancykit.UserSession) error {
		if _, exist := db[publicID]; !exist {
			return mock.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
