// Package twofactorsessionsql provides a auto-generated package which contains a sql CRUD API for the specific TwoFactorSession struct in package pkg.
//
//
package twofactorsessionsql

import (
	"fmt"

	"errors"

	dsql "database/sql"

	"github.com/influx6/faux/db"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/db/sql"

	"github.com/influx6/faux/db/sql/tables"

	"github.com/gokit/tenancykit/pkg"
)

// errors ...
var (
	ErrNotFound = errors.New("record not found")
)

// mapFields defines a type for a map that exposes a Fields() method.
type mapFields map[string]interface{}

// Fields returns the map itself and provides a method to match the sql.TableField interface.
func (m mapFields) Fields() (map[string]interface{}, error) {
	return m, nil
}

// TwoFactorSessionFields defines an interface which exposes method to return a map of all
// attributes associated with the defined structure as decided by the structure.
type TwoFactorSessionFields interface {
	Fields() (map[string]interface{}, error)
}

// TwoFactorSessionConsumer defines an interface which accepts a map of data which will be consumed
// into the giving implementing structure as decided by the structure.
type TwoFactorSessionConsumer interface {
	Consume(map[string]interface{}) error
}

// TwoFactorSessionDB defines a structure which provide DB CRUD operations
// using sql as the underline db.
type TwoFactorSessionDB struct {
	col     string
	sx      sql.DB
	dx      *sql.SQL
	metrics metrics.Metrics
	table   db.TableIdentity
}

// New returns a new instance of TwoFactorSessionDB.
func New(table string, m metrics.Metrics, sx sql.DB, tm ...tables.TableMigration) *TwoFactorSessionDB {
	dx := sql.New(m, sx, tm...)

	return &TwoFactorSessionDB{
		sx:      sx,
		dx:      dx,
		col:     table,
		metrics: m,
		table:   db.TableName{Name: table},
	}
}

// Delete attempts to remove the record from the db using the provided publicID.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given pkg.TwoFactorSession struct.
func (mdb *TwoFactorSessionDB) Delete(ctx context.Context, publicID string) error {
	m := metrics.NewTrace("TwoFactorSessionDB.Delete")
	defer mdb.metrics.Emit(metrics.Info("TwoFactorSessionDB.Delete"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("publicID", publicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Delete(mdb.table, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("table", mdb.col), metrics.With("publicID", publicID), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Deleted record"), metrics.With("table", mdb.col), metrics.With("publicID", publicID))

	return nil
}

// Create attempts to add the record into the db using the provided instance of the
// pkg.TwoFactorSession.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given TwoFactorSession struct.
func (mdb *TwoFactorSessionDB) Create(ctx context.Context, elem pkg.TwoFactorSession) error {
	m := metrics.NewTrace("TwoFactorSessionDB.Create")
	defer mdb.metrics.Emit(metrics.Info("TwoFactorSessionDB.Create"), metrics.With("publicID", elem.PublicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.With("publicID", elem.PublicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Save(mdb.table, elem); err != nil {
		mdb.metrics.Emit(
			metrics.Errorf("Failed to create TwoFactorSession record"),
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
func (mdb *TwoFactorSessionDB) GetAllByOrder(ctx context.Context, order string, orderby string) ([]pkg.TwoFactorSession, error) {
	res, _, err := mdb.GetAll(ctx, order, orderby, -1, -1)
	return res, err
}

// GetAll retrieves all records from the db and returns a slice of pkg.TwoFactorSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given TwoFactorSession struct.
func (mdb *TwoFactorSessionDB) GetAll(ctx context.Context, order string, orderby string, page int, responsePerPage int) ([]pkg.TwoFactorSession, int, error) {
	m := metrics.NewTrace("TwoFactorSessionDB.GetAll")
	defer mdb.metrics.Emit(metrics.Info("TwoFactorSessionDB.GetAll"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))

		return nil, -1, err
	}

	var total int
	var ritems []pkg.TwoFactorSession

	items, total, err := mdb.dx.GetAllPerPage(mdb.table, order, orderby, page, responsePerPage)
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of TwoFactorSession type from db"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return nil, total, ErrNotFound
		}
		return nil, total, err
	}

	for _, item := range items {
		var elem pkg.TwoFactorSession
		if err := elem.Consume(item); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to consume record data for TwoFactorSession from db"), metrics.With("table", mdb.col), metrics.With("data", item), metrics.With("error", err.Error()))
			return nil, -1, err
		}
		ritems = append(ritems, elem)
	}

	return ritems, total, nil
}

// GetByField retrieves a record from the db using the field key and value,
// returns the pkg.TwoFactorSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given TwoFactorSession struct.
func (mdb *TwoFactorSessionDB) GetByField(ctx context.Context, key string, value interface{}) (pkg.TwoFactorSession, error) {
	m := metrics.NewTrace("TwoFactorSessionDB.Get")
	defer mdb.metrics.Emit(metrics.Info("TwoFactorSessionDB.Get"), metrics.With(key, value), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("table", mdb.col), metrics.With("error", err.Error()))

		return pkg.TwoFactorSession{}, err
	}

	var elem pkg.TwoFactorSession

	if err := mdb.dx.Get(mdb.table, &elem, key, value); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of TwoFactorSession type from db"),
			metrics.With("table", mdb.col),
			metrics.With("error", err.Error()))

		if err == dsql.ErrNoRows {
			return pkg.TwoFactorSession{}, ErrNotFound
		}

		return pkg.TwoFactorSession{}, err
	}

	return elem, nil
}

// Get retrieves a record from the db using the publicID and returns the pkg.TwoFactorSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given TwoFactorSession struct.
func (mdb *TwoFactorSessionDB) Get(ctx context.Context, publicID string) (pkg.TwoFactorSession, error) {
	m := metrics.NewTrace("TwoFactorSessionDB.Get")
	defer mdb.metrics.Emit(metrics.Info("TwoFactorSessionDB.Get"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return pkg.TwoFactorSession{}, err
	}

	var elem pkg.TwoFactorSession

	if err := mdb.dx.Get(mdb.table, &elem, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of TwoFactorSession type from db"),
			metrics.With("table", mdb.col),
			metrics.With("error", err.Error()))

		if err == dsql.ErrNoRows {
			return pkg.TwoFactorSession{}, ErrNotFound
		}
		return pkg.TwoFactorSession{}, err
	}

	return elem, nil
}

// Update uses a record from the db using the publicID and returns the pkg.TwoFactorSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given TwoFactorSession struct.
func (mdb *TwoFactorSessionDB) Update(ctx context.Context, publicID string, elem pkg.TwoFactorSession) error {
	m := metrics.NewTrace("TwoFactorSessionDB.Update")
	defer mdb.metrics.Emit(metrics.Info("TwoFactorSessionDB.Update"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to finish, context has expired"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Update(mdb.table, elem, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update TwoFactorSession record"), metrics.With("query", elem), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("query", elem))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the table.
func (mdb *TwoFactorSessionDB) Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error {
	m := metrics.NewTrace("TwoFactorSessionDB.Exec")
	defer mdb.metrics.Emit(metrics.Info("TwoFactorSessionDB.Exec"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := fx(mdb.dx, mdb.sx); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Operation executed"), metrics.With("table", mdb.col))

	return nil
}
