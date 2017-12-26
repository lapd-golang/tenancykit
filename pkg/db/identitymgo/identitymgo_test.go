package identitymgo_test

import (
	"os"

	"time"

	"context"

	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	mdb "github.com/gokit/tenancykit/pkg/db/identitymgo"

	fixtures "github.com/gokit/tenancykit/pkg/db/identitymgo/fixtures"
)

var (
	events = metrics.New(custom.StackDisplay(os.Stdout))

	config = mdb.Config{
		Mode:     mgo.Monotonic,
		DB:       os.Getenv("USERCLAIMJWT_MONGO_TEST_DB"),
		Host:     os.Getenv("USERCLAIMJWT_MONGO_TEST_HOST"),
		User:     os.Getenv("USERCLAIMJWT_MONGO_TEST_USER"),
		AuthDB:   os.Getenv("USERCLAIMJWT_MONGO_TEST_AUTHDB"),
		Password: os.Getenv("USERCLAIMJWT_MONGO_TEST_PASSWORD"),
	}

	testCol = "identity_test_collection"
)

// TestGetIdentity validates the retrieval of a Identity
// record from a mongodb.
func TestGetIdentity(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadIdentityJSON(fixtures.IdentityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Identity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Identity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Identity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Identity into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for Identity from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for Identity from db.")
}

// TestGetAllIdentity validates the retrieval of all Identity
// record from a mongodb.
func TestGetAllIdentity(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadIdentityJSON(fixtures.IdentityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Identity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Identity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Identity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Identity into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for Identity from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Identity from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Identity from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Identity from db.")
}

// TestGetAllIdentityOrderBy validates the retrieval of all Identity
// record from a mongodb.
func TestGetAllIdentityByOrder(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadIdentityJSON(fixtures.IdentityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Identity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Identity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Identity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Identity into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
	if err != nil {
		tests.Failed("Successfully retrieved all records for Identity from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Identity from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Identity from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Identity from db.")
}

// TestIdentityCreate validates the creation of a Identity
// record with a mongodb.
func TestIdentityCreate(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadIdentityJSON(fixtures.IdentityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Identity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Identity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Identity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Identity into db.")
}

// TestIdentityUpdate validates the update of a Identity
// record with a mongodb.
func TestIdentityUpdate(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadIdentityJSON(fixtures.IdentityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Identity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Identity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Identity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Identity into db.")

	elem2, err := fixtures.LoadIdentityJSON(fixtures.IdentityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Identity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Identity record")

	elem2.PublicID = elem.PublicID

	if err := api.Update(ctx, elem2.PublicID, elem2); err != nil {
		tests.Failed("Successfully updated record for Identity into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for Identity into db.")
}

// TestIdentityDelete validates the removal of a Identity
// record from a mongodb.
func TestIdentityDelete(t *testing.T) {
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadIdentityJSON(fixtures.IdentityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Identity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Identity record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Identity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Identity into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for Identity into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for Identity into db.")

	if _, err = api.Get(ctx, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for Identity into db.")
	}
	tests.Passed("Successfully failed to get deleted record for Identity into db.")
}
