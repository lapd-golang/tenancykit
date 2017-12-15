// Package mdb provides a auto-generated package which contains a mongo CRUD API for the specific User struct in package users.
//
//
package userdb

import (
	"fmt"
	"strings"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/metrics"

	"github.com/gokit/tenancykit/users"
)

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

// UserBSON defines an interface which exposes method to return a bson.M type
// which contains all related fields for the giving  object.
type UserBSON interface {
	BSON() bson.M
}

// UserBSONConsumer defines an interface which accepts a map of data which will be consumed
// into the giving implementing structure as decided by the structure.
type UserBSONConsumer interface {
	BSONConsume(bson.M) error
}

// Mongod defines a interface which exposes a method for retrieving a
// mongo.Database and mongo.Session.
type Mongod interface {
	New() (*mgo.Database, *mgo.Session, error)
}

// UserDB defines a structure which provide DB CRUD operations
// using mongo as the underline db.
type UserDB struct {
	col             string
	db              Mongod
	metrics         metrics.Metrics
	ensuredIndex    bool
	incompleteIndex bool
	indexes         []mgo.Index
}

// New returns a new instance of UserDB.
func New(col string, m metrics.Metrics, mo Mongod, indexes ...mgo.Index) *UserDB {
	return &UserDB{
		db:      mo,
		col:     col,
		metrics: m,
		indexes: indexes,
	}
}

// ensureIndex attempts to ensure all provided indexes into the specific collection.
func (mdb *UserDB) ensureIndex() error {
	if mdb.ensuredIndex {
		return nil
	}

	m := metrics.NewTrace("UserDB.ensureIndex")
	defer mdb.metrics.Emit(metrics.Info("UserDB.ensureIndex"), metrics.WithTrace(m.End()))

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
func (mdb *UserDB) Count(ctx context.Context) (int, error) {
	m := metrics.NewTrace("UserDB.Count")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Count"), metrics.WithTrace(m.End()))

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

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to get record count"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return -1, err
	}

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
// on the given users.User struct.
func (mdb *UserDB) Delete(ctx context.Context, publicID string) error {
	m := metrics.NewTrace("UserDB.Delete")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Delete"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

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

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := database.C(mdb.col).Remove(query); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("publicID", publicID), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Deleted record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("publicID", publicID))

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

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.With("publicID", elem.PublicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if fields, ok := interface{}(elem).(UserBSON); ok {
		if err := database.C(mdb.col).Insert(fields.BSON()); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to create User record"), metrics.With("collection", mdb.col), metrics.With("elem", elem), metrics.With("error", err.Error()))
			return err
		}

		mdb.metrics.Emit(metrics.Info("Create record"), metrics.With("collection", mdb.col), metrics.With("elem", elem))

		return nil
	}

	if fields, ok := interface{}(elem).(UserFields); ok {
		if err := database.C(mdb.col).Insert(bson.M(fields.Fields())); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to create User record"), metrics.With("collection", mdb.col), metrics.With("elem", elem), metrics.With("error", err.Error()))
			return err
		}

		mdb.metrics.Emit(metrics.Info("Create record"), metrics.With("collection", mdb.col), metrics.With("elem", elem))

		return nil
	}

	query := bson.M(map[string]interface{}{

		"created_at": elem.Created,

		"hash": elem.Hash,

		"private_id": elem.PrivateID,

		"public_id": elem.PublicID,

		"two_factor_auth": elem.TwoFactorAuth,

		"updated_at": elem.Updated,

		"username": elem.Username,
	})

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.With("publicID", elem.PublicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := database.C(mdb.col).Insert(query); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create User record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Create record"), metrics.With("collection", mdb.col), metrics.With("query", query))

	return nil
}

