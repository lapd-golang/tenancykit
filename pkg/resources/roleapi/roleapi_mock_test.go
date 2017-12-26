package roleapi_test

import (
	"errors"
	"fmt"
	"strings"

	"context"

	httpapi "github.com/gokit/tenancykit/pkg/resources/roleapi"

	"github.com/gokit/tenancykit/pkg"
)

// ErrNotFound is returned when a record is not found.
var ErrNotFound = errors.New("not found")

var _ httpapi.RoleBackend = (*mocker)(nil)

// mocker defines an structure which implements the APIOperator for providing
// a mock usage in tests for use with the "Role" http API.
type mocker struct {
	db map[string]pkg.Role
}

// newMocker returns a new instance of a mocker.
func newMocker() *mocker {
	return &mocker{
		db: map[string]pkg.Role{},
	}
}

// Count provides the operation to remove a giving record identified by ID.
func (m *mocker) Count(ctx context.Context) (int, error) {
	return len(m.db), nil
}

// Delete provides the operation to remove a giving record identified by ID.
func (m *mocker) Delete(ctx context.Context, publicID string) error {
	if _, exist := m.db[publicID]; !exist {
		return ErrNotFound
	}
	delete(m.db, publicID)
	return nil
}

// GetAllByOrder returns a slice of all available record of type pkg.Role.
func (m *mocker) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Role, error) {
	var records []pkg.Role
	for _, elem := range m.db {
		records = append(records, elem)
	}
	return records, nil
}

// GetAll returns a slice of all available record of type pkg.Role.
func (m *mocker) GetAll(ctx context.Context, order string, orderBy string, page int, responserPerPage int) ([]pkg.Role, int, error) {
	var records []pkg.Role
	for _, elem := range m.db {
		records = append(records, elem)
	}
	return records, len(records), nil
}

// GetByField retrieves a record based on the provided publicID.
func (m *mocker) GetByField(ctx context.Context, key string, value interface{}) (pkg.Role, error) {
	val := fmt.Sprintf("%+s", value)
	for _, record := range m.db {
		switch strings.ToLower(key) {
		case "public_id":
			if record.PublicID == val {
				return record, nil
			}
		}
	}

	return pkg.Role{}, ErrNotFound
}

// Get retrieves a record based on the provided publicID.
func (m *mocker) Get(ctx context.Context, publicID string) (pkg.Role, error) {
	elem, exist := m.db[publicID]
	if !exist {
		return elem, ErrNotFound
	}
	return elem, nil
}

// Update updates a giving record with the given new value.
func (m *mocker) Update(ctx context.Context, publicID string, elem pkg.Role) error {
	if _, exist := m.db[publicID]; !exist {
		return ErrNotFound
	}

	m.db[publicID] = elem
	return nil

}

// Create adds a new record into the giving record store.
func (m *mocker) Create(ctx context.Context, elem pkg.RoleName) (pkg.Role, error) {

	newElem, err := createRole(ctx, elem)
	if err != nil {
		return pkg.Role{}, err
	}
	m.db[newElem.PublicID] = newElem
	return newElem, nil

}

// createRole creates a new pkg.Role record from the given create type.
func createRole(ctx context.Context, elem pkg.RoleName) (pkg.Role, error) {
	return pkg.NewRole(elem), nil
}
