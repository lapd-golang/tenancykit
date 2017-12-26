// Package identitymgo provides a auto-generated package which contains a sql CRUD API for the specific Identity struct in package userclaimjwt.
//
//
package identitymgo

import (
	"errors"
	"time"

	"fmt"

	"strings"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"context"

	"github.com/influx6/faux/metrics"

	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"
)

// errors ...
var (
	ErrNotFound = errors.New("record not found")
)

// Config provides configuration for connecting to a db.
type Config struct {
	Host     string
	AuthDB   string
	DB       string
	User     string
	Password string
	Mode     mgo.Mode
}

// GetSession attempts to retrieve the giving session for the given config.
func GetSession(config Config) (*mgo.Session, error) {
	info := mgo.DialInfo{
		Addrs:    []string{config.Host},
		Timeout:  60 * time.Second,
		Database: config.AuthDB,
		Username: config.User,
		Password: config.Password,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	ses, err := mgo.DialWithInfo(&info)
	if err != nil {
		return nil, err
	}

	if config.Mode == 0 {
		config.Mode = mgo.Monotonic
	}

	ses.SetMode(config.Mode, true)

	return ses, nil
}

// NewMongoDB returns a new instance of a MongoServer.
func NewMongoDB(config Config) MongoDB {
	var mn mongoServer
	mn.Config = config
	return &mn
}

// mongoServer defines a mongo connection manager that builds
// allows usage of a giving configuration to generate new mongo
// sessions and database instances.
type mongoServer struct {
	Config
}

// New returns a new session and database from the giving configuration.
func (m *mongoServer) New() (*mgo.Database, *mgo.Session, error) {
	ses, err := GetSession(m.Config)
	if err != nil {
		return nil, nil, err
	}

	return ses.DB(m.Config.DB), ses, nil
}

// IdentityFields defines an interface which exposes method to return a map of all
// attributes associated with the defined structure as decided by the structure.
type IdentityFields interface {
	Fields() (map[string]interface{}, error)
}

// IdentityConsumer defines an interface which accepts a map of data which will be consumed
// into the giving implementing structure as decided by the structure.
type IdentityConsumer interface {
	Consume(map[string]interface{}) error
}

// MongoDB defines a interface which exposes a method for retrieving a
// mongo.Database and mongo.Session.
type MongoDB interface {
	New() (*mgo.Database, *mgo.Session, error)
}

// IdentityDB defines a structure which provide DB CRUD operations
// using mongo as the underline db.
type IdentityDB struct {
	col             string
	db              MongoDB
	metrics         metrics.Metrics
	ensuredIndex    bool
	incompleteIndex bool
	indexes         []mgo.Index
}

// New returns a new instance of IdentityDB.
func New(col string, m metrics.Metrics, mo MongoDB, indexes ...mgo.Index) *IdentityDB {
	return &IdentityDB{
		db:      mo,
		col:     col,
		metrics: m,
		indexes: indexes,
	}
}

// ensureIndex attempts to ensure all provided indexes into the specific collection.
func (mdb *IdentityDB) ensureIndex() error {
	if mdb.ensuredIndex {
		return nil
	}

	m := metrics.NewTrace("IdentityDB.ensureIndex")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.ensureIndex"), metrics.WithTrace(m.End()))

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
func (mdb *IdentityDB) Count(ctx context.Context) (int, error) {
	m := metrics.NewTrace("IdentityDB.Count")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.Count"), metrics.WithTrace(m.End()))

	if isContextExpired(ctx) {
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
// on the given userclaimjwt.Identity struct.
func (mdb *IdentityDB) Delete(ctx context.Context, publicID string) error {
	m := metrics.NewTrace("IdentityDB.Delete")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.Delete"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if isContextExpired(ctx) {
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
// userclaimjwt.Identity.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Identity struct.
func (mdb *IdentityDB) Create(ctx context.Context, elem userclaimjwt.Identity) error {
	m := metrics.NewTrace("IdentityDB.Create")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.Create"), metrics.With("publicID", elem.PublicID), metrics.WithTrace(m.End()))

	if isContextExpired(ctx) {
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

	query := bson.M(map[string]interface{}{

		"data": elem.Data,

		"expires": elem.Expires,

		"last_login": elem.IssuedAt,

		"public_id": elem.PublicID,

		"refresh_interval": elem.RefreshInterval,

		"refresh_token": elem.RefreshToken,

		"target_id": elem.TargetID,
	})

	if err := database.C(mdb.col).Insert(query); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create Identity record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Create record"), metrics.With("collection", mdb.col), metrics.With("query", query))

	return nil
}

// GetAll retrieves all records from the db and returns a slice of userclaimjwt.Identity type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Identity struct.
func (mdb *IdentityDB) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]userclaimjwt.Identity, int, error) {
	m := metrics.NewTrace("IdentityDB.GetAll")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.GetAll"), metrics.WithTrace(m.End()))

	switch strings.ToLower(order) {
	case "dsc", "desc":
		orderBy = "-" + orderBy
	}

	if isContextExpired(ctx) {
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

	var ritems []userclaimjwt.Identity

	if err := database.C(mdb.col).Find(query).Skip(indexToStart).Limit(totalWanted).Sort(orderBy).All(&ritems); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Identity type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return nil, -1, ErrNotFound
		}
		return nil, -1, err
	}

	return ritems, totalRecords, nil
}

// GetAllByOrder retrieves all records from the db and returns a slice of userclaimjwt.Identity type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Identity struct.
func (mdb *IdentityDB) GetAllByOrder(ctx context.Context, order, orderBy string) ([]userclaimjwt.Identity, error) {
	m := metrics.NewTrace("IdentityDB.GetAllByOrder")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.GetAllByOrder"), metrics.WithTrace(m.End()))

	switch strings.ToLower(order) {
	case "dsc", "desc":
		orderBy = "-" + orderBy
	}

	if isContextExpired(ctx) {
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

	var items []userclaimjwt.Identity
	if err := database.C(mdb.col).Find(query).Sort(orderBy).All(&items); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Identity type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return items, nil

}

// GetByField retrieves a record from the db using the provided field key and value
// returns the userclaimjwt.Identity type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Identity struct.
func (mdb *IdentityDB) GetByField(ctx context.Context, key string, value interface{}) (userclaimjwt.Identity, error) {
	m := metrics.NewTrace("IdentityDB.GetByField")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.GetByField"), metrics.With(key, value), metrics.WithTrace(m.End()))

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return userclaimjwt.Identity{}, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return userclaimjwt.Identity{}, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return userclaimjwt.Identity{}, err
	}

	defer session.Close()

	query := bson.M{key: value}

	var item userclaimjwt.Identity

	if err := database.C(mdb.col).Find(query).One(&item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Identity type from db"), metrics.With("query", query), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return userclaimjwt.Identity{}, ErrNotFound
		}
		return userclaimjwt.Identity{}, err
	}

	return item, nil

}

// Get retrieves a record from the db using the publicID and returns the userclaimjwt.Identity type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Identity struct.
func (mdb *IdentityDB) Get(ctx context.Context, publicID string) (userclaimjwt.Identity, error) {
	m := metrics.NewTrace("IdentityDB.Get")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.Get"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return userclaimjwt.Identity{}, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return userclaimjwt.Identity{}, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return userclaimjwt.Identity{}, err
	}

	defer session.Close()

	query := bson.M{"public_id": publicID}

	var item userclaimjwt.Identity

	if err := database.C(mdb.col).Find(query).One(&item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Identity type from db"), metrics.With("query", query), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return userclaimjwt.Identity{}, ErrNotFound
		}
		return userclaimjwt.Identity{}, err
	}

	return item, nil

}

// Update uses a record from the db using the publicID and returns the userclaimjwt.Identity type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Identity struct.
func (mdb *IdentityDB) Update(ctx context.Context, publicID string, elem userclaimjwt.Identity) error {
	m := metrics.NewTrace("IdentityDB.Update")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.Update"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if isContextExpired(ctx) {
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

	queryData := bson.M(map[string]interface{}{

		"data": elem.Data,

		"expires": elem.Expires,

		"last_login": elem.IssuedAt,

		"public_id": elem.PublicID,

		"refresh_interval": elem.RefreshInterval,

		"refresh_token": elem.RefreshToken,

		"target_id": elem.TargetID,
	})
	if err := database.C(mdb.col).Update(query, queryData); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update Identity record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("data", queryData), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("query", query))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the collection.
func (mdb *IdentityDB) Exec(ctx context.Context, fx func(col *mgo.Collection) error) error {
	m := metrics.NewTrace("IdentityDB.Exec")
	defer mdb.metrics.Emit(metrics.Info("IdentityDB.Exec"), metrics.WithTrace(m.End()))

	if isContextExpired(ctx) {
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

func isContextExpired(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}