// GetAllPerPage retrieves all records from the db and returns a slice of users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) GetAllPerPage(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]users.User, int, error) {
	m := metrics.NewTrace("UserDB.GetAll")
	defer mdb.metrics.Emit(metrics.Info("UserDB.GetAll"), metrics.WithTrace(m.End()))

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

	mdb.metrics.Emit(metrics.Info("DB:Query:GetAllPerPage").WithFields(metrics.Field{
		"starting_index":       indexToStart,
		"total_records_wanted": totalWanted,
		"order":                order,
		"orderBy":              orderBy,
		"page":                 page,
		"responsePerPage":      responsePerPage,
	}))

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, -1, err
	}

	defer session.Close()

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, -1, err
	}

	query := bson.M{}

	var items []users.User

	if err := database.C(mdb.col).Find(query).Skip(indexToStart).Limit(totalWanted).Sort(orderBy).All(&items); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of User type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))

		return nil, -1, err
	}

	return items, totalRecords, nil

}

// GetAllByOrder retrieves all records from the db and returns a slice of users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) GetAllByOrder(ctx context.Context, order, orderBy string) ([]users.User, error) {
	m := metrics.NewTrace("UserDB.GetAllByOrder")
	defer mdb.metrics.Emit(metrics.Info("UserDB.GetAllByOrder"), metrics.WithTrace(m.End()))

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

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")

		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return nil, err
	}

	query := bson.M{}

	var items []users.User

	if err := database.C(mdb.col).Find(query).Sort(orderBy).All(&items); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of User type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))

		return nil, err
	}

	return items, nil

}

// GetByField retrieves a record from the db using the provided field key and value
// returns the users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) GetByField(ctx context.Context, key string, value interface{}) (users.User, error) {
	m := metrics.NewTrace("UserDB.GetByField")
	defer mdb.metrics.Emit(metrics.Info("UserDB.GetByField"), metrics.With(key, value), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return users.User{}, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return users.User{}, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return users.User{}, err
	}

	defer session.Close()

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return users.User{}, err
	}

	query := bson.M{key: value}

	var item users.User

	if err := database.C(mdb.col).Find(query).One(&item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of User type from db"), metrics.With("query", query), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return users.User{}, err
	}

	return item, nil

}

// Get retrieves a record from the db using the publicID and returns the users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) Get(ctx context.Context, publicID string) (users.User, error) {
	m := metrics.NewTrace("UserDB.Get")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Get"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return users.User{}, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return users.User{}, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return users.User{}, err
	}

	defer session.Close()

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return users.User{}, err
	}

	query := bson.M{"public_id": publicID}

	var item users.User

	if err := database.C(mdb.col).Find(query).One(&item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of User type from db"), metrics.With("query", query), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return users.User{}, err
	}

	return item, nil

}

// Update uses a record from the db using the publicID and returns the users.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) Update(ctx context.Context, publicID string, elem users.User) error {
	m := metrics.NewTrace("UserDB.Update")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Update"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

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

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to finish, context has expired"), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	query := bson.M{"public_id": publicID}

	if fields, ok := interface{}(elem).(UserBSON); ok {
		if err := database.C(mdb.col).Update(query, fields.BSON()); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to update User record"), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("query", query), metrics.With("error", err.Error()))

			return err
		}

		mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("query", query), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("data", fields.BSON()))

		return nil
	}

	if fields, ok := interface{}(elem).(UserFields); ok {
		if err := database.C(mdb.col).Update(query, fields.Fields()); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to update User record"), metrics.With("query", query), metrics.With("public_id", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
			return err
		}

		mdb.metrics.Emit(metrics.Info("Create record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("data", fields.Fields()), metrics.With("public_id", publicID))

		return nil
	}

	queryData := bson.M(map[string]interface{}{

		"created_at": elem.Created,

		"hash": elem.Hash,

		"private_id": elem.PrivateID,

		"public_id": elem.PublicID,

		"two_factor_auth": elem.TwoFactorAuth,

		"updated_at": elem.Updated,

		"username": elem.Username,
	})

	if err := database.C(mdb.col).Update(query, queryData); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update User record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("data", queryData), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("query", query))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the collection.
func (mdb *UserDB) Exec(ctx context.Context, fx func(col *mgo.Collection) error) error {
	m := metrics.NewTrace("UserDB.Exec")
	defer mdb.metrics.Emit(metrics.Info("UserDB.Exec"), metrics.WithTrace(m.End()))

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

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to finish, context has expired"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := fx(database.C(mdb.col)); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Operation executed").
		With("collection", mdb.col))

	return nil
}
