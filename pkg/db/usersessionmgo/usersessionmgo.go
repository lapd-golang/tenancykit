// Package usersessionmgo provides a auto-generated package which contains a sql CRUD API for the specific UserSession struct in package pkg.
//
//
package usersessionmgo

import (
	"errors"
	"fmt"

	"strings"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/metrics"

	"github.com/gokit/tenancykit/pkg"
)

// errors ...
var (
	ErrNotFound = errors.New("record not found")
)

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

// Mongod defines a interface which exposes a method for retrieving a
// mongo.Database and mongo.Session.
type Mongod interface {
	New() (*mgo.Database, *mgo.Session, error)
}

// UserSessionDB defines a structure which provide DB CRUD operations
// using mongo as the underline db.
type UserSessionDB struct {
	col             string
	db              Mongod
	metrics         metrics.Metrics
	ensuredIndex    bool
	incompleteIndex bool
	indexes         []mgo.Index
}

// New returns a new instance of UserSessionDB.
func New(col string, m metrics.Metrics, mo Mongod, indexes ...mgo.Index) *UserSessionDB {
	return &UserSessionDB{
		db:      mo,
		col:     col,
		metrics: m,
		indexes: indexes,
	}
}

// ensureIndex attempts to ensure all provided indexes into the specific collection.
func (mdb *UserSessionDB) ensureIndex() error {
	if mdb.ensuredIndex {
		return nil
	}

	m := metrics.NewTrace("UserSessionDB.ensureIndex")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.ensureIndex"), metrics.WithTrace(m.End()))

	if len(mdb.indexes) == 0 {
		return nil
	}

	// If we had an error before index was complete, then skip, we cant not
	// stop all ops because of failed index.
	if !mdb.ensuredIndex && mdb.incompleteIndex {
		return nil
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session for index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	defer session.Close()

	collection := database.C(mdb.col)

	for _, index := range mdb.indexes {
		if err := collection.EnsureIndex(index); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to ensure session index"), metrics.With("collection", mdb.col), metrics.With("index", index), metrics.With("error", err.Error()))

			mdb.incompleteIndex = true
			return err
		}

		mdb.metrics.Emit(metrics.Info("Succeeded in ensuring collection index"), metrics.With("collection", mdb.col), metrics.With("index", index))
	}

	mdb.ensuredIndex = true

	mdb.metrics.Emit(metrics.Info("Finished adding index"), metrics.With("collection", mdb.col))
	return nil
}

// Count attempts to return the total number of record from the db.
func (mdb *UserSessionDB) Count(ctx context.Context) (int, error) {
	m := metrics.NewTrace("UserSessionDB.Count")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Count"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")

		mdb.metrics.Emit(metrics.Errorf("Failed to get record count"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return -1, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return -1, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to get record count"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return -1, err
	}

	defer session.Close()

	query := bson.M{}
	total, err := database.C(mdb.col).Find(query).Count()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to get record count"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))

		return -1, err
	}

	mdb.metrics.Emit(metrics.Info("Deleted record"), metrics.With("collection", mdb.col), metrics.With("query", query))

	return total, err
}

// Delete attempts to remove the record from the db using the provided publicID.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given pkg.UserSession struct.
func (mdb *UserSessionDB) Delete(ctx context.Context, publicID string) error {
	m := metrics.NewTrace("UserSessionDB.Delete")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Delete"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	defer session.Close()

	query := bson.M{
		"publicID": publicID,
	}

	if err := database.C(mdb.col).Remove(query); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("publicID", publicID), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Deleted record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("publicID", publicID))

	return nil
}

// Create attempts to add the record into the db using the provided instance of the
// pkg.UserSession.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) Create(ctx context.Context, elem pkg.UserSession) error {
	m := metrics.NewTrace("UserSessionDB.Create")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Create"), metrics.With("publicID", elem.PublicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.With("publicID", elem.PublicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("publicID", elem.PublicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	defer session.Close()

	fields, err := elem.Fields()
	if err != nil {
		mdb.metrics.Emit(
			metrics.Errorf("Failed to get Fields() for UserSession record"),
			metrics.With("collection", mdb.col),
			metrics.With("elem", elem),
			metrics.With("error", err.Error()),
		)
		return err
	}

	if err := database.C(mdb.col).Insert(bson.M(fields)); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create UserSession record"), metrics.With("collection", mdb.col), metrics.With("elem", elem), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Create record"), metrics.With("collection", mdb.col), metrics.With("elem", elem))

	return nil
}

