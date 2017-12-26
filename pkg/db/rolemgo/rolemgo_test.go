package rolemgo_test

import (
	"os"

	"time"

	"context"

	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	mdb "github.com/gokit/tenancykit/pkg/db/rolemgo"

	fixtures "github.com/gokit/tenancykit/pkg/db/rolemgo/fixtures"
)

var (
	events = metrics.New(custom.StackDisplay(os.Stdout))

	config = mdb.Config{
		Mode:     mgo.Monotonic,
		DB:       os.Getenv("PKG_MONGO_TEST_DB"),
		Host:     os.Getenv("PKG_MONGO_TEST_HOST"),
		User:     os.Getenv("PKG_MONGO_TEST_USER"),
		AuthDB:   os.Getenv("PKG_MONGO_TEST_AUTHDB"),
		Password: os.Getenv("PKG_MONGO_TEST_PASSWORD"),
	}

	testCol = "role_test_collection"
)

// TestGetRole validates the retrieval of a Role
// record from a mongodb.
func TestGetRole(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

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
// record from a mongodb.
func TestGetAllRole(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

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

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for Role from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Role from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Role from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Role from db.")
}

// TestGetAllRoleOrderBy validates the retrieval of all Role
// record from a mongodb.
func TestGetAllRoleByOrder(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

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
// record with a mongodb.
func TestRoleCreate(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

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
// record with a mongodb.
func TestRoleUpdate(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

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
// record from a mongodb.
func TestRoleDelete(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

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
