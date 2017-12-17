package usersql_test

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

	sqldb "github.com/gokit/tenancykit/db/usersql"
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

	testCol = "user_test"
)

// TestGetUser validates the retrieval of a User
// record from a sqldb.
func TestGetUser(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(userCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for User from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for User from db.")
}

// TestGetAllUser validates the retrieval of all User
// record from a sqldb.
func TestGetAllUser(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(userCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for User from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for User from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for User from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for User from db.")
}

// TestUserCreate validates the creation of a User
// record with a sqldb.
func TestUserCreate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(userCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")
}

// TestUserUpdate validates the update of a User
// record with a sqldb.
func TestUserUpdate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(userCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	//TODO: Update something.

	if err := api.Update(ctx, elem.PublicID, elem); err != nil {
		tests.Failed("Successfully updated record for User into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for User into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for User from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for User from db.")
}

// TestUserDelete validates the removal of a User
// record from a sqldb.
func TestUserDelete(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(userCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for User into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for User into db.")
}
