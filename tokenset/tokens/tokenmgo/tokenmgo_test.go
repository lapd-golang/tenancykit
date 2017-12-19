package tokenmgo_test

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

	mdb "github.com/gokit/tenancykit/tokenset/tokens/tokenmgo"
)

var (
	events = metrics.New(custom.StackDisplay(os.Stdout))

	config = mongo.Config{
		Mode:     mgo.Monotonic,
		DB:       os.Getenv("tokens_MONGO_TEST_DB"),
		Host:     os.Getenv("tokens_MONGO_TEST_HOST"),
		User:     os.Getenv("tokens_MONGO_TEST_USER"),
		AuthDB:   os.Getenv("tokens_MONGO_TEST_AUTHDB"),
		Password: os.Getenv("tokens_MONGO_TEST_PASSWORD"),
	}

	testCol = "token_test_collection"
)

// TestGetToken validates the retrieval of a Token
// record from a mongodb.
func TestGetToken(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for Token from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for Token from db.")
}

// TestGetAllToken validates the retrieval of all Token
// record from a mongodb.
func TestGetAllToken(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for Token from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Token from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Token from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Token from db.")
}

// TestGetAllTokenOrderBy validates the retrieval of all Token
// record from a mongodb.
func TestGetAllTokenByOrder(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
	if err != nil {
		tests.Failed("Successfully retrieved all records for Token from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for Token from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for Token from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for Token from db.")
}

// TestTokenCreate validates the creation of a Token
// record with a mongodb.
func TestTokenCreate(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")
}

// TestTokenUpdate validates the update of a Token
// record with a mongodb.
func TestTokenUpdate(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	//TODO: Update something.

	if err := api.Update(ctx, elem.PublicID, elem); err != nil {
		tests.Failed("Successfully updated record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for Token into db.")
}

// TestTokenDelete validates the removal of a Token
// record from a mongodb.
func TestTokenDelete(t *testing.T) {
	api := mdb.New(testCol, events, mongo.New(config))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for Token record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for Token record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully added record for Token into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for Token into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for Token into db.")
}
