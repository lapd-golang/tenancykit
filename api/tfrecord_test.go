package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/resources/tfrecordapi"
	userFixtures "github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"

	"context"

	"github.com/influx6/faux/httputil/httptesting"

	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	"github.com/gokit/tenancykit/pkg/resources/tfrecordapi/fixtures"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestTwoFactorRecordAPI(t *testing.T) {
	m := metrics.New()
	tfdb := mock.TFRecordBackend()
	userdb := mock.UserBackend()
	tf := api.NewTwoFactorAPI(m, tfdb, userdb)

	userCreateBody, err := userFixtures.LoadCreateJSON(userFixtures.UserCreateJSON)
	if err != nil {
		tests.FailedWithError(err, "Should successfully loaded create user fixture")
	}
	tests.Passed("Should successfully loaded create user fixture")

	userCreateBody.TwoFactorAuth = false
	userCreateBody.PasswordConfirm = userCreateBody.Password

	users := backends.UserBackend{UserDBBackend: userdb}
	userRecord, err := users.Create(context.Background(), userCreateBody)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded user record")
	}
	tests.Passed("Should have successfully loaded user record")

	testUserEnabledTwoFactor(t, userRecord, tf, tfdb)
	testUserQRTwoFactor(t, userRecord, tf, tfdb)
	testUserSecretCodeTwoFactor(t, userRecord, tf, tfdb)
	testUserDisabledTwoFactor(t, userRecord, tf, tfdb)
	testTwoFactorRecordCreate(t, tf, tfdb)
	testTwoFactorRecordCount(t, tf, tfdb)
	testTwoFactorRecordGetAll(t, tf, tfdb)
	testTwoFactorRecordGet(t, tf, tfdb)
	testTwoFactorRecordUpdate(t, tf, tfdb)
	testTwoFactorRecordDelete(t, tf, tfdb)
}

func testUserEnabledTwoFactor(t *testing.T, user pkg.User, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When enabling twofactor auth for user with TwoFactorRecordAPI")
	{
		response := httptest.NewRecorder()
		resource := httptesting.NewRequest("INFO", "/tfrecords/"+user.PublicID, nil, response)
		resource.Set("public_id", user.PublicID)
		if err := tf.EnableTwoFactor(resource); err != nil {
			tests.FailedWithError(err, "Should have successfully made info request")
		}
		tests.Passed("Should have successfully created record")

		if response.Code != http.StatusNoContent {
			tests.Failed("Should have received status code 204")
		}
		tests.Passed("Should have received status code 204")

		if _, err := pkg.GetTwoFactorRecord(resource); err != nil {
			tests.FailedWithError(err, "Should have new tf record in context")
		}
		tests.Passed("Should have new tf record in context")
	}
}

func testUserDisabledTwoFactor(t *testing.T, user pkg.User, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When disabling twofactor auth for user with TwoFactorRecordAPI")
	{
		response := httptest.NewRecorder()
		resource := httptesting.NewRequest("INFO", "/tfrecords/"+user.PublicID, nil, response)
		resource.Set("public_id", user.PublicID)
		if err := tf.DisableTwoFactor(resource); err != nil {
			tests.FailedWithError(err, "Should have successfully made info request")
		}
		tests.Passed("Should have successfully created record")

		if response.Code != http.StatusNoContent {
			tests.Failed("Should have received status code 204")
		}
		tests.Passed("Should have received status code 204")

		if _, err := db.GetByField(context.Background(), "user_id", user.PublicID); err == nil {
			tests.Failed("Should have failed to find user tf record")
		}
		tests.Passed("Should have failed to find user tf record")
	}
}

func testUserSecretCodeTwoFactor(t *testing.T, user pkg.User, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When retrieving qr for user's twofactor auth with TwoFactorRecordAPI")
	{
		response := httptest.NewRecorder()
		resource := httptesting.NewRequest("INFO", "/tfrecords/"+user.PublicID, nil, response)
		resource.Set("public_id", user.PublicID)
		if err := tf.TwoFactorQRCode(resource); err != nil {
			tests.FailedWithError(err, "Should have successfully made info request")
		}
		tests.Passed("Should have successfully created record")

		if response.Code != http.StatusOK {
			tests.Failed("Should have received status code 200")
		}
		tests.Passed("Should have received status code 200")

		if response.Header().Get("Content-Type") != "text/plain" {
			tests.Failed("Should have received content type 'text/plain'")
		}
		tests.Passed("Should have received content type 'text/plain'")

		if response.Body.Len() == 0 {
			tests.Failed("Should have received response body with content")
		}
		tests.Passed("Should have received response body with content")
	}
}

