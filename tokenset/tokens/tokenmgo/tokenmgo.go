// Package tokenmgo provides a auto-generated package which contains a sql CRUD API for the specific Token struct in package tokens.
//
//
package tokenmgo

import (
	"fmt"

	"strings"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/metrics"

	"github.com/gokit/tenancykit/tokenset/tokens"
)

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

// Mongod defines a interface which exposes a method for retrieving a
// mongo.Database and mongo.Session.
type Mongod interface {
	New() (*mgo.Database, *mgo.Session, error)
}

// TokenDB defines a structure which provide DB CRUD operations
// using mongo as the underline db.
type TokenDB struct {
	col             string
	db              Mongod
	metrics         metrics.Metrics
	ensuredIndex    bool
	incompleteIndex bool
	indexes         []mgo.Index
}

// New returns a new instance of TokenDB.
func New(col string, m metrics.Metrics, mo Mongod, indexes ...mgo.Index) *TokenDB {
	return &TokenDB{
		db:      mo,
		col:     col,
		metrics: m,
		indexes: indexes,
	}
}

// ensureIndex attempts to ensure all provided indexes into the specific collection.
func (mdb *TokenDB) ensureIndex() error {
	if mdb.ensuredIndex {
		return nil
	}

	m := metrics.NewTrace("TokenDB.ensureIndex")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.ensureIndex"), metrics.WithTrace(m.End()))

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
func (mdb *TokenDB) Count(ctx context.Context) (int, error) {
	m := metrics.NewTrace("TokenDB.Count")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Count"), metrics.WithTrace(m.End()))

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
// on the given tokens.Token struct.
func (mdb *TokenDB) Delete(ctx context.Context, publicID string) error {
	m := metrics.NewTrace("TokenDB.Delete")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Delete"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

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
		return err
	}

	mdb.metrics.Emit(metrics.Info("Deleted record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("publicID", publicID))

	return nil
}

// Create attempts to add the record into the db using the provided instance of the
// tokens.Token.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) Create(ctx context.Context, elem tokens.Token) error {
	m := metrics.NewTrace("TokenDB.Create")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Create"), metrics.With("publicID", elem.PublicID), metrics.WithTrace(m.End()))

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
			metrics.Errorf("Failed to get Fields() for Token record"),
			metrics.With("collection", mdb.col),
			metrics.With("elem", elem),
			metrics.With("error", err.Error()),
		)
		return err
	}

	if err := database.C(mdb.col).Insert(bson.M(fields)); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create Token record"), metrics.With("collection", mdb.col), metrics.With("elem", elem), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Create record"), metrics.With("collection", mdb.col), metrics.With("elem", elem))

	return nil
}

// GetAllPerPage retrieves all records from the db and returns a slice of tokens.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) GetAllPerPage(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]tokens.Token, int, error) {
	m := metrics.NewTrace("TokenDB.GetAll")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.GetAll"), metrics.WithTrace(m.End()))

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

	var ritems []tokens.Token

	var ditems []map[string]interface{}
	if err := database.C(mdb.col).Find(query).Skip(indexToStart).Limit(totalWanted).Sort(orderBy).All(&ditems); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Token type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		return nil, -1, err
	}

	for _, item := range ditems {
		var elem tokens.Token
		if err := elem.Consume(item); err != nil {
			return nil, -1, err
		}
		ritems = append(ritems, elem)
	}

	return ritems, totalRecords, nil
}

// GetAllByOrder retrieves all records from the db and returns a slice of tokens.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) GetAllByOrder(ctx context.Context, order, orderBy string) ([]tokens.Token, error) {
	m := metrics.NewTrace("TokenDB.GetAllByOrder")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.GetAllByOrder"), metrics.WithTrace(m.End()))

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
			metrics.Errorf("Failed to retrieve all records of Token type from db"),
			metrics.With("collection", mdb.col),
			metrics.With("query", query),
			metrics.With("error", err.Error()),
		)
		return nil, err
	}

	var ritems []tokens.Token
	for _, item := range ditems {
		var elem tokens.Token
		if err := elem.Consume(item); err != nil {
			return nil, err
		}
		ritems = append(ritems, elem)
	}

	return ritems, nil

}

