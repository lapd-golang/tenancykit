package mock

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tenancykit/tokenset/tokens/types"

	"github.com/gokit/tenancykit/db/mock"
	"github.com/gokit/tenancykit/tokenset/tokens"
	"github.com/gokit/tenancykit/tokenset/tokens/mocks/tokendbbackendimpl"
	"github.com/influx6/faux/context"
)

// TokenSetBackend returns a new instance of the tokendbbackendimpl for mocking.
func TokenSetBackend() types.TokenDBBackend {
	var mocker tokendbbackendimpl.TokenDBBackendImpl

	db := make(map[string]tokens.Token)

	mocker.CreateFunc = func(ctx context.Context, elem tokens.Token) error {
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

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderby string) ([]tokens.Token, error) {
		var records []tokens.Token
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tokens.Token, int, error) {
		var records []tokens.Token
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (tokens.Token, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, mock.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (tokens.Token, error) {
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

		return tokens.Token{}, mock.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem tokens.Token) error {
		if _, exist := db[publicID]; !exist {
			return mock.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
