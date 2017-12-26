package activitymgo_test

import (
	"os"

	"time"

	"context"

	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	mdb "github.com/gokit/tenancykit/pkg/db/activitymgo"

	fixtures "github.com/gokit/tenancykit/pkg/db/activitymgo/fixtures"
)

var (
	config = mdb.Config{
		Mode:     mgo.Monotonic,
		DB:       os.Getenv("PKG_MONGO_TEST_DB"),
		Host:     os.Getenv("PKG_MONGO_TEST_HOST"),
		User:     os.Getenv("PKG_MONGO_TEST_USER"),
		AuthDB:   os.Getenv("PKG_MONGO_TEST_AUTHDB"),
		Password: os.Getenv("PKG_MONGO_TEST_PASSWORD"),
	}

	testCol = "activity_test_collection"
)

// TestGetActivity validates the retrieval of a Activity
// record from a mongodb.
func TestGetActivity(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadActivityJSON(fixtures.ActivityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Activity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Activity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Activity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Activity into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for Activity from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for Activity from db.")
}

// TestGetAllActivity validates the retrieval of all Activity
// record from a mongodb.
func TestGetAllActivity(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadActivityJSON(fixtures.ActivityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Activity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Activity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Activity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Activity into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for Activity from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Activity from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Activity from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Activity from db.")
}

// TestGetAllActivityOrderBy validates the retrieval of all Activity
// record from a mongodb.
func TestGetAllActivityByOrder(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadActivityJSON(fixtures.ActivityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Activity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Activity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Activity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Activity into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
	if err != nil {
		tests.Failed("Successfully retrieved all records for Activity from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Activity from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Activity from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Activity from db.")
}

// TestActivityCreate validates the creation of a Activity
// record with a mongodb.
func TestActivityCreate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadActivityJSON(fixtures.ActivityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Activity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Activity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Activity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Activity into db.")
}

// TestActivityUpdate validates the update of a Activity
// record with a mongodb.
func TestActivityUpdate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadActivityJSON(fixtures.ActivityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Activity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Activity record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Activity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Activity into db.")

	elem2, err := fixtures.LoadActivityJSON(fixtures.ActivityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Activity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Activity record")

	elem2.PublicID = elem.PublicID

	if err := api.Update(ctx, elem2.PublicID, elem2); err != nil {
		tests.Failed("Successfully updated record for Activity into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for Activity into db.")
}

// TestActivityDelete validates the removal of a Activity
// record from a mongodb.
func TestActivityDelete(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadActivityJSON(fixtures.ActivityJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Activity record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Activity record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Activity into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Activity into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for Activity into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for Activity into db.")

	if _, err = api.Get(ctx, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for Activity into db.")
	}
	tests.Passed("Successfully failed to get deleted record for Activity into db.")
}
