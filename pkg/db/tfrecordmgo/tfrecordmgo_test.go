package tfrecordmgo_test

import (
	"os"

	"time"

	"context"

	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	mdb "github.com/gokit/tenancykit/pkg/db/tfrecordmgo"

	fixtures "github.com/gokit/tenancykit/pkg/db/tfrecordmgo/fixtures"
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

	testCol = "tfrecord_test_collection"
)

// TestGetTFRecord validates the retrieval of a TFRecord
// record from a mongodb.
func TestGetTFRecord(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTFRecordJSON(fixtures.TFRecordJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TFRecord record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TFRecord record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TFRecord into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for TFRecord from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for TFRecord from db.")
}

// TestGetAllTFRecord validates the retrieval of all TFRecord
// record from a mongodb.
func TestGetAllTFRecord(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTFRecordJSON(fixtures.TFRecordJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TFRecord record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TFRecord record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TFRecord into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for TFRecord from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for TFRecord from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for TFRecord from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for TFRecord from db.")
}

// TestGetAllTFRecordOrderBy validates the retrieval of all TFRecord
// record from a mongodb.
func TestGetAllTFRecordByOrder(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTFRecordJSON(fixtures.TFRecordJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TFRecord record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TFRecord record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TFRecord into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
	if err != nil {
		tests.Failed("Successfully retrieved all records for TFRecord from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for TFRecord from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for TFRecord from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for TFRecord from db.")
}

// TestTFRecordCreate validates the creation of a TFRecord
// record with a mongodb.
func TestTFRecordCreate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTFRecordJSON(fixtures.TFRecordJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TFRecord record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TFRecord record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TFRecord into db.")
}

// TestTFRecordUpdate validates the update of a TFRecord
// record with a mongodb.
func TestTFRecordUpdate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTFRecordJSON(fixtures.TFRecordJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TFRecord record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TFRecord record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TFRecord into db.")

	elem2, err := fixtures.LoadTFRecordJSON(fixtures.TFRecordJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TFRecord record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TFRecord record")

	elem2.PublicID = elem.PublicID

	if err := api.Update(ctx, elem2.PublicID, elem2); err != nil {
		tests.Failed("Successfully updated record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for TFRecord into db.")
}

// TestTFRecordDelete validates the removal of a TFRecord
// record from a mongodb.
func TestTFRecordDelete(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadTFRecordJSON(fixtures.TFRecordJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TFRecord record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TFRecord record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TFRecord into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for TFRecord into db.")

	if _, err = api.Get(ctx, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for TFRecord into db.")
	}
	tests.Passed("Successfully failed to get deleted record for TFRecord into db.")
}
