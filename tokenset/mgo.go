package tokenset

import (
	"fmt"

	"github.com/gokit/tenancykit/tokenset/tokens"
	"github.com/gokit/tenancykit/tokenset/tokens/tokenmgo"
	"github.com/influx6/faux/context"
	"github.com/influx6/faux/metrics"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SetMgoDB implements the tokenset.Set interface on top of
// mongodb.
type SetMgoDB struct {
	db *tokenmgo.TokenDB
	m  metrics.Metrics
}

// MakeSetMgo returns a new instance of SetMgoDB which has it's index set and
// giving collection named {{domain}}_tokenset_collection.
func MakeSetMgo(domain string, m metrics.Metrics, mo tokenmgo.Mongod) *SetMgoDB {
	domainName := fmt.Sprintf("%s_tokenset_collection", domain)

	return NewSetMgoDB(m, tokenmgo.New(domainName, m, mo, mgo.Index{
		Unique: true,
		Key:    []string{"value"},
		Name:   "value_index",
	}))
}

// NewSetMgoDB returns a new instance of the SetMgoDB
func NewSetMgoDB(m metrics.Metrics, db *tokenmgo.TokenDB) *SetMgoDB {
	return &SetMgoDB{
		m:  m,
		db: db,
	}
}

// HasToken returns true/false if giving token exists or not within underline
// db. Returns an error if call to db failed.
func (tf *SetMgoDB) HasToken(ctx context.Context, targetID string, token string) (bool, error) {
	var status bool
	return status, tf.db.Exec(ctx, func(col *mgo.Collection) error {
		query := bson.M{"target_id": targetID, "value": token}
		tf.m.Emit(
			metrics.Info("Check if token exists"),
			metrics.With("target_id", targetID),
			metrics.With("value", token),
		)
		total, err := col.Find(query).Count()
		if err != nil {
			return err
		}
		if total != 0 {
			status = true
		}
		return nil
	})
}

// AddToken adds giving underline tokendb into db, returning error if
// it fails to do so, or call to db errors out.
func (tf *SetMgoDB) AddToken(ctx context.Context, targetID string, token string) (tokens.Token, error) {
	tf.m.Emit(
		metrics.Info("Add token"),
		metrics.With("target_id", targetID),
		metrics.With("value", token),
	)
	tokened := tokens.New(targetID, token)
	return tokened, tf.db.Create(ctx, tokened)
}
