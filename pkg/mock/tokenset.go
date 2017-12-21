package mock

import (
	"errors"
	"fmt"
	"strings"

	"context"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/mocks"
	"github.com/gokit/tenancykit/pkg/db/mocks/tokendbbackendimpl"
	"github.com/gokit/tenancykit/pkg/db/types"
)

// TokenSetBackend returns a new instance of the tokendbbackendimpl for mocking.
func TokenSetBackend() types.TokenDBBackend {
	var mocker tokendbbackendimpl.TokenDBBackendImpl

	db := make(map[string]pkg.Token)

	mocker.CountFunc = func(ctx context.Context) (int, error) {
		return len(db), nil
	}
	mocker.CreateFunc = func(ctx context.Context, elem pkg.Token) error {
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

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderby string) ([]pkg.Token, error) {
		var records []pkg.Token
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Token, int, error) {
		var records []pkg.Token
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (pkg.Token, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, mocks.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (pkg.Token, error) {
		val := fmt.Sprintf("%+s", value)
		for _, record := range db {
			switch strings.ToLower(key) {
			case "target_id":
				if record.TargetID == val {
					return record, nil
				}
			case "public_id":
				if record.PublicID == val {
					return record, nil
				}
			case "value":
				if record.Value == val {
					return record, nil
				}
			}
		}

		return pkg.Token{}, mocks.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem pkg.Token) error {
		if _, exist := db[publicID]; !exist {
			return mocks.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
