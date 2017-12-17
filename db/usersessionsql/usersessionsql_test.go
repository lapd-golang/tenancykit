package usersessionsql_test

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

	sqldb "github.com/gokit/tenancykit/db/usersessionsql"
)

var (
	events = metrics.New(custom.StackDisplay(os.Stdout))

	config = sql.Config{
		DBName:       os.Getenv("tenancykit_SQL_TEST_DB"),
		User:         os.Getenv("tenancykit_SQL_TEST_USER"),
		DBIP:         os.Getenv("tenancykit_SQL_TEST_ADDR"),
		DBPort:       os.Getenv("tenancykit_SQL_TEST_PORT"),
		DBDriver:     os.Getenv("tenancykit_SQL_TEST_Driver"),
		UserPassword: os.Getenv("tenancykit_SQL_TEST_PASSWORD"),
	}

	testCol = "usersession_test"
)

// TestGetUserSession validates the retrieval of a UserSession
// record from a sqldb.
func TestGetUserSession(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(usersessionCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for UserSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for UserSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for UserSession into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for UserSession from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for UserSession from db.")
}

// TestGetAllUserSession validates the retrieval of all UserSession
// record from a sqldb.
func TestGetAllUserSession(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(usersessionCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for UserSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for UserSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for UserSession into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for UserSession from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for UserSession from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for UserSession from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for UserSession from db.")
}

// TestUserSessionCreate validates the creation of a UserSession
// record with a sqldb.
func TestUserSessionCreate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(usersessionCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for UserSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for UserSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for UserSession into db.")
}

// TestUserSessionUpdate validates the update of a UserSession
// record with a sqldb.
func TestUserSessionUpdate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(usersessionCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for UserSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for UserSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for UserSession into db.")

	//TODO: Update something.

	if err := api.Update(ctx, elem.PublicID, elem); err != nil {
		tests.Failed("Successfully updated record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for UserSession into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for UserSession from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for UserSession from db.")
}

// TestUserSessionDelete validates the removal of a UserSession
// record from a sqldb.
func TestUserSessionDelete(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(usersessionCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for UserSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for UserSession record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for UserSession into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for UserSession into db.")
}
