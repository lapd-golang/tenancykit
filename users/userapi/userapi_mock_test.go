package userapi_test

import (
	"errors"
	"fmt"

	"golang.org/x/sync/syncmap"

	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit/users"
)

// MockAPIOperator defines an structure which implements the APIOperator for providing
// a mock usage in tests for use with the Unconvertible Type http API.
type MockAPIOperator struct {
	store *syncmap.Map
}

// NewMockAPIOperator returns a new instance of a MockAPIOperator.
func NewMockAPIOperator() *MockAPIOperator {
	return &MockAPIOperator{
		store: new(syncmap.Map),
	}
}

// Delete provides the operation to remove a giving record identified by ID.
func (m *MockAPIOperator) Delete(ctx context.Context, publicID string) error {
	if _, has := m.store.Load(publicID); !has {
		return fmt.Errorf("Record does not exists with id %q", publicID)
	}

	m.store.Delete(publicID)
	return nil
}

// GetAll returns a slice of all available record of type users.User.
func (m *MockAPIOperator) GetAll(ctx context.Context, order string, orderBy string, page int, responserPerPage int) ([]users.User, int, error) {
	var records []users.User

	m.store.Range(func(k, v interface{}) bool {
		if elem, ok := v.(users.User); ok {
			records = append(records, elem)
		}

		return true
	})

	return records, len(records), nil
}

// Get retrieves a record based on the provided publicID.
func (m *MockAPIOperator) Get(ctx context.Context, publicID string) (users.User, error) {
	elem, found := m.store.Load(publicID)
	if !found {
		return users.User{}, fmt.Errorf("Record does not exists with id %q", publicID)
	}

	rElem, ok := elem.(users.User)
	if !ok {
		return users.User{}, errors.New("Record does not match type")
	}

	return rElem, nil
}

// Update updates a giving record with the given new value.
func (m *MockAPIOperator) Update(ctx context.Context, publicID string, elem users.User) error {

	m.store.Store(publicID, elem)
	return nil

}

// Create adds a new record into the giving record store.
func (m *MockAPIOperator) Create(ctx context.Context, elem users.User) (users.User, error) {

	m.store.Store(elem.PublicID, elem)
	return elem, nil

}
