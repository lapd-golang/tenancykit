package twofactorsessionmgo_test

import (
	"os"

	"time"

	"context"

	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	mdb "github.com/gokit/tenancykit/pkg/db/twofactorsessionmgo"

	fixtures "github.com/gokit/tenancykit/pkg/db/twofactorsessionmgo/fixtures"
)

var (
	config = mdb.Config{
		Mode:     mgo.Monotonic,
		DB:       os.Getenv("TENANCYKIT_MONGO_TEST_DB"),
		Host:     os.Getenv("TENANCYKIT_MONGO_TEST_HOST"),
		User:     os.Getenv("TENANCYKIT_MONGO_TEST_USER"),
		AuthDB:   os.Getenv("TENANCYKIT_MONGO_TEST_AUTHDB"),
		Password: os.Getenv("TENANCYKIT_MONGO_TEST_PASSWORD"),
	}

	testCol = "twofactorsession_test_collection"
)

// TestGetTwoFactorSession validates the retrieval of a TwoFactorSession
// record from a mongodb.
func TestGetTwoFactorSession(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTwoFactorSessionJSON(fixtures.TwoFactorSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TwoFactorSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TwoFactorSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TwoFactorSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TwoFactorSession into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for TwoFactorSession from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for TwoFactorSession from db.")
}

// TestGetAllTwoFactorSession validates the retrieval of all TwoFactorSession
// record from a mongodb.
func TestGetAllTwoFactorSession(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTwoFactorSessionJSON(fixtures.TwoFactorSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TwoFactorSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TwoFactorSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TwoFactorSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TwoFactorSession into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for TwoFactorSession from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for TwoFactorSession from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for TwoFactorSession from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for TwoFactorSession from db.")
}

// TestGetAllTwoFactorSessionOrderBy validates the retrieval of all TwoFactorSession
// record from a mongodb.
func TestGetAllTwoFactorSessionByOrder(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTwoFactorSessionJSON(fixtures.TwoFactorSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TwoFactorSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TwoFactorSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TwoFactorSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TwoFactorSession into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
	if err != nil {
		tests.Failed("Successfully retrieved all records for TwoFactorSession from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for TwoFactorSession from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for TwoFactorSession from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for TwoFactorSession from db.")
}

// TestTwoFactorSessionCreate validates the creation of a TwoFactorSession
// record with a mongodb.
func TestTwoFactorSessionCreate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTwoFactorSessionJSON(fixtures.TwoFactorSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TwoFactorSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TwoFactorSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TwoFactorSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TwoFactorSession into db.")
}

// TestTwoFactorSessionUpdate validates the update of a TwoFactorSession
// record with a mongodb.
func TestTwoFactorSessionUpdate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTwoFactorSessionJSON(fixtures.TwoFactorSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TwoFactorSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TwoFactorSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TwoFactorSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TwoFactorSession into db.")

	elem2, err := fixtures.LoadTwoFactorSessionJSON(fixtures.TwoFactorSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TwoFactorSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TwoFactorSession record")

	elem2.PublicID = elem.PublicID

	if err := api.Update(ctx, elem2.PublicID, elem2); err != nil {
		tests.Failed("Successfully updated record for TwoFactorSession into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for TwoFactorSession into db.")
}

// TestTwoFactorSessionDelete validates the removal of a TwoFactorSession
// record from a mongodb.
func TestTwoFactorSessionDelete(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTwoFactorSessionJSON(fixtures.TwoFactorSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TwoFactorSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TwoFactorSession record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TwoFactorSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TwoFactorSession into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for TwoFactorSession into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for TwoFactorSession into db.")

	if _, err = api.Get(ctx, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for TwoFactorSession into db.")
	}
	tests.Passed("Successfully failed to get deleted record for TwoFactorSession into db.")
}
