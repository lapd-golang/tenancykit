package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg"
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

	users := api.UserBackend{UserDBBackend: userdb}
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
