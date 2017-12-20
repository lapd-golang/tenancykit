package tokenset

import (
	"fmt"

	"github.com/gokit/tenancykit/pkg/db/tokenmgo"
	"github.com/gokit/tenancykit/pkg/db/tokensql"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/influx6/faux/db/sql"
	"github.com/influx6/faux/db/sql/tables"
	"github.com/influx6/faux/metrics"
	mgo "gopkg.in/mgo.v2"
)

// MGOTokenSet returns a new instance of TokenSet using Mongodb as underline database,
// which has it's index set and giving collection named {{domain}}_tokenset_collection.
func MGOTokenSet(domain string, m metrics.Metrics, mo tokenmgo.Mongod) types.TokenDBBackend {
	domainName := fmt.Sprintf("%s_tokenset_collection", domain)
	return tokenmgo.New(domainName, m, mo, mgo.Index{
		Unique: true,
		Key:    []string{"value"},
		Name:   "value_index",
	})
}

// SQLTokenSet returns a new instance of TokenSet using any provided sql database,
// which has it's index set and giving collection named {{domain}}_tokenset.
func SQLTokenSet(domain string, m metrics.Metrics, db sql.DB) types.TokenDBBackend {
	domainName := fmt.Sprintf("%s_token_set", domain)
	return tokensql.New(domainName, m, db, tables.TableMigration{
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
	})
}
