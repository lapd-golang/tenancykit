package tokensql_test

import (
	"os"

	"time"

	"testing"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/db/sql"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/metrics/custom"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lib/pq"

	_ "github.com/mattn/go-sqlite3"

	sqldb "github.com/gokit/tenancykit/tokenset/tokens/tokensql"
)

var (
	events = metrics.New(custom.StackDisplay(os.Stdout))

	config = sql.Config{
		DBName:       os.Getenv("tokens_SQL_DB"),
		User:         os.Getenv("tokens_SQL_USER"),
		DBIP:         os.Getenv("tokens_SQL_ADDR"),
		DBPort:       os.Getenv("tokens_SQL_PORT"),
		DBDriver:     os.Getenv("tokens_SQL_Driver"),
		UserPassword: os.Getenv("tokens_SQL_PASSWORD"),
	}

	testCol = "token_test"
)

// TestGetToken validates the retrieval of a Token
// record from a sqldb.
func TestGetToken(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for Token from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for Token from db.")
}

// TestGetAllToken validates the retrieval of all Token
// record from a sqldb.
func TestGetAllToken(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for Token from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Token from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Token from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Token from db.")
}

// TestTokenCreate validates the creation of a Token
// record with a sqldb.
func TestTokenCreate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")
}

// TestTokenUpdate validates the update of a Token
// record with a sqldb.
func TestTokenUpdate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	//TODO: Update something.

	if err := api.Update(ctx, elem.PublicID, elem); err != nil {
		tests.Failed("Successfully updated record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for Token into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for Token from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for Token from db.")
}

// TestTokenDelete validates the removal of a Token
// record from a sqldb.
func TestTokenDelete(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for Token into db.")
}
