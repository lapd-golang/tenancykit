package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg/resources/usersessionapi"

	"github.com/gokit/tenancykit/pkg/resources/tfrecordapi"
	"github.com/influx6/faux/httputil/httptesting"

	"github.com/gokit/tenancykit/pkg"

	"context"

	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	tenantFixtures "github.com/gokit/tenancykit/pkg/resources/tenantapi/fixtures"
	userFixtures "github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"
	"github.com/gokit/tenancykit/pkg/resources/usersessionapi/fixtures"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestMultiTenantUserSessionAPI(t *testing.T) {
	m := metrics.New()
	tfdb := mock.TFRecordBackend()
	ttdb := mock.TenantDBBackend()
	tfsdb := mock.TFSessionBackend()
	ufsdb := mock.UserSessionBackend()
	udb := mock.UserBackend()

	tf := api.NewUserSessionAPI(m, udb, ufsdb, tfsdb, ttdb, tfdb)

	createTenant, err := tenantFixtures.LoadCreateJSON(tenantFixtures.TenantCreateJSON)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded tenant record")
	}
	tests.Passed("Should have successfully loaded tenant record")

	tenants := backends.TenantBackend{TenantDBBackend: ttdb}
	tenantRecord, err := tenants.Create(context.Background(), createTenant)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created tenant")
	}
	tests.Passed("Should have successfully created tenants")

	userCreateBody, err := userFixtures.LoadCreateJSON(userFixtures.UserCreateJSON)
	if err != nil {
		tests.FailedWithError(err, "Should successfully loaded create user fixture")
	}
	tests.Passed("Should successfully loaded create user fixture")

	userCreateBody.TenantID = tenantRecord.PublicID
	userCreateBody.PasswordConfirm = userCreateBody.Password

	users := backends.UserBackend{UserDBBackend: udb}
	userRecord, err := users.Create(context.Background(), userCreateBody)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded user record")
	}
	tests.Passed("Should have successfully loaded user record")

	testMultiTenantSessionLoginAndLogout(t, userRecord, userCreateBody, tenantRecord, tf, ufsdb)
	testMultiTenantSessionCreate(t, userRecord, userCreateBody, tenantRecord, tf, ufsdb)
	testMultiTenantSessionCount(t, tf, ufsdb)
	testMultiTenantSessionGetAll(t, userRecord, tenantRecord, tf, ufsdb)
	testMultiTenantSessionGet(t, userRecord, tenantRecord, tf, ufsdb)
	testMultiTenantSessionUpdate(t, userRecord, tenantRecord, tf, ufsdb)
	testMultiTenantSessionDelete(t, userRecord, tenantRecord, tf, ufsdb)

	os.RemoveAll("./keys")
}

