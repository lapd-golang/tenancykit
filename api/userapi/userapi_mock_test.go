package userapi_test

import (
	"errors"
	"fmt"

	"golang.org/x/sync/syncmap"

	"github.com/influx6/faux/context"

	"github.com/gokit/tenancykit"
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

// GetAll returns a slice of all available record of type tenancykit.User.
func (m *MockAPIOperator) GetAll(ctx context.Context, order string, orderBy string, page int, responserPerPage int) ([]tenancykit.User, int, error) {
	var records []tenancykit.User

	m.store.Range(func(k, v interface{}) bool {
		if elem, ok := v.(tenancykit.User); ok {
			records = append(records, elem)
		}

		return true
	})

	return records, len(records), nil
}

// Get retrieves a record based on the provided publicID.
func (m *MockAPIOperator) Get(ctx context.Context, publicID string) (tenancykit.User, error) {
	elem, found := m.store.Load(publicID)
	if !found {
		return tenancykit.User{}, fmt.Errorf("Record does not exists with id %q", publicID)
	}

	rElem, ok := elem.(tenancykit.User)
	if !ok {
		return tenancykit.User{}, errors.New("Record does not match type")
	}

	return rElem, nil
}

// Update updates a giving record with the given new value.
func (m *MockAPIOperator) Update(ctx context.Context, publicID string, elem tenancykit.UpdateUser) error {

	nowElem, ok := m.store.Load(publicID).(tenancykit.User)
	if !ok {
		return errors.New("Not found/invalid type")
	}

	newElem, err := updateUser(ctx, nowElem, elem)
	if err != nil {
		return err
	}

	m.store.Store(newElem.PublicID, newElem)

	return nil

}

// Create adds a new record into the giving record store.
func (m *MockAPIOperator) Create(ctx context.Context, elem tenancykit.CreateUser) (tenancykit.User, error) {

	newElem, err := createUser(ctx, elem)
	if err != nil {
		return tenancykit.User{}, err
	}

	m.store.Store(newElem.PublicID, newElem)

	return newElem, nil

}

// updateUser updates a exiting tenancykit.User record with the given update item.
func UpdateUser(ctx context.Context, record tenancykit.User, updater tenancykit.UpdateUser) (tenancykit.User, error) {

	// TODO(developer):
	// Override function contents with what should happen.

	return tenancykit.User{}, errors.New("Not Implemented")
}

// createUser creates a new tenancykit.User record from the given create type.
func createUser(ctx context.Context, elem tenancykit.CreateUser) (tenancykit.User, error) {

	// TODO(developer):
	// Override function contents with what should happen.

	return tenancykit.User{}, errors.New("Not Implemented")
}
