package mock

import (
	"fmt"

	"strings"

	"errors"

	"context"

	"github.com/gokit/tokens"

	userclaimjwt "github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"
)

// IdentityBackend returns a new instance of userclaimjwt.IdentityDBBackend
func IdentityBackend() userclaimjwt.IdentityDBBackend {
	var mocker IdentityDB

	db := make(map[string]userclaimjwt.Identity)

	mocker.CountFunc = func(ctx context.Context) (int, error) {
		return len(db), nil
	}

	mocker.CreateFunc = func(ctx context.Context, elem userclaimjwt.Identity) error {
		if _, exist := db[elem.PublicID]; exist {
			return errors.New("record already exists")
		}

		db[elem.PublicID] = elem
		return nil
	}

	mocker.DeleteFunc = func(ctx context.Context, publicID string) error {
		if _, exist := db[publicID]; !exist {
			return userclaimjwt.ErrNotFound
		}
		delete(db, publicID)
		return nil
	}

	mocker.GetAllByOrderFunc = func(ctx context.Context, order string, orderBy string) ([]userclaimjwt.Identity, error) {
		var records []userclaimjwt.Identity
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, nil
	}

	mocker.GetAllFunc = func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]userclaimjwt.Identity, int, error) {
		var records []userclaimjwt.Identity
		for _, elem := range db {
			records = append(records, elem)
		}
		return records, len(records), nil
	}

	mocker.GetFunc = func(ctx context.Context, publicID string) (userclaimjwt.Identity, error) {
		elem, exist := db[publicID]
		if !exist {
			return elem, userclaimjwt.ErrNotFound
		}
		return elem, nil
	}

	mocker.GetByFieldFunc = func(ctx context.Context, key string, value interface{}) (userclaimjwt.Identity, error) {
		val := fmt.Sprintf("%+s", value)
		for _, record := range db {
			switch strings.ToLower(key) {
			case "refresh_token":
				if record.RefreshToken == val {
					return record, nil
				}
			case "target_id":
				if record.TargetID == val {
					return record, nil
				}
			case "public_id":
				if record.PublicID == val {
					return record, nil
				}
			}
		}

		return userclaimjwt.Identity{}, userclaimjwt.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem userclaimjwt.Identity) error {
		if _, exist := db[publicID]; !exist {
			return userclaimjwt.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}

// IdentityDB defines a concrete struct which implements the methods for the
// userclaimjwt.IdentityDBBackend interface. All methods will panic, so add necessary internal logic.
type IdentityDB struct {
	CountFunc func(ctx context.Context) (int, error)

	DeleteFunc func(ctx context.Context, publicID string) error

	CreateFunc func(ctx context.Context, elem userclaimjwt.Identity) error

	GetFunc func(ctx context.Context, publicID string) (userclaimjwt.Identity, error)

	UpdateFunc func(ctx context.Context, publicID string, elem userclaimjwt.Identity) error

	GetAllByOrderFunc func(ctx context.Context, order string, orderBy string) ([]userclaimjwt.Identity, error)

	GetByFieldFunc func(ctx context.Context, key string, value interface{}) (userclaimjwt.Identity, error)

	GetAllFunc func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]userclaimjwt.Identity, int, error)
}

// Count implements the IdentityDBBackend.Count() method for IdentityDB.
func (impl IdentityDB) Count(ctx context.Context) (int, error) {

	ret1, ret2 := impl.CountFunc(ctx)
	return ret1, ret2

}

// Delete implements the IdentityDBBackend.Delete() method for IdentityDB.
func (impl IdentityDB) Delete(ctx context.Context, publicID string) error {

	ret1 := impl.DeleteFunc(ctx, publicID)
	return ret1

}

// Create implements the IdentityDBBackend.Create() method for IdentityDB.
func (impl IdentityDB) Create(ctx context.Context, elem userclaimjwt.Identity) error {

	ret1 := impl.CreateFunc(ctx, elem)
	return ret1

}

// Get implements the IdentityDBBackend.Get() method for IdentityDB.
func (impl IdentityDB) Get(ctx context.Context, publicID string) (userclaimjwt.Identity, error) {

	ret1, ret2 := impl.GetFunc(ctx, publicID)
	return ret1, ret2

}

// Update implements the IdentityDBBackend.Update() method for IdentityDB.
func (impl IdentityDB) Update(ctx context.Context, publicID string, elem userclaimjwt.Identity) error {

	ret1 := impl.UpdateFunc(ctx, publicID, elem)
	return ret1

}

// GetAllByOrder implements the IdentityDBBackend.GetAllByOrder() method for IdentityDB.
func (impl IdentityDB) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]userclaimjwt.Identity, error) {

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)
	return ret1, ret2

}

// GetByField implements the IdentityDBBackend.GetByField() method for IdentityDB.
func (impl IdentityDB) GetByField(ctx context.Context, key string, value interface{}) (userclaimjwt.Identity, error) {

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)
	return ret1, ret2

}

// GetAll implements the IdentityDBBackend.GetAll() method for IdentityDB.
func (impl IdentityDB) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]userclaimjwt.Identity, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)
	return ret1, ret2, ret3

}

// TokenBackend returns a in-memory token set.
func TokenBackend() tokens.TokenSet {
	var mocker tokenSetImpl

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

// TokenSetImpl defines a concrete struct which implements the methods for the
// TokenSet interface. All methods will panic, so add necessary internal logic.
type tokenSetImpl struct {
	RemoveFunc func(ctx context.Context, target string, token string) error

	HasFunc func(ctx context.Context, target string, token string) (bool, error)

	AddFunc func(ctx context.Context, target string, token string) (tokens.Token, error)
}

// Remove implements the TokenSet.Remove() method for TokenSetImpl.
func (impl tokenSetImpl) Remove(ctx context.Context, target string, token string) error {

	ret1 := impl.RemoveFunc(ctx, target, token)
	return ret1

}

// Has implements the TokenSet.Has() method for TokenSetImpl.
func (impl tokenSetImpl) Has(ctx context.Context, target string, token string) (bool, error) {

	ret1, ret2 := impl.HasFunc(ctx, target, token)
	return ret1, ret2

}

// Add implements the TokenSet.Add() method for TokenSetImpl.
func (impl tokenSetImpl) Add(ctx context.Context, target string, token string) (tokens.Token, error) {

	ret1, ret2 := impl.AddFunc(ctx, target, token)
	return ret1, ret2

}
