package tokenset

import (
	"fmt"

	"github.com/gokit/tenancykit/tokenset/tokens"
	"github.com/gokit/tenancykit/tokenset/tokens/tokenmgo"
	"github.com/gokit/tenancykit/tokenset/tokens/tokensql"
	"github.com/gokit/tenancykit/tokenset/tokens/types"
	"github.com/influx6/faux/context"
	"github.com/influx6/faux/db/sql"
	"github.com/influx6/faux/db/sql/tables"
	"github.com/influx6/faux/metrics"
	mgo "gopkg.in/mgo.v2"
)

// TokenSet implements the tokenset.TokenSet interface on top of
// mongodb.
type TokenSet struct {
	m  metrics.Metrics
	db types.TokenDBBackend
}

// MGOTokenSet returns a new instance of TokenSet using Mongodb as underline database,
// which has it's index set and giving collection named {{domain}}_tokenset_collection.
func MGOTokenSet(domain string, m metrics.Metrics, mo tokenmgo.Mongod) TokenSet {
	domainName := fmt.Sprintf("%s_tokenset_collection", domain)
	return NewTokenSet(m, tokenmgo.New(domainName, m, mo, mgo.Index{
		Unique: true,
		Key:    []string{"value"},
		Name:   "value_index",
	}))
}

// SQLTokenSet returns a new instance of TokenSet using any provided sql database,
// which has it's index set and giving collection named {{domain}}_tokenset.
func SQLTokenSet(domain string, m metrics.Metrics, db sql.DB) TokenSet {
	domainName := fmt.Sprintf("%s_token_set", domain)
	return NewTokenSet(m, tokensql.New(domainName, m, db, tables.TableMigration{
		TableName: domainName,
		Fields: []tables.FieldMigration{
			{
				FieldName:     "id",
				FieldType:     "integer",
				AutoIncrement: true,
				PrimaryKey:    true,
				NotNull:       true,
			},
			{
				NotNull:   true,
				FieldName: "value",
				FieldType: "varchar(255)",
			},
			{
				NotNull:   true,
				FieldName: "public_id",
				FieldType: "varchar(255)",
			},
			{
				NotNull:   true,
				FieldName: "target_id",
				FieldType: "varchar(255)",
			},
		},
	}))
}

// NewTokenSet returns a new instance of the TokenSet
func NewTokenSet(m metrics.Metrics, db types.TokenDBBackend) TokenSet {
	return TokenSet{
		m:  m,
		db: db,
	}
}

// Has returns true/false if giving token exists or not within underline
// db. Returns an error if call to db failed.
func (tf TokenSet) Has(ctx context.Context, targetID string, token string) (bool, error) {
	_, err := tf.db.GetByField(ctx, "target_id", token)
	if err != nil && (err == tokensql.ErrNotFound || err == tokenmgo.ErrNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// Add adds giving underline tokendb into db, returning error if
// it fails to do so, or call to db errors out.
func (tf TokenSet) Add(ctx context.Context, targetID string, token string) (tokens.Token, error) {
	tf.m.Emit(
		metrics.Info("Add token"),
		metrics.With("target_id", targetID),
		metrics.With("value", token),
	)
	tokened := tokens.New(targetID, token)
	return tokened, tf.db.Create(ctx, tokened)
}
