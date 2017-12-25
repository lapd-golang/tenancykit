package mock

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/mocks/identitydbbackendimpl"
	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"
)

func IdentityBackend() identitydbbackendimpl.IdentityDBBackendImpl {
	var mocker identitydbbackendimpl.IdentityDBBackendImpl

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
			return pkg.ErrNotFound
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
			return elem, pkg.ErrNotFound
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

		return userclaimjwt.Identity{}, pkg.ErrNotFound
	}

	mocker.UpdateFunc = func(ctx context.Context, publicID string, elem userclaimjwt.Identity) error {
		if _, exist := db[publicID]; !exist {
			return pkg.ErrNotFound
		}

		db[publicID] = elem
		return nil
	}

	return mocker
}
