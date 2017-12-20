// Package tokensql provides a auto-generated package which contains a sql CRUD API for the specific Token struct in package pkg.
//
//
package tokensql

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

// TokenFields defines an interface which exposes method to return a map of all
// attributes associated with the defined structure as decided by the structure.
type TokenFields interface {
	Fields() (map[string]interface{}, error)
}

// TokenConsumer defines an interface which accepts a map of data which will be consumed
// into the giving implementing structure as decided by the structure.
type TokenConsumer interface {
	Consume(map[string]interface{}) error
}

// TokenDB defines a structure which provide DB CRUD operations
// using sql as the underline db.
type TokenDB struct {
	col     string
	sx      sql.DB
	dx      *sql.SQL
	metrics metrics.Metrics
	table   db.TableIdentity
}

// New returns a new instance of TokenDB.
func New(table string, m metrics.Metrics, sx sql.DB, tm ...tables.TableMigration) *TokenDB {
	dx := sql.New(m, sx, tm...)

	return &TokenDB{
		sx:      sx,
		dx:      dx,
		col:     table,
		metrics: m,
		table:   db.TableName{Name: table},
	}
}

// Delete attempts to remove the record from the db using the provided publicID.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given pkg.Token struct.
func (mdb *TokenDB) Delete(ctx context.Context, publicID string) error {
	m := metrics.NewTrace("TokenDB.Delete")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Delete"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

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
// pkg.Token.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) Create(ctx context.Context, elem pkg.Token) error {
	m := metrics.NewTrace("TokenDB.Create")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Create"), metrics.With("publicID", elem.PublicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.With("publicID", elem.PublicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Save(mdb.table, elem); err != nil {
		mdb.metrics.Emit(
			metrics.Errorf("Failed to create Token record"),
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
func (mdb *TokenDB) GetAllByOrder(ctx context.Context, order string, orderby string) ([]pkg.Token, error) {
	res, _, err := mdb.GetAll(ctx, order, orderby, -1, -1)
	return res, err
}

// GetAll retrieves all records from the db and returns a slice of pkg.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) GetAll(ctx context.Context, order string, orderby string, page int, responsePerPage int) ([]pkg.Token, int, error) {
	m := metrics.NewTrace("TokenDB.GetAll")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.GetAll"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))

		return nil, -1, err
	}

	var total int
	var ritems []pkg.Token

	items, total, err := mdb.dx.GetAllPerPage(mdb.table, order, orderby, page, responsePerPage)
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Token type from db"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return nil, total, ErrNotFound
		}
		return nil, total, err
	}

	for _, item := range items {
		var elem pkg.Token
		if err := elem.Consume(item); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to consume record data for Token from db"), metrics.With("table", mdb.col), metrics.With("data", item), metrics.With("error", err.Error()))
			return nil, -1, err
		}
		ritems = append(ritems, elem)
	}

	return ritems, total, nil
}

// GetByField retrieves a record from the db using the field key and value,
// returns the pkg.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) GetByField(ctx context.Context, key string, value interface{}) (pkg.Token, error) {
	m := metrics.NewTrace("TokenDB.Get")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Get"), metrics.With(key, value), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("table", mdb.col), metrics.With("error", err.Error()))

		return pkg.Token{}, err
	}

	var elem pkg.Token

	if err := mdb.dx.Get(mdb.table, &elem, key, value); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Token type from db"),
			metrics.With("table", mdb.col),
			metrics.With("error", err.Error()))

		if err == dsql.ErrNoRows {
			return pkg.Token{}, ErrNotFound
		}

		return pkg.Token{}, err
	}

	return elem, nil
}

// Get retrieves a record from the db using the publicID and returns the pkg.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) Get(ctx context.Context, publicID string) (pkg.Token, error) {
	m := metrics.NewTrace("TokenDB.Get")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Get"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return pkg.Token{}, err
	}

	var elem pkg.Token

	if err := mdb.dx.Get(mdb.table, &elem, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Token type from db"),
			metrics.With("table", mdb.col),
			metrics.With("error", err.Error()))

		if err == dsql.ErrNoRows {
			return pkg.Token{}, ErrNotFound
		}
		return pkg.Token{}, err
	}

	return elem, nil
}

// Update uses a record from the db using the publicID and returns the pkg.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) Update(ctx context.Context, publicID string, elem pkg.Token) error {
	m := metrics.NewTrace("TokenDB.Update")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Update"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to finish, context has expired"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Update(mdb.table, elem, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update Token record"), metrics.With("query", elem), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("query", elem))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the table.
func (mdb *TokenDB) Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error {
	m := metrics.NewTrace("TokenDB.Exec")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Exec"), metrics.WithTrace(m.End()))

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
