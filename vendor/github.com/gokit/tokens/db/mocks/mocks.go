package mocks

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tokens"
	"github.com/gokit/tokens/db/types"
	"github.com/gokit/tokens/mock/tokendbbackendimpl"
	"github.com/gokit/tokens/mock/tokensetimpl"
)

// TokenBackendImpl implenents types.TokenBackend.
type TokenBackendImpl struct {
	types.TokenDBBackend
}

// Create creates token with db and returns instance.
func (tb TokenBackendImpl) Create(ctx context.Context, t tokens.Token) (tokens.Token, error) {
	return t, tb.TokenDBBackend.Create(ctx, t)
}

// TokenBackend returns a new instance of the tokendbbackendimpl for mocking.
func TokenBackend() types.TokenDBBackend {
	var mocker tokendbbackendimpl.TokenDBBackendImpl

	db := make(map[string]tokens.Token)

	mocker.CountFunc = func(ctx context.Context) (int, error) {
		return len(db), nil
	}
	mocker.CreateFunc = func(ctx context.Context, elem tokens.Token) error {
		if _, exist := db[elem.PublicID]; exist {
			return errors.New("record already exists")
		}

		db[elem.PublicID] = elem
		return nil
	}

	mocker.DeleteFunc = func(ctx context.Context, publicID string) error {
		if _, exist := db[publicID]; !exist {
			return tokens.ErrNotFound
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
			return elem, tokens.ErrNotFound
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

		return tokens.Token{}, tokens.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem tokens.Token) error {
		if _, exist := db[publicID]; !exist {
			return tokens.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}

// TokenSetBackend returns a new instance of the tokens.TokenSet for mocking.
func TokenSetBackend() tokens.TokenSet {
	var mocker tokensetimpl.TokenSetImpl

	db := map[string][]tokens.Token{}

	mocker.HasFunc = func(ctx context.Context, id string, token string) (bool, error) {
		if set, ok := db[id]; ok {
			for _, item := range set {
				if item.Value != token {
					continue
				}

				return true, nil
			}
		}

		return false, nil
	}

	mocker.RemoveFunc = func(ctx context.Context, id string, token string) error {
		if set, ok := db[id]; ok {
			for index, item := range set {
				if item.Value != token {
					continue
				}

				set = append(set[:index], set[index+1:]...)
				db[id] = set
				return nil
			}
		}

		return tokens.ErrNotFound
	}

	mocker.AddFunc = func(ctx context.Context, id string, token string) (tokens.Token, error) {
		if set, ok := db[id]; ok {
			for _, item := range set {
				if item.Value == token {
					return item, nil
				}
			}

			tok := tokens.NewToken(id, token)
			db[id] = append(set, tok)
			return tok, nil
		}

		tok := tokens.NewToken(id, token)
		db[id] = []tokens.Token{tok}
		return tok, nil
	}

	return mocker
}
