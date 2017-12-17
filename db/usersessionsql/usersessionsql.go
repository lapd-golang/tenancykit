// Package usersessionsql provides a auto-generated package which contains a sql CRUD API for the specific UserSession struct in package tenancykit.
//
//
package usersessionsql

import (
	"fmt"

	"github.com/influx6/faux/db"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/db/sql"

	"github.com/influx6/faux/db/sql/tables"

	"github.com/gokit/tenancykit"
)

// mapFields defines a type for a map that exposes a Fields() method.
type mapFields map[string]interface{}

// Fields returns the map itself and provides a method to match the sql.TableField interface.
func (m mapFields) Fields() (map[string]interface{}, error) {
	return m, nil
}

// UserSessionFields defines an interface which exposes method to return a map of all
// attributes associated with the defined structure as decided by the structure.
type UserSessionFields interface {
	Fields() (map[string]interface{}, error)
}

// UserSessionConsumer defines an interface which accepts a map of data which will be consumed
// into the giving implementing structure as decided by the structure.
type UserSessionConsumer interface {
	Consume(map[string]interface{}) error
}

// UserSessionDB defines a structure which provide DB CRUD operations
// using sql as the underline db.
type UserSessionDB struct {
	col     string
	sx      sql.DB
	dx      *sql.SQL
	metrics metrics.Metrics
	table   db.TableIdentity
}

// New returns a new instance of UserSessionDB.
func New(table string, m metrics.Metrics, sx sql.DB, tm ...tables.TableMigration) *UserSessionDB {
	dx := sql.New(m, sx, tm...)

	return &UserSessionDB{
		sx:      sx,
		dx:      dx,
		col:     table,
		metrics: m,
		table:   db.TableName{Name: table},
	}
}

// Delete attempts to remove the record from the db using the provided publicID.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given tenancykit.UserSession struct.
func (mdb *UserSessionDB) Delete(ctx context.Context, publicID string) error {
	m := metrics.NewTrace("UserSessionDB.Delete")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Delete"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("publicID", publicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Delete(mdb.table, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("table", mdb.col), metrics.With("publicID", publicID), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Deleted record"), metrics.With("table", mdb.col), metrics.With("publicID", publicID))

	return nil
}

// Create attempts to add the record into the db using the provided instance of the
// tenancykit.UserSession.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) Create(ctx context.Context, elem tenancykit.UserSession) error {
	m := metrics.NewTrace("UserSessionDB.Create")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Create"), metrics.With("publicID", elem.PublicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.With("publicID", elem.PublicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Save(mdb.table, elem); err != nil {
		mdb.metrics.Emit(
			metrics.Errorf("Failed to create UserSession record"),
			metrics.With("table", mdb.col),
			metrics.With("elem", elem),
			metrics.With("error", err.Error()),
		)

		return err
	}

	mdb.metrics.Emit(metrics.Info("Create record"),
		metrics.With("table", mdb.col),
		metrics.With("elem", elem))

	return nil
}

// GetAllByOrder attempts to retrieve all elements from db using provided order and orderby
// values.
func (mdb *UserSessionDB) GetAllByOrder(ctx context.Context, order string, orderby string) ([]tenancykit.UserSession, int, error) {
	return mdb.GetAll(ctx, order, orderby, -1, -1)
}

// GetAll retrieves all records from the db and returns a slice of tenancykit.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) GetAll(ctx context.Context, order string, orderby string, page int, responsePerPage int) ([]tenancykit.UserSession, int, error) {
	m := metrics.NewTrace("UserSessionDB.GetAll")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.GetAll"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))

		return nil, -1, err
	}

	var total int
	var ritems []tenancykit.UserSession

	items, total, err := mdb.dx.GetAllPerPage(mdb.table, order, orderby, page, responsePerPage)
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of UserSession type from db"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return nil, total, err
	}

	for _, item := range items {
		var elem tenancykit.UserSession
		if err := elem.Consume(item); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to consume record data for UserSession from db"), metrics.With("table", mdb.col), metrics.With("data", item), metrics.With("error", err.Error()))
			return nil, -1, err
		}
		ritems = append(ritems, elem)
	}

	return ritems, total, nil
}

// GetByField retrieves a record from the db using the field key and value,
// returns the tenancykit.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) GetByField(ctx context.Context, key string, value string) (tenancykit.UserSession, error) {
	m := metrics.NewTrace("UserSessionDB.Get")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Get"), metrics.With(key, value), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("table", mdb.col), metrics.With("error", err.Error()))

		return tenancykit.UserSession{}, err
	}

	var elem tenancykit.UserSession

	if err := mdb.dx.Get(mdb.table, &elem, key, value); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of UserSession type from db"),
			metrics.With("table", mdb.col),
			metrics.With("error", err.Error()))

		return tenancykit.UserSession{}, err
	}

	return elem, nil
}

// Get retrieves a record from the db using the publicID and returns the tenancykit.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) Get(ctx context.Context, publicID string) (tenancykit.UserSession, error) {
	m := metrics.NewTrace("UserSessionDB.Get")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Get"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return tenancykit.UserSession{}, err
	}

	var elem tenancykit.UserSession

	if err := mdb.dx.Get(mdb.table, &elem, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of UserSession type from db"),
			metrics.With("table", mdb.col),
			metrics.With("error", err.Error()))

		return tenancykit.UserSession{}, err
	}

	return elem, nil
}

// Update uses a record from the db using the publicID and returns the tenancykit.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) Update(ctx context.Context, publicID string, elem tenancykit.UserSession) error {
	m := metrics.NewTrace("UserSessionDB.Update")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Update"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to finish, context has expired"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Update(mdb.table, elem, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update UserSession record"), metrics.With("query", elem), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("query", elem))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the table.
func (mdb *UserSessionDB) Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error {
	m := metrics.NewTrace("UserSessionDB.Exec")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Exec"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := fx(mdb.dx, mdb.sx); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Operation executed"), metrics.With("table", mdb.col))

	return nil
}
