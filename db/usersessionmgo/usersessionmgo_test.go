package usersessionmgo_test

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

	mdb "github.com/gokit/tenancykit/db/usersessionmgo"
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

	testCol = "usersession_test_collection"
)

// TestGetUserSession validates the retrieval of a UserSession
// record from a mongodb.
func TestGetUserSession(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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
// record from a mongodb.
func TestGetAllUserSession(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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

// TestGetAllUserSessionOrderBy validates the retrieval of all UserSession
// record from a mongodb.
func TestGetAllUserSessionByOrder(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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
	api := mdb.New(testCol, events, mongo.New(config))

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
// record with a mongodb.
func TestUserSessionUpdate(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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
}

// TestUserSessionDelete validates the removal of a UserSession
// record from a mongodb.
func TestUserSessionDelete(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

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
