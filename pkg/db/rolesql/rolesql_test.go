package rolesql_test

import (
	"os"

	"time"

	"testing"

	"context"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/db/sql"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lib/pq"

	_ "github.com/mattn/go-sqlite3"

	sqldb "github.com/gokit/tenancykit/pkg/db/rolesql"

	fixtures "github.com/gokit/tenancykit/pkg/db/rolesql/fixtures"
)

var (
	config = sql.Config{
		DBName:       os.Getenv("PKG_SQL_TEST_DB"),
		User:         os.Getenv("PKG_SQL_TEST_USER"),
		DBIP:         os.Getenv("PKG_SQL_TEST_ADDR"),
		DBPort:       os.Getenv("PKG_SQL_TEST_PORT"),
		DBDriver:     os.Getenv("PKG_SQL_TEST_Driver"),
		UserPassword: os.Getenv("PKG_SQL_TEST_PASSWORD"),
	}

	testCol = "role_test"
)

// TestGetRole validates the retrieval of a Role
// record from a sqldb.
func TestGetRole(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadRoleJSON(fixtures.RoleJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Role record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Role record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Role into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Role into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for Role from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for Role from db.")
}

// TestGetAllRole validates the retrieval of all Role
// record from a sqldb.
func TestGetAllRole(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadRoleJSON(fixtures.RoleJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Role record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Role record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Role into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Role into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
	if err != nil {
		tests.Failed("Successfully retrieved all records for Role from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Role from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Role from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Role from db.")
}

// TestRoleCreate validates the creation of a Role
// record with a sqldb.
func TestRoleCreate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadRoleJSON(fixtures.RoleJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Role record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Role record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Role into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Role into db.")
}

// TestRoleUpdate validates the update of a Role
// record with a sqldb.
func TestRoleUpdate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadRoleJSON(fixtures.RoleJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Role record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Role record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Role into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Role into db.")

	elem2, err := fixtures.LoadRoleJSON(fixtures.RoleJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Role record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Role record")

	elem2.PublicID = elem.PublicID

	if err := api.Update(ctx, elem2.PublicID, elem2); err != nil {
		tests.Failed("Successfully updated record for Role into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for Role into db.")
}

// TestRoleDelete validates the removal of a Role
// record from a sqldb.
func TestRoleDelete(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadRoleJSON(fixtures.RoleJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Role record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Role record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Role into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Role into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for Role into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for Role into db.")

	if _, err = api.Get(ctx, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for Role into db.")
	}
	tests.Passed("Successfully failed to get deleted record for Role into db.")
}
