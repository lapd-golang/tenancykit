package tenancykit_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gokit/tenancykit/pkg/resources/twofactorsessionapi"

	"github.com/gokit/tenancykit/pkg/resources/twofactorsessionapi/fixtures"
	"github.com/influx6/faux/context"
	"github.com/influx6/faux/httputil/httptesting"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestTwoFactorSessionAPI(t *testing.T) {
	m := metrics.New()
	tfdb := mock.TFSessionBackend()
	tfset := backends.TokenBackend{TokenDBBackend: mock.TokenSetBackend()}
	tf := tenancykit.NewTwoFactorSessionAPI(m, tfset, tfdb)

	testTwoFactorSessionCreate(t, tf, tfdb)
	testTwoFactorSessionGetAll(t, tf, tfdb)
	testTwoFactorSessionGet(t, tf, tfdb)
	testTwoFactorSessionUpdate(t, tf, tfdb)
	testTwoFactorSessionDelete(t, tf, tfdb)
}

func testTwoFactorSessionCreate(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When creating a twofactor session record using the TwoFactorSessionAPI")
	{
		createResponse := httptest.NewRecorder()
		createResource := httptesting.Post("/tfsessions", bytes.NewBufferString(fixtures.TwoFactorSessionCreateJSON), createResponse)
		if err := tf.Create(createResource); err != nil {
			tests.Info("JSON: %+s", fixtures.TwoFactorSessionCreateJSON)
			tests.FailedWithError(err, "Should have successfully created record")
		}
		tests.Passed("Should have successfully created record")

		if createResponse.Code != http.StatusCreated {
			tests.Failed("Should have received Status 201")
		}
		tests.Passed("Should have received Status 201")

		if createResponse.Body == nil {
			tests.Failed("Should have received body response")
		}
		tests.Passed("Should have received body response")

		if _, err := fixtures.LoadTwoFactorSessionJSON(createResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new record response")
		}
		tests.Passed("Should have successfully received new record response")
	}
}

func testTwoFactorSessionGetAll(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When getting all twofactor session record using the TwoFactorSessionAPI")
	{
		_, total, err := db.GetAll(context.New(), "", "", 0, 0)
		if err != nil {
			tests.FailedWithError(err, "Should have retrieved all results from backend")
		}
		tests.Passed("Should have retrieved all results from backend")

		if total == 0 {
			tests.Failed("Should have received atleast one record from backend")
		}
		tests.Passed("Should have received atleast one record from backend")

		getResponse := httptest.NewRecorder()
		getAll := httptesting.Get("/tfsessions/", nil, getResponse)
		if err := tf.GetAll(getAll); err != nil {
			tests.FailedWithError(err, "Should have successfully created record")
		}
		tests.Passed("Should have successfully created record")

		if getResponse.Code != http.StatusOK {
			tests.Failed("Should have received Status 202")
		}
		tests.Passed("Should have received Status 202")

		if getResponse.Body == nil {
			tests.Failed("Should have received body response")
		}
		tests.Passed("Should have received body response")

		var records twofactorsessionapi.TwoFactorSessionRecords
		if err = json.Unmarshal(getResponse.Body.Bytes(), &records); err != nil {
			tests.Info("Records: %+q", getResponse.Body.String())
			tests.FailedWithError(err, "Should have successfully received records response")
		}
		tests.Passed("Should have successfully received new records")

		if records.TotalRecords != total {
			tests.Failed("Should have retrieved same number of records from db")
		}
		tests.Passed("Should have retrieved same number of records from db")
	}
}

