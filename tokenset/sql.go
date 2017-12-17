package tokenset

import (
	"fmt"

	"github.com/gokit/tenancykit/tokenset/tokens"
	"github.com/gokit/tenancykit/tokenset/tokens/tokensql"
	"github.com/influx6/faux/context"
	"github.com/influx6/faux/db/sql"
	"github.com/influx6/faux/db/sql/tables"
	"github.com/influx6/faux/metrics"
)

// SetSQLDB implements the tokenset.Set interface on top of
// mongodb.
type SetSQLDB struct {
	table string
	db    *tokensql.TokenDB
	m     metrics.Metrics
}

// MakeSetSQL returns a new instance of SetMgoDB which has it's index set and
// giving collection named {{domain}}_tokenset.
func MakeSetSQL(domain string, m metrics.Metrics, db sql.DB) *SetSQLDB {
	domainName := fmt.Sprintf("%s_token_set", domain)
	return NewSetSQLDB(domainName, m, tokensql.New(domainName, m, db, tables.TableMigration{
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

// NewSetSQLDB returns a new instance of the SetSQLDB.
func NewSetSQLDB(domain string, m metrics.Metrics, db *tokensql.TokenDB) *SetSQLDB {
	return &SetSQLDB{
		m:     m,
		db:    db,
		table: domain,
	}
}

// HasToken returns true/false if giving token exists or not within underline
// db. Returns an error if call to db failed.
func (tf *SetSQLDB) HasToken(ctx context.Context, targetID string, token string) (bool, error) {
	var status bool
	return status, tf.db.Exec(ctx, func(qx *sql.SQL, dx sql.DB) error {
		db, err := dx.New()
		if err != nil {
			return err
		}

		query := fmt.Sprintf("SELECT * FROM %s WHERE value=%+q target_id=%+q", tf.table, targetID, token)
		row := db.QueryRowx(query)
		if err := row.Err(); err != nil {
			return err
		}

		res, err := row.SliceScan()
		if err != nil {
			return err
		}

		if len(res) != 0 {
			status = true
		}
		return nil
	})
}

// AddToken adds giving underline tokendb into db, returning error if
// it fails to do so, or call to db errors out.
func (tf *SetSQLDB) AddToken(ctx context.Context, targetID string, token string) (tokens.Token, error) {
	tf.m.Emit(
		metrics.Info("Add token"),
		metrics.With("target_id", targetID),
		metrics.With("value", token),
	)
	tokened := tokens.New(targetID, token)
	return tokened, tf.db.Create(ctx, tokened)
}