// GetByField retrieves a record from the db using the provided field key and value
// returns the tokens.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) GetByField(ctx context.Context, key string, value interface{}) (tokens.Token, error) {
	m := metrics.NewTrace("TokenDB.GetByField")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.GetByField"), metrics.With(key, value), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return tokens.Token{}, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return tokens.Token{}, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With(key, value), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))

		return tokens.Token{}, err
	}

	defer session.Close()

	query := bson.M{key: value}

	var item map[string]interface{}

	if err := database.C(mdb.col).Find(query).One(&item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Token type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))

		return tokens.Token{}, err
	}

	var elem tokens.Token

	if err := elem.Consume(item); err != nil {
		return tokens.Token{}, err
	}

	return elem, nil

}

// Get retrieves a record from the db using the publicID and returns the tokens.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) Get(ctx context.Context, publicID string) (tokens.Token, error) {
	m := metrics.NewTrace("TokenDB.Get")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Get"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

	if context.IsExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return tokens.Token{}, err
	}

	if err := mdb.ensureIndex(); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to apply index"), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return tokens.Token{}, err
	}

	database, session, err := mdb.db.New()
	if err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create session"), metrics.With("publicID", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
		return tokens.Token{}, err
	}

	defer session.Close()

	query := bson.M{"public_id": publicID}

	var item map[string]interface{}

	if err := database.C(mdb.col).Find(query).One(&item); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve all records of Token type from db"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("error", err.Error()))
		return tokens.Token{}, err
	}

	var elem tokens.Token

	if err := elem.Consume(item); err != nil {
		return tokens.Token{}, err
	}

	return elem, nil

}

// Update uses a record from the db using the publicID and returns the tokens.Token type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given Token struct.
func (mdb *TokenDB) Update(ctx context.Context, publicID string, elem tokens.Token) error {
	m := metrics.NewTrace("TokenDB.Update")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Update"), metrics.With("publicID", publicID), metrics.WithTrace(m.End()))

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

	if fielder, ok := interface{}(elem).(TokenFields); ok {
		fields, err := fielder.Fields()
		if err != nil {
			mdb.metrics.Emit(
				metrics.Errorf("Failed to get Fields() for Token record"),
				metrics.With("collection", mdb.col),
				metrics.With("elem", elem),
				metrics.With("error", err.Error()),
			)
			return err
		}

		if err := database.C(mdb.col).Update(query, fields); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to update Token record"), metrics.With("query", query), metrics.With("public_id", publicID), metrics.With("collection", mdb.col), metrics.With("error", err.Error()))
			return err
		}

		mdb.metrics.Emit(
			metrics.Info("Create record"),
			metrics.With("collection", mdb.col),
			metrics.With("query", query),
			metrics.With("data", fields),
			metrics.With("public_id", publicID),
		)

		return nil
	}

	queryData := bson.M(map[string]interface{}{

		"public_id": elem.PublicID,

		"target_id": elem.TargetID,

		"value": elem.Value,
	})

	if err := database.C(mdb.col).Update(query, queryData); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update Token record"), metrics.With("collection", mdb.col), metrics.With("query", query), metrics.With("data", queryData), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("collection", mdb.col), metrics.With("public_id", publicID), metrics.With("query", query))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the collection.
func (mdb *TokenDB) Exec(ctx context.Context, fx func(col *mgo.Collection) error) error {
	m := metrics.NewTrace("TokenDB.Exec")
	defer mdb.metrics.Emit(metrics.Info("TokenDB.Exec"), metrics.WithTrace(m.End()))

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
		return err
	}

	mdb.metrics.Emit(metrics.Info("Operation executed"), metrics.With("collection", mdb.col))

	return nil
}
