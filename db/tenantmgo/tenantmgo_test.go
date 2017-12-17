package tenantmgo_test

import (
	"os"

	"time"

	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/db/mongo"

	"github.com/influx6/faux/metrics/custom"

	mdb "github.com/gokit/tenancykit/db/tenantmgo"
)

var (
	events = metrics.New(custom.StackDisplay(os.Stdout))

	config = mongo.Config{
		Mode:     mgo.Monotonic,
		DB:       os.Getenv("tenancykit_MONGO_TEST_DB"),
		Host:     os.Getenv("tenancykit_MONGO_TEST_HOST"),
		User:     os.Getenv("tenancykit_MONGO_TEST_USER"),
		AuthDB:   os.Getenv("tenancykit_MONGO_TEST_AUTHDB"),
		Password: os.Getenv("tenancykit_MONGO_TEST_PASSWORD"),
	}

	testCol = "tenant_test_collection"
)

// TestGetTenant validates the retrieval of a Tenant
// record from a mongodb.
func TestGetTenant(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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
// record from a mongodb.
func TestGetAllTenant(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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

// TestGetAllTenantOrderBy validates the retrieval of all Tenant
// record from a mongodb.
func TestGetAllTenantByOrder(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
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
// record with a mongodb.
func TestTenantCreate(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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
// record with a mongodb.
func TestTenantUpdate(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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
}

// TestTenantDelete validates the removal of a Tenant
// record from a mongodb.
func TestTenantDelete(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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
