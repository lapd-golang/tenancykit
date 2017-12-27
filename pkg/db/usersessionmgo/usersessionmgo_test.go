package usersessionmgo_test

import (
	"os"

	"time"

	"context"

	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	mdb "github.com/gokit/tenancykit/pkg/db/usersessionmgo"

	fixtures "github.com/gokit/tenancykit/pkg/db/usersessionmgo/fixtures"
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

	testCol = "usersession_test_collection"
)

// TestGetUserSession validates the retrieval of a UserSession
// record from a mongodb.
func TestGetUserSession(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserSessionJSON(fixtures.UserSessionJSON)
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
// record from a mongodb.
func TestGetAllUserSession(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserSessionJSON(fixtures.UserSessionJSON)
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

// TestGetAllUserSessionOrderBy validates the retrieval of all UserSession
// record from a mongodb.
func TestGetAllUserSessionByOrder(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserSessionJSON(fixtures.UserSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for UserSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for UserSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for UserSession into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
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
// record with a mongodb.
func TestUserSessionCreate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserSessionJSON(fixtures.UserSessionJSON)
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
// record with a mongodb.
func TestUserSessionUpdate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserSessionJSON(fixtures.UserSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for UserSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for UserSession record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully added record for UserSession into db.")

	elem2, err := fixtures.LoadUserSessionJSON(fixtures.UserSessionJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for UserSession record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for UserSession record")

	elem2.PublicID = elem.PublicID

	if err := api.Update(ctx, elem2.PublicID, elem2); err != nil {
		tests.Failed("Successfully updated record for UserSession into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for UserSession into db.")
}

// TestUserSessionDelete validates the removal of a UserSession
// record from a mongodb.
func TestUserSessionDelete(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserSessionJSON(fixtures.UserSessionJSON)
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

	if _, err = api.Get(ctx, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for UserSession into db.")
	}
	tests.Passed("Successfully failed to get deleted record for UserSession into db.")
}
