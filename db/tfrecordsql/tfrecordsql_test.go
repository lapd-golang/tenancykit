package tfrecordsql_test

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

	sqldb "github.com/gokit/tenancykit/db/tfrecordsql"
)

var (
	events = metrics.New(custom.StackDisplay(os.Stdout))

	config = sql.Config{
		DBName:       os.Getenv("tenancykit_SQL_TEST_DB"),
		User:         os.Getenv("tenancykit_SQL_TEST_USER"),
		DBIP:         os.Getenv("tenancykit_SQL_TEST_ADDR"),
		DBPort:       os.Getenv("tenancykit_SQL_TEST_PORT"),
		DBDriver:     os.Getenv("tenancykit_SQL_TEST_Driver"),
		UserPassword: os.Getenv("tenancykit_SQL_TEST_PASSWORD"),
	}

	testCol = "tfrecord_test"
)

// TestGetTFRecord validates the retrieval of a TFRecord
// record from a sqldb.
func TestGetTFRecord(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tfrecordCreateJSON)
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
// record from a sqldb.
func TestGetAllTFRecord(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tfrecordCreateJSON)
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

// TestTFRecordCreate validates the creation of a TFRecord
// record with a sqldb.
func TestTFRecordCreate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tfrecordCreateJSON)
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
// record with a sqldb.
func TestTFRecordUpdate(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tfrecordCreateJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for TFRecord record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for TFRecord record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully added record for TFRecord into db.")

	//TODO: Update something.

	if err := api.Update(ctx, elem.PublicID, elem); err != nil {
		tests.Failed("Successfully updated record for TFRecord into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for TFRecord into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for TFRecord from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for TFRecord from db.")
}

// TestTFRecordDelete validates the removal of a TFRecord
// record from a sqldb.
func TestTFRecordDelete(t *testing.T) {
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx := context.WithTimeout(context.NewValueBag(), 10*time.Second)

	elem, err := loadJSONFor(tfrecordCreateJSON)
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
}