func testMultiTenantSessionCount(t *testing.T, us api.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When getting info on records using the UserSessionAPI")
	{
		infoResponse := httptest.NewRecorder()
		infoResource := httptesting.NewRequest("INFO", "/tenants", nil, infoResponse)
		if err := us.Info(infoResource); err != nil {
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

		var info usersessionapi.UserSessionInfo
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

func testMultiTenantSessionLoginAndLogout(t *testing.T, user pkg.User, usercreate pkg.CreateUser, tenant pkg.Tenant, tf api.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When login and logging out using the UserSessionAPI")
	{
		var login pkg.CreateUserSession
		login.Email = user.Email
		login.Password = usercreate.Password
		login.Expiration = 24 * time.Hour

		loginBodyJSON, err := json.Marshal(login)
		if err != nil {
			tests.Info("JSON: %+s", loginBodyJSON)
			tests.FailedWithError(err, "Should successfully marshal create user record")
		}
		tests.Passed("Should successfully marshal create user record")

		loginResponse := httptest.NewRecorder()
		logginUser := httptesting.Post("/sessions/login", bytes.NewBuffer(loginBodyJSON), loginResponse)
		if err := tf.Login(logginUser); err != nil {
			tests.FailedWithError(err, "Should have successfully authenticated user")
		}
		tests.Passed("Should have successfully authenticated user")

		if loginResponse.Code != http.StatusOK {
			tests.Info("Received Status: ", loginResponse.Code)
			tests.Failed("Should have received status no 200")
		}
		tests.Passed("Should have received status no 200")

		authHeader := loginResponse.Header().Get("Authorization")
		if strings.TrimSpace(authHeader) == "" {
			tests.Failed("Should have received Authorization header")
		}
		tests.Passed("Should have received Authorization header")

		tests.Info("User Authorization Header: %+q", authHeader)

		var authentication pkg.UserAuthorization
		if err := json.Unmarshal(loginResponse.Body.Bytes(), &authentication); err != nil {
			tests.Info("JSON: %+q", loginResponse.Body.Bytes())
			tests.FailedWithError(err, "Should have successfully unserialized data")
		}
		tests.Passed("Should have successfully unserialized data")

		if auth := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer")); auth != authentication.Token {
			tests.Failed("Should have authentication data match incoming response")
		}
		tests.Passed("Should have authentication data match incoming response")

		newLoginResponse := httptest.NewRecorder()
		newLogginUser := httptesting.Post("/sessions/login", nil, newLoginResponse)
		newLogginUser.Request().Header.Set("Authorization", authHeader)
		if err := tf.GetUser(newLogginUser); err != nil {
			tests.FailedWithError(err, "Should have successfully authenticated user with header")
		}
		tests.Passed("Should have successfully authenticated user with header")

		if newLoginResponse.Code != http.StatusOK {
			tests.Info("Received Status: %d", newLoginResponse.Code)
			tests.Failed("Should have received status no 204")
		}
		tests.Passed("Should have received status no 204")

		newLogoutResponse := httptest.NewRecorder()
		newLogoutUser := httptesting.Post("/sessions/login", nil, newLogoutResponse)
		newLogoutUser.Request().Header.Set("Authorization", authHeader)
		if err := tf.Logout(newLogoutUser); err != nil {
			tests.FailedWithError(err, "Should have successfully logged out authenticated user with header")
		}
		tests.Passed("Should have successfully logged out authenticated user with header")

		if newLogoutResponse.Code != http.StatusNoContent {
			tests.Info("Received Status: ", newLogoutResponse.Code)
			tests.Failed("Should have received status no 204")
		}
		tests.Passed("Should have received status no 204")

		err = tf.GetUser(newLogginUser)
		if err == nil {
			tests.Failed("Should have failed to authenticate user with header")
		}
		tests.PassedWithError(err, "Should have failed to authenticate user with header")
	}
}

func testMultiTenantSessionCreate(t *testing.T, user pkg.User, usercreate pkg.CreateUser, tenant pkg.Tenant, tf api.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When creating a two factor record using the UserSessionAPI")
	{
		createBody, err := fixtures.LoadCreateJSON(fixtures.UserSessionCreateJSON)
		if err != nil {
			tests.FailedWithError(err, "Should successfully loaded create user fixture")
		}
		tests.Passed("Should successfully loaded create user fixture")

		createBody.Email = user.Email
		createBody.Password = usercreate.Password
		createBodyJSON, err := json.Marshal(createBody)
		if err != nil {
			tests.Info("JSON: %+s", createBodyJSON)
			tests.FailedWithError(err, "Should successfully marshal create user record")
		}
		tests.Passed("Should successfully marshal create user record")

		createResponse := httptest.NewRecorder()
		createUser := httptesting.Post("/sessions", bytes.NewBuffer(createBodyJSON), createResponse)
		if err := tf.Create(createUser); err != nil {
			tests.Info("JSON: %+s", createBodyJSON)
			tests.FailedWithError(err, "Should have successfully created user")
		}
		tests.Passed("Should have successfully created user")

		if createResponse.Code != http.StatusCreated {
			tests.Failed("Should have received Status 201")
		}
		tests.Passed("Should have received Status 201")

		if createResponse.Body == nil {
			tests.Failed("Should have received body response")
		}
		tests.Passed("Should have received body response")

		if _, err := fixtures.LoadCreateJSON(createResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new user response")
		}
		tests.Passed("Should have successfully received new user response")
	}
}

func testMultiTenantSessionGetAll(t *testing.T, user pkg.User, tenant pkg.Tenant, tf api.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When retrieving all two factor records using the UserSessionAPI")
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

func testMultiTenantSessionGet(t *testing.T, user pkg.User, tenant pkg.Tenant, tf api.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When retrieving a two factor record using the UserSessionAPI")
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

		if _, err = fixtures.LoadUserSessionJSON(getResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new record response")
		}
		tests.Passed("Should have successfully received new record response")
	}
}

func testMultiTenantSessionUpdate(t *testing.T, user pkg.User, tenant pkg.Tenant, tf api.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When updating a two factor record using the UserSessionAPI")
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

		beforeDomain := record.Token
		record.Token = "wombat.gmail.com"

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

		if updatedRecord.Token != record.Token {
			tests.Info("Before: %+q", beforeDomain)
			tests.Info("After: %+q", updatedRecord.Token)
			tests.Info("Expected: %+q", record.Token)
			tests.Failed("Should have successfully update record field")
		}
		tests.Passed("Should have successfully update record field")
	}
}

func testMultiTenantSessionDelete(t *testing.T, user pkg.User, tenant pkg.Tenant, tf api.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When deleting a two factor record using the UserSessionAPI")
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