func testUserQRTwoFactor(t *testing.T, user pkg.User, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When retrieving qr for user's twofactor auth with TwoFactorRecordAPI")
	{
		response := httptest.NewRecorder()
		resource := httptesting.NewRequest("INFO", "/tfrecords/"+user.PublicID, nil, response)
		resource.Set("public_id", user.PublicID)
		if err := tf.TwoFactorQRImage(resource); err != nil {
			tests.FailedWithError(err, "Should have successfully made info request")
		}
		tests.Passed("Should have successfully created record")

		if response.Code != http.StatusOK {
			tests.Failed("Should have received status code 204")
		}
		tests.Passed("Should have received status code 204")

		if response.Header().Get("Content-Type") != "image/png" {
			tests.Failed("Should have received content type 'image/png'")
		}
		tests.Passed("Should have received content type 'image/png'")

		if response.Body.Len() == 0 {
			tests.Failed("Should have received response body with content")
		}
		tests.Passed("Should have received response body with content")
	}
}

func testTwoFactorRecordCount(t *testing.T, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When getting info on records using the TwoFactorRecordAPI")
	{
		infoResponse := httptest.NewRecorder()
		infoResource := httptesting.NewRequest("INFO", "/tfrecords", nil, infoResponse)
		if err := tf.Info(infoResource); err != nil {
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

		var info tfrecordapi.TFRecordInfo
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

func testTwoFactorRecordCreate(t *testing.T, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When creating a twofactor record using the TwoFactorRecordAPI")
	{
		createResponse := httptest.NewRecorder()
		createResource := httptesting.Post("/tfrecords", bytes.NewBufferString(fixtures.TFRecordCreateJSON), createResponse)
		if err := tf.Create(createResource); err != nil {
			tests.Info("JSON: %+s", fixtures.TFRecordCreateJSON)
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

		if _, err := fixtures.LoadTFRecordJSON(createResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new record response")
		}
		tests.Passed("Should have successfully received new record response")
	}
}

func testTwoFactorRecordGetAll(t *testing.T, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When retrieving all twofactor records using the TwoFactorRecordAPI")
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
		getAll := httptesting.Get("/tfrecords/", nil, getResponse)
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

		var records tfrecordapi.TFRecordRecords
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

func testTwoFactorRecordGet(t *testing.T, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When retrieving a twofactor record using the TwoFactorRecordAPI")
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
		getRecord := httptesting.Post("/tfrecords/"+record.PublicID, nil, getResponse)
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

		if _, err = fixtures.LoadTFRecordJSON(getResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new record response")
		}
		tests.Passed("Should have successfully received new record response")
	}
}

func testTwoFactorRecordUpdate(t *testing.T, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When updating a twofactor record using the TwoFactorRecordAPI")
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

		beforeDomain := record.Domain
		record.Domain = "wombat.gmail.com"

		recordJSON, err := json.Marshal(record)
		if err != nil {
			tests.Info("JSON: %#v", record)
			tests.FailedWithError(err, "Should successfully marshal user record")
		}
		tests.Passed("Should successfully marshal user record")

		getResponse := httptest.NewRecorder()
		getRecord := httptesting.Put("/tfrecords/"+record.PublicID, bytes.NewBuffer(recordJSON), getResponse)
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

		updatedRecord, err := db.Get(context.Background(), record.PublicID)
		if err != nil {
			tests.FailedWithError(err, "Should have succesfully retrieved update record")
		}
		tests.Passed("Should have succesfully retrieved update record")

		if updatedRecord.Domain != record.Domain {
			tests.Info("Before: %+q", beforeDomain)
			tests.Info("After: %+q", updatedRecord.Domain)
			tests.Info("Expected: %+q", record.Domain)
			tests.Failed("Should have successfully update record field")
		}
		tests.Passed("Should have successfully update record field")
	}
}

func testTwoFactorRecordDelete(t *testing.T, tf api.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When deleting a twofactor record using the TwoFactorRecordAPI")
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
		getRecord := httptesting.Delete("/tfrecords/"+record.PublicID, nil, getResponse)
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

		if _, err := db.Get(context.Background(), record.PublicID); err == nil {
			tests.Failed("Should have succesfully failed to get deleted record")
		}
		tests.Passed("Should have succesfully failed to get deleted record")
	}
}