func testTwoFactorSessionGet(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When retrieving a twofactor session record using the TwoFactorSessionAPI")
	{
		records, total, err := db.GetAll(context.New(), "", "", 0, 0)
		if err != nil {
			tests.FailedWithError(err, "Should have retrieved all results from backend")
		}
		tests.Passed("Should have retrieved all results from backend")

		if total == 0 {
			tests.Failed("Should have received atleast one record from backend")
		}
		tests.Passed("Should have received atleast one record from backend")

		record := records[0]

		getResponse := httptest.NewRecorder()
		getRecord := httptesting.Post("/tfsessions/"+record.PublicID, nil, getResponse)
		getRecord.Bag().Set("public_id", record.PublicID)

		if err := tf.Get(getRecord); err != nil {
			tests.Info("Record: %#v", record)
			tests.FailedWithError(err, "Should have successfully created record")
		}
		tests.Passed("Should have successfully created record")

		if getResponse.Code != http.StatusOK {
			tests.Failed("Should have received Status 202")
		}
		tests.Passed("Should have received Status 202")

		if getResponse.Body == nil {
			tests.Failed("Should have received body response")
		}
		tests.Passed("Should have received body response")

		if _, err = fixtures.LoadTwoFactorSessionJSON(getResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new record response")
		}
		tests.Passed("Should have successfully received new record response")
	}
}

func testTwoFactorSessionUpdate(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When updating a twofactor session record using the TwoFactorSessionAPI")
	{
		records, total, err := db.GetAll(context.New(), "", "", 0, 0)
		if err != nil {
			tests.FailedWithError(err, "Should have retrieved all results from backend")
		}
		tests.Passed("Should have retrieved all results from backend")

		if total == 0 {
			tests.Failed("Should have received atleast one record from backend")
		}
		tests.Passed("Should have received atleast one record from backend")

		record := records[0]

		beforeDomain := record.Done
		record.Done = true

		recordJSON, err := json.Marshal(record)
		if err != nil {
			tests.Info("JSON: %#v", record)
			tests.FailedWithError(err, "Should successfully marshal user record")
		}
		tests.Passed("Should successfully marshal user record")

		getResponse := httptest.NewRecorder()
		getRecord := httptesting.Put("/tfsessions/"+record.PublicID, bytes.NewBuffer(recordJSON), getResponse)
		getRecord.Bag().Set("public_id", record.PublicID)

		if err := tf.Update(getRecord); err != nil {
			tests.Info("Record: %#v", record)
			tests.FailedWithError(err, "Should have successfully created record")
		}
		tests.Passed("Should have successfully created record")

		if getResponse.Code != http.StatusNoContent {
			tests.Failed("Should have received Status 202")
		}
		tests.Passed("Should have received Status 202")

		updatedRecord, err := db.Get(context.New(), record.PublicID)
		if err != nil {
			tests.FailedWithError(err, "Should have succesfully retrieved update record")
		}
		tests.Passed("Should have succesfully retrieved update record")

		if updatedRecord.Done != record.Done {
			tests.Info("Before: %+q", beforeDomain)
			tests.Info("After: %+q", updatedRecord.Done)
			tests.Info("Expected: %+q", record.Done)
			tests.Failed("Should have successfully update record field")
		}
		tests.Passed("Should have successfully update record field")
	}
}

func testTwoFactorSessionDelete(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When deleting a twofactor session record using the TwoFactorSessionAPI")
	{
		records, total, err := db.GetAll(context.New(), "", "", 0, 0)
		if err != nil {
			tests.FailedWithError(err, "Should have retrieved all results from backend")
		}
		tests.Passed("Should have retrieved all results from backend")

		if total == 0 {
			tests.Failed("Should have received atleast one record from backend")
		}
		tests.Passed("Should have received atleast one record from backend")

		record := records[0]

		getResponse := httptest.NewRecorder()
		getRecord := httptesting.Delete("/tfsessions/"+record.PublicID, nil, getResponse)
		getRecord.Bag().Set("public_id", record.PublicID)

		if err := tf.Delete(getRecord); err != nil {
			tests.Info("Record: %#v", record)
			tests.FailedWithError(err, "Should have successfully created record")
		}
		tests.Passed("Should have successfully created record")

		if getResponse.Code != http.StatusNoContent {
			tests.Failed("Should have received Status 202")
		}
		tests.Passed("Should have received Status 202")

		if _, err := db.Get(context.New(), record.PublicID); err == nil {
			tests.Failed("Should have succesfully failed to get deleted record")
		}
		tests.Passed("Should have succesfully failed to get deleted record")
	}
}