// GetAll retrieves all records from the db and returns a slice of pkg.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.UserSession, int, error) {
	m := metrics.NewTrace("UserSessionDB.GetAll")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.GetAll"), metrics.WithTrace(m.End()))

	switch strings.ToLower(order) {
	case "dsc", "desc":
		orderBy = "-" + orderBy
	}

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, -1, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, -1, err
	}

	if page <= 0 && responsePerPage <= 0 {
		records, err := mdb.GetAllByOrder(ctx, order, orderBy)
		return records, len(records), err
	}

	// Get total number of records.
	totalRecords, err := mdb.Count(ctx)
	if err != nil {
		return nil, -1, err
	}

	var totalWanted, indexToStart int

	if page <= 1 && responsePerPage > 0 {
		totalWanted = responsePerPage
		indexToStart = 0
	} else {
		totalWanted = responsePerPage * page
		indexToStart = totalWanted / 2

		if page > 1 {
			indexToStart++
		}
	}

	mdb.metrics.Emit(
		metrics.Info("DB:Query:GetAllPerPage"),
		metrics.WithFields(metrics.Field{
			"starting_index":       indexToStart,
			"total_records_wanted": totalWanted,
			"order":                order,
			"orderBy":              orderBy,
			"page":                 page,
			"responsePerPage":      responsePerPage,
		}),
	)

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, -1, err
	}

	defer session.Close()

	query := bson.M{}

	var ritems []pkg.UserSession

	var ditems []map[string]interface{}
	if err := database.C(mdb.col).Find(query).Skip(indexToStart).Limit(totalWanted).Sort(orderBy).All(&ditems); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of UserSession type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return nil, -1, ErrNotFound
		}
		return nil, -1, err
	}

	for _, item := range ditems {
		var elem pkg.UserSession
		if err := elem.Consume(item); err != nil {
			return nil, -1, err
		}
		ritems = append(ritems, elem)
	}

	return ritems, totalRecords, nil
}

// GetAllByOrder retrieves all records from the db and returns a slice of pkg.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) GetAllByOrder(ctx context.Context, order, orderBy string) ([]pkg.UserSession, error) {
	m := metrics.NewTrace("UserSessionDB.GetAllByOrder")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.GetAllByOrder"), metrics.WithTrace(m.End()))

	switch strings.ToLower(order) {
	case "dsc", "desc":
		orderBy = "-" + orderBy
	}

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")

		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, err
	}

	defer session.Close()

	query := bson.M{}

	var ditems []map[string]interface{}
	if err := database.C(mdb.col).Find(query).Sort(orderBy).All(&ditems); err != nil {
		mdb.metrics.Emit(
			metrics.Errorf("Failed to retrieve all records of UserSession type from db"),
			metrics.With("collection", mdb.col),
			metrics.With("query", query),
			metrics.With("error", err.Error()),
		)
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var ritems []pkg.UserSession
	for _, item := range ditems {
		var elem pkg.UserSession
		if err := elem.Consume(item); err != nil {
			return nil, err
		}
		ritems = append(ritems, elem)
	}

	return ritems, nil

}

// GetByField retrieves a record from the db using the provided field key and value
// returns the pkg.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) GetByField(ctx context.Context, key string, value interface{}) (pkg.UserSession, error) {
	m := metrics.NewTrace("UserSessionDB.GetByField")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.GetByField"), metrics.With(key, value), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return pkg.UserSession{}, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return pkg.UserSession{}, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return pkg.UserSession{}, err
	}

	defer session.Close()

	query := bson.M{key: value}

	var item map[string]interface{}

	if err := database.C(mdb.col).Find(query).One(&item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of UserSession type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return pkg.UserSession{}, ErrNotFound
		}
		return pkg.UserSession{}, err
	}

	var elem pkg.UserSession

	if err := elem.Consume(item); err != nil {
		return pkg.UserSession{}, err
	}

	return elem, nil

}

// Get retrieves a record from the db using the publicID and returns the pkg.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) Get(ctx context.Context, publicID string) (pkg.UserSession, error) {
	m := metrics.NewTrace("UserSessionDB.Get")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Get"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return pkg.UserSession{}, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return pkg.UserSession{}, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return pkg.UserSession{}, err
	}

	defer session.Close()

	query := bson.M{"public_id": publicID}

	var item map[string]interface{}

	if err := database.C(mdb.col).Find(query).One(&item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of UserSession type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return pkg.UserSession{}, ErrNotFound
		}
		return pkg.UserSession{}, err
	}

	var elem pkg.UserSession

	if err := elem.Consume(item); err != nil {
		return pkg.UserSession{}, err
	}

	return elem, nil

}

// Update uses a record from the db using the publicID and returns the pkg.UserSession type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given UserSession struct.
func (mdb *UserSessionDB) Update(ctx context.Context, publicID string, elem pkg.UserSession) error {
	m := metrics.NewTrace("UserSessionDB.Update")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Update"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to finish, context has expired"), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	defer session.Close()

	query := bson.M{"public_id": publicID}

	fields, err := elem.Fields()
	if err != nil {
		mdb.metrics.Emit(
			metrics.Errorf("Failed to get Fields() for UserSession record"),
			metrics.With("collection", mdb.col),
			metrics.With("elem", elem),
			metrics.With("error", err.Error()),
		)
		return err
	}

	if err := database.C(mdb.col).Update(query, fields); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update UserSession record"), metrics.With("query", query), metrics.With("public_id", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(
		metrics.Info("Create record"),
		metrics.With("collection", mdb.col),
		metrics.With("query", query),
		metrics.With("data", fields),
		metrics.With("public_id", publicID),
	)

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("query", query))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the collection.
func (mdb *UserSessionDB) Exec(ctx context.Context, fx func(col *mgo.Collection) error) error {
	m := metrics.NewTrace("UserSessionDB.Exec")
	defer mdb.metrics.Emit(metrics.Info("UserSessionDB.Exec"), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	defer session.Close()

	if err := fx(database.C(mdb.col)); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Operation executed"), metrics.With("collection", mdb.col))

	return nil
}