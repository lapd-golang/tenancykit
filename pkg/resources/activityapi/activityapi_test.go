package activityapi_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"encoding/json"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/httputil/httptesting"

	"github.com/influx6/faux/metrics/custom"

	httpapi "github.com/gokit/tenancykit/pkg/resources/activityapi"

	"github.com/gokit/tenancykit/pkg/resources/activityapi/fixtures"
)

func TestActivityAPI(t *testing.T) {
	logs := metrics.New()
	if testing.Verbose() {
		logs = metrics.New(custom.StackDisplay(os.Stdout))
	}

	backend := newMocker()
	api := httpapi.New(logs, backend)

	testCreate(t, api, backend)
	testInfo(t, api, backend)
	testGet(t, api, backend)
	testGetAll(t, api, backend)
	testUpdate(t, api, backend)
	testDelete(t, api, backend)
}

func testInfo(t *testing.T, api httpapi.ActivityHTTP, db httpapi.ActivityBackend) {
	tests.Header("When we create a Activity with ActivityAPI")
	{
		infoResponse := httptest.NewRecorder()
		infoResource := httptesting.NewRequest("INFO", "/activity", nil, infoResponse)
		if err := api.Info(infoResource); err != nil {
			tests.FailedWithError(err, "Should have successfully made info request")
		}
		tests.Passed("Should have successfully created record")

		if infoResponse.Code != http.StatusOK {
			tests.Failed("Should have received Status 200")
		}
		tests.Passed("Should have received Status 200")

		if infoResponse.Body == nil {
			tests.Failed("Should have received body response")
		}
		tests.Passed("Should have received body response")

		var info httpapi.ActivityInfo
		if err := json.Unmarshal(infoResponse.Body.Bytes(), &info); err != nil {
			tests.FailedWithError(err, "Should have successfully collected record info")
		}
		tests.Passed("Should have successfully collected record info")

		if info.Total == 0 {
			tests.Failed("Should have atleast one record in backend")
		}
		tests.Passed("Should have atleast one record in backend")
	}
}

func testCreate(t *testing.T, api httpapi.ActivityHTTP, db httpapi.ActivityBackend) {
	tests.Header("When we get info for all Activity records with ActivityAPI")
	{
		createResponse := httptest.NewRecorder()
		createResource := httptesting.Post("/activity", bytes.NewBufferString(fixtures.ActivityCreateJSON), createResponse)
		if err := api.Create(createResource); err != nil {
			tests.Info("JSON: %+s", fixtures.ActivityCreateJSON)
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

		if _, err := fixtures.LoadActivityJSON(createResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new record response")
		}
		tests.Passed("Should have successfully received new record response")
	}
}

func testGetAll(t *testing.T, api httpapi.ActivityHTTP, db httpapi.ActivityBackend) {
	tests.Header("When we get all Activity records with ActivityAPI")
	{
		_, total, err := db.GetAll(context.Background(), "", "", 0, 0)
		if err != nil {
			tests.FailedWithError(err, "Should have retrieved all results from backend")
		}
		tests.Passed("Should have retrieved all results from backend")

		if total == 0 {
			tests.Failed("Should have received atleast one record from backend")
		}
		tests.Passed("Should have received atleast one record from backend")

		getResponse := httptest.NewRecorder()
		getAll := httptesting.Get("/activity/", nil, getResponse)
		if err := api.GetAll(getAll); err != nil {
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

		var records httpapi.ActivityRecords
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

func testGet(t *testing.T, api httpapi.ActivityHTTP, db httpapi.ActivityBackend) {
	tests.Header("When we get one Activity record with ActivityAPI")
	{
		records, total, err := db.GetAll(context.Background(), "", "", 0, 0)
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
		getRecord := httptesting.Post("/activity/"+record.PublicID, nil, getResponse)
		getRecord.Bag().Set("public_id", record.PublicID)

		if err := api.Get(getRecord); err != nil {
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

		if _, err = fixtures.LoadActivityJSON(getResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new record response")
		}
		tests.Passed("Should have successfully received new record response")
	}
}

func testUpdate(t *testing.T, api httpapi.ActivityHTTP, db httpapi.ActivityBackend) {
	tests.Header("When we update one Activity record with ActivityAPI")
	{
		records, total, err := db.GetAll(context.Background(), "", "", 0, 0)
		if err != nil {
			tests.FailedWithError(err, "Should have retrieved all results from backend")
		}
		tests.Passed("Should have retrieved all results from backend")

		if total == 0 {
			tests.Failed("Should have received atleast one record from backend")
		}
		tests.Passed("Should have received atleast one record from backend")

		record := records[0]

		updateRecord, err := fixtures.LoadUpdateJSON(fixtures.ActivityUpdateJSON)
		if err != nil {
			tests.FailedWithError(err, "Should have successfully loaded update data")
		}
		tests.Passed("Should have successfully loaded update data")

		updateRecord.PublicID = record.PublicID

		recordJSON, err := json.Marshal(record)
		if err != nil {
			tests.Info("JSON: %#v", record)
			tests.FailedWithError(err, "Should successfully marshal user record")
		}
		tests.Passed("Should successfully marshal user record")

		getResponse := httptest.NewRecorder()
		getRecord := httptesting.Put("/activity/"+record.PublicID, bytes.NewBuffer(recordJSON), getResponse)
		getRecord.Bag().Set("public_id", record.PublicID)

		if err := api.Update(getRecord); err != nil {
			tests.Info("Record: %#v", record)
			tests.FailedWithError(err, "Should have successfully created record")
		}
		tests.Passed("Should have successfully created record")

		if getResponse.Code != http.StatusNoContent {
			tests.Failed("Should have received Status 202")
		}
		tests.Passed("Should have received Status 202")

		if _, err = db.Get(context.Background(), record.PublicID); err != nil {
			tests.FailedWithError(err, "Should have successfully retrieved update record")
		}
		tests.Passed("Should have successfully retrieved update record")
	}
}

func testDelete(t *testing.T, api httpapi.ActivityHTTP, db httpapi.ActivityBackend) {
	tests.Header("When we delete one Activity record with ActivityAPI")
	{
		records, total, err := db.GetAll(context.Background(), "", "", 0, 0)
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
		getRecord := httptesting.Delete("/activity/"+record.PublicID, nil, getResponse)
		getRecord.Bag().Set("public_id", record.PublicID)

		if err := api.Delete(getRecord); err != nil {
			tests.Info("Record: %#v", record)
			tests.FailedWithError(err, "Should have successfully created record")
		}
		tests.Passed("Should have successfully created record")

		if getResponse.Code != http.StatusNoContent {
			tests.Failed("Should have received Status 202")
		}
		tests.Passed("Should have received Status 202")

		if _, err := db.Get(context.Background(), record.PublicID); err == nil {
			tests.Failed("Should have successfully failed to get deleted record")
		}
		tests.Passed("Should have successfully failed to get deleted record")
	}
}
