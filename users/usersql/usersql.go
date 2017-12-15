// Package usersql provides a auto-generated package which contains a sql CRUD API for the specific User struct in package users.
//
//
package usersql

import (
	"errors"
	"fmt"

	"github.com/influx6/faux/db"
	"github.com/jmoiron/sqlx"

	"github.com/influx6/faux/db/sql"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/db/sql/tables"

	"github.com/gokit/tenancykit/users"
)

// mapFields defines a type for a map that exposes a Fields() method.
type mapFields map[string]interface{}

// Fields returns the map itself and provides a method to match the sql.TableField interface.
func (m mapFields) Fields() map[string]interface{} {
	return m
}

// UserFields defines an interface which exposes method to return a map of all
// attributes associated with the defined structure as decided by the structure.
type UserFields interface {
	Fields() map[string]interface{}
}

// UserConsumer defines an interface which accepts a map of data which will be consumed
// into the giving implementing structure as decided by the structure.
type UserConsumer interface {
	Consume(map[string]interface{}) error
}

// UserDB defines a structure which provide DB CRUD operations
// using sql as the underline db.
type UserDB struct {
	col     string
	sx      sql.DB
	dx      *sql.SQL
	metrics metrics.Metrics
	table   db.TableIdentity
}

// New returns a new instance of UserDB.
func New(table string, m metrics.Metrics, sx sql.DB, tm ...tables.TableMigration) *UserDB {
	dx := sql.New(m, sx, tm...)

	return &UserDB{
		sx:      sx,
		dx:      dx,
		col:     table,
		metrics: m,
		table:   db.TableName{Name: table},
	}
}

// Delete attempts to remove the record from the db using the provided publicID.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given users.User struct.
func (mdb *UserDB) Delete(ctx context.Context, publicID string) error {
	m := metrics.NewTrace("UserDB.Delete")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Delete"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.WithFields("publicID", publicID), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))
		return err
	}

	if err := mdb.dx.Delete(mdb.table, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.WithFields("table", mdb.col), metrics.WithFields("publicID", publicID), metrics.WithFields("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Deleted record"), metrics.WithFields("table", mdb.col), metrics.WithFields("publicID", publicID))

	return nil
}

// Create attempts to add the record into the db using the provided instance of the
// users.User.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) Create(ctx context.Context, elem users.User) error {
	m := metrics.NewTrace("UserDB.Create")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Create"), metrics.With("publicID", elem.PublicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.WithFields("publicID", elem.PublicID), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))
		return err
	}

	if fields, ok := interface{}(elem).(UserFields); ok {
		if err := mdb.dx.Save(mdb.table, mapFields(fields.Fields())); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to create User record"), metrics.WithFields("table", mdb.col), metrics.WithFields("elem", elem), metrics.WithFields("error", err.Error()))

			return err
		}

		mdb.metrics.Emit(metrics.Info("Create record"), metrics.WithFields("table", mdb.col), metrics.WithFields("elem", elem))

		return nil
	}

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")

		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.WithFields("publicID", elem.PublicID), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))
		return err
	}

	content := mapFields(map[string]interface{}{

		"created_at": elem.Created,

		"hash": elem.Hash,

		"private_id": elem.PrivateID,

		"public_id": elem.PublicID,

		"two_factor_auth": elem.TwoFactorAuth,

		"updated_at": elem.Updated,

		"username": elem.Username,
	})

	if err := mdb.dx.Save(mdb.table, content); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create User record"), metrics.WithFields("table", mdb.col), metrics.WithFields("query", content), metrics.WithFields("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Create record"), metrics.WithFields("table", mdb.col), metrics.WithFields("query", content))
	return nil
}

// GetAll retrieves all records from the db and returns a slice of users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) GetAll(ctx context.Context, order string, orderby string, page int, responsePerPage int) ([]users.User, int, error) {
	m := metrics.NewTrace("UserDB.GetAll")
	defer mdb.metrics.Emit(metrics.Info("UserDB.GetAll"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))

		return nil, -1, err
	}

	var ritems []users.User

	total, err := mdb.dx.GetAllPerPageBy(mdb.table, order, orderby, page, responsePerPage, func(rows *sqlx.Rows) error {
		for rows.Next() {
			var ritem users.User

			if err := rows.StructScan(&ritem); err != nil {
				mdb.metrics.Emit(metrics.Errorf(err), metrics.WithFields(metrics.Field{
					"err":   err,
					"table": mdb.table.Table(),
				}))

				return err
			}

			ritems = append(ritems, ritem)
		}

		return nil
	})

	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to consume record data for User from db"), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))

		return nil, total, err
	}

	return ritems, total, nil

}

// GetByField retrieves a record from the db using the field key and value,
// returns the users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) GetByField(ctx context.Context, key string, value string) (users.User, error) {
	m := metrics.NewTrace("UserDB.Get")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Get"), metrics.With(key, value), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.WithFields(key, value), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))

		return users.User{}, err
	}

	item, err := mdb.dx.Get(mdb.table, key, value)
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of User type from db"), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))

		return users.User{}, err
	}

	var elem users.User

	if _, ok := elem.(UserConsumer); !ok {
		return elem, errors.New("Only UserConsumer allowed")
	}

	if err := elem.Consume(item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to consume record data for User from db"), metrics.WithFields("table", mdb.col), metrics.WithFields("data", item), metrics.WithFields("error", err.Error()))

		return users.User{}, err
	}

	return elem, nil
}

// Get retrieves a record from the db using the publicID and returns the users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) Get(ctx context.Context, publicID string) (users.User, error) {
	m := metrics.NewTrace("UserDB.Get")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Get"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.WithFields("publicID", publicID), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))
		return users.User{}, err
	}

	var elem users.User

	if err := mdb.dx.GetBy(mdb.table, func(row *sqlx.Row) error {
		if err := row.StructScan(&elem); err != nil {
			mdb.metrics.Emit(metrics.Errorf(err), metrics.WithFields(metrics.Field{
				"err":   err,
				"table": mdb.table.Table(),
			}))

			return err
		}

		return nil
	}, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to consume record data for User from db"), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))

		return users.User{}, err
	}

	return elem, nil

}

// Update uses a record from the db using the publicID and returns the users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) Update(ctx context.Context, publicID string, elem users.User) error {
	m := metrics.NewTrace("UserDB.Update")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Update"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to finish, context has expired"), metrics.WithFields("table", mdb.col), metrics.WithFields("public_id", publicID), metrics.WithFields("error", err.Error()))
		return err
	}

	var data mapFields

	if fields, ok := interface{}(elem).(UserFields); ok {
		data = mapFields(fields.Fields())
	} else {
		data = mapFields(map[string]interface{}{

			"created_at": elem.Created,

			"hash": elem.Hash,

			"private_id": elem.PrivateID,

			"public_id": elem.PublicID,

			"two_factor_auth": elem.TwoFactorAuth,

			"updated_at": elem.Updated,

			"username": elem.Username,
		})
	}

	if err := mdb.dx.Update(mdb.table, data, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update User record"), metrics.WithFields("query", data), metrics.WithFields("table", mdb.col), metrics.WithFields("public_id", publicID), metrics.WithFields("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.WithFields("table", mdb.col), metrics.WithFields("public_id", publicID), metrics.WithFields("query", data))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the table.
func (mdb *UserDB) Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error {
	m := metrics.NewTrace("UserDB.Exec")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Exec"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))
		return err
	}

	if err := fx(mdb.dx, mdb.sx); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.WithFields("table", mdb.col), metrics.WithFields("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Operation executed"), metrics.WithFields("table", mdb.col))

	return nil
}
