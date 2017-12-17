package tenantsql_test

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

	sqldb "github.com/gokit/tenancykit/tenantsql"
)

var (
	events = metrics.New(custom.StackDisplay(os.Stdout))

	config = sql.Config{
		DBName:       os.Getenv("tenancykit_SQL_DB"),
		User:         os.Getenv("tenancykit_SQL_USER"),
		DBIP:         os.Getenv("tenancykit_SQL_ADDR"),
		DBPort:       os.Getenv("tenancykit_SQL_PORT"),
		DBDriver:     os.Getenv("tenancykit_SQL_Driver"),
		UserPassword: os.Getenv("tenancykit_SQL_PASSWORD"),
	}

	testCol = "tenant_test"
)

// TestGetTenant validates the retrieval of a Tenant
// record from a sqldb.
func TestGetTenant(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tenantCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Tenant record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Tenant record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Tenant into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Tenant into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for Tenant from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for Tenant from db.")
}

// TestGetAllTenant validates the retrieval of all Tenant
// record from a sqldb.
func TestGetAllTenant(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tenantCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Tenant record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Tenant record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Tenant into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Tenant into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for Tenant from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Tenant from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Tenant from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Tenant from db.")
}

// TestTenantCreate validates the creation of a Tenant
// record with a sqldb.
func TestTenantCreate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tenantCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Tenant record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Tenant record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Tenant into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Tenant into db.")
}

// TestTenantUpdate validates the update of a Tenant
// record with a sqldb.
func TestTenantUpdate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tenantCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Tenant record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Tenant record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Tenant into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Tenant into db.")

	//TODO: Update something.

	if err := api.Update(ctx, elem.PublicID, elem); err != nil {
		tests.Failed("Successfully updated record for Tenant into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for Tenant into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for Tenant from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for Tenant from db.")
}

// TestTenantDelete validates the removal of a Tenant
// record from a sqldb.
func TestTenantDelete(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tenantCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Tenant record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Tenant record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Tenant into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Tenant into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for Tenant into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for Tenant into db.")
}
