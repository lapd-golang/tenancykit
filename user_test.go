package tenancykit_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"context"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/pkg"
	"github.com/influx6/faux/httputil/httptesting"
	"github.com/influx6/faux/tests"

	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	tenantFixtures "github.com/gokit/tenancykit/pkg/resources/tenantapi/fixtures"
	"github.com/gokit/tenancykit/pkg/resources/userapi"
	"github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"
	"github.com/influx6/faux/metrics"
)

func TestUserAPI(t *testing.T) {

	m := metrics.New()
	userdb := mock.UserBackend()
	tenantdb := mock.TenantDBBackend()
	tfdb := mock.TFRecordBackend()

	users := tenancykit.NewUserAPI(
		pkg.GoogleAuthenticatorUserCodeLength,
		m,
		userdb,
		tenantdb,
		tfdb,
	)

	createTenant, err := tenantFixtures.LoadCreateJSON(tenantFixtures.TenantCreateJSON)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded tenant record")
	}
	tests.Passed("Should have successfully loaded tenant record")

	tenants := backends.TenantBackend{TenantDBBackend: tenantdb}
	tenantRecord, err := tenants.Create(context.Background(), createTenant)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created tenant")
	}
	tests.Passed("Should have successfully created tenants")

	testUserCreate(t, tenantRecord, users, userdb)
	testUserGetAll(t, tenantRecord, users, userdb)
	testUserGet(t, tenantRecord, users, userdb)
	testUserUpdate(t, tenantRecord, users, userdb)
	testUserDelete(t, tenantRecord, users, userdb)
	os.RemoveAll("./keys")
}

func testUserCreate(t *testing.T, tenant pkg.Tenant, users tenancykit.UserAPI, db types.UserDBBackend) {
	tests.Header("When creating new user with User API")
	{
		createBody, err := fixtures.LoadCreateJSON(fixtures.UserCreateJSON)
		if err != nil {
			tests.FailedWithError(err, "Should successfully loaded create user fixture")
		}
		tests.Passed("Should successfully loaded create user fixture")

		createBody.TenantID = tenant.PublicID
		createBody.PasswordConfirm = createBody.Password
		createBodyJSON, err := json.Marshal(createBody)
		if err != nil {
			tests.Info("JSON: %+s", createBodyJSON)
			tests.FailedWithError(err, "Should successfully marshal create user record")
		}
		tests.Passed("Should successfully marshal create user record")

		createResponse := httptest.NewRecorder()
		createUser := httptesting.Post("/users", bytes.NewBuffer(createBodyJSON), createResponse)
		if err := users.Create(createUser); err != nil {
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

		if _, err := fixtures.LoadUserJSON(createResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new user response")
		}
		tests.Passed("Should have successfully received new user response")
	}
}

func testUserGetAll(t *testing.T, tenant pkg.Tenant, users tenancykit.UserAPI, db types.UserDBBackend) {
	tests.Header("When retrieving all users with User API")
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

		getUserResponse := httptest.NewRecorder()
		getAllUser := httptesting.Get("/users/", nil, getUserResponse)
		if err := users.GetAll(getAllUser); err != nil {
			tests.FailedWithError(err, "Should have successfully created user")
		}
		tests.Passed("Should have successfully created user")

		if getUserResponse.Code != http.StatusOK {
			tests.Failed("Should have received Status 202")
		}
		tests.Passed("Should have received Status 202")

		if getUserResponse.Body == nil {
			tests.Failed("Should have received body response")
		}
		tests.Passed("Should have received body response")

		var records userapi.UserRecords
		if err = json.Unmarshal(getUserResponse.Body.Bytes(), &records); err != nil {
			tests.Info("Users: %+q", getUserResponse.Body.String())
			tests.FailedWithError(err, "Should have successfully received users response")
		}
		tests.Passed("Should have successfully received new users")

		if records.TotalRecords != total {
			tests.Failed("Should have retrieved same number of users from db")
		}
		tests.Passed("Should have retrieved same number of users from db")
	}
}

func testUserGet(t *testing.T, tenant pkg.Tenant, users tenancykit.UserAPI, db types.UserDBBackend) {
	tests.Header("When retrieving user with User API")
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

		user := records[0]

		getUserResponse := httptest.NewRecorder()
		getUser := httptesting.Post("/users/"+user.PublicID, nil, getUserResponse)
		getUser.Bag().Set("public_id", user.PublicID)

		if err := users.Get(getUser); err != nil {
			tests.Info("User: %#v", user)
			tests.FailedWithError(err, "Should have successfully created user")
		}
		tests.Passed("Should have successfully created user")

		if getUserResponse.Code != http.StatusOK {
			tests.Failed("Should have received Status 202")
		}
		tests.Passed("Should have received Status 202")

		if getUserResponse.Body == nil {
			tests.Failed("Should have received body response")
		}
		tests.Passed("Should have received body response")

		if _, err = fixtures.LoadUserJSON(getUserResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new user response")
		}
		tests.Passed("Should have successfully received new user response")
	}
}

func testUserUpdate(t *testing.T, tenant pkg.Tenant, users tenancykit.UserAPI, db types.UserDBBackend) {
	tests.Header("When updating user with User API")
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

		user := records[0]

		var userUpdate pkg.UpdateUser

		beforeUsername := user.Email
		userUpdate.Email = "djackman201@gmail.com"
		userUpdate.Password = "bombabastick"
		userUpdate.PasswordConfirm = userUpdate.Password

		updateBodyJSON, err := json.Marshal(userUpdate)
		if err != nil {
			tests.Info("JSON: %+s", user)
			tests.FailedWithError(err, "Should successfully marshal user record")
		}
		tests.Passed("Should successfully marshal user record")

		updateUserResponse := httptest.NewRecorder()
		updateUser := httptesting.Put("/users/"+user.PublicID, bytes.NewReader(updateBodyJSON), updateUserResponse)
		updateUser.Bag().Set("public_id", user.PublicID)

		if err := users.Update(updateUser); err != nil {
			tests.Info("User: %#v", updateBodyJSON)
			tests.FailedWithError(err, "Should have successfully created user")
		}
		tests.Passed("Should have successfully created user")

		if updateUserResponse.Code != http.StatusNoContent {
			tests.Failed("Should have received Status 204")
		}
		tests.Passed("Should have received Status 204")

		updatedRecord, err := db.Get(context.Background(), user.PublicID)
		if err != nil {
			tests.FailedWithError(err, "Should have succesfully retrieved update record")
		}
		tests.Passed("Should have succesfully retrieved update record")

		if updatedRecord.Email != userUpdate.Email {
			tests.Info("Before: %+q", beforeUsername)
			tests.Info("After: %+q", updatedRecord.Email)
			tests.Info("Expected: %+q", userUpdate.Email)
			tests.Failed("Should have successfully update record field")
		}
		tests.Passed("Should have successfully update record field")

		if err := updatedRecord.Authenticate(userUpdate.Password); err != nil {
			tests.FailedWithError(err, "Should have successfully confirmed password change")
		}
		tests.Passed("Should have successfully confirmed password change")
	}
}

func testUserDelete(t *testing.T, tenant pkg.Tenant, users tenancykit.UserAPI, db types.UserDBBackend) {
	tests.Header("When deleting user with User API")
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

		user := records[0]

		deleteUserResponse := httptest.NewRecorder()
		deleteUser := httptesting.Delete("/users/"+user.PublicID, nil, deleteUserResponse)
		deleteUser.Bag().Set("public_id", user.PublicID)

		if err := users.Delete(deleteUser); err != nil {
			tests.Info("User: %#v", user)
			tests.FailedWithError(err, "Should have successfully created user")
		}
		tests.Passed("Should have successfully created user")

		if deleteUserResponse.Code != http.StatusNoContent {
			tests.Failed("Should have received Status 204")
		}
		tests.Passed("Should have received Status 204")

		if _, err := db.Get(context.Background(), user.PublicID); err == nil {
			tests.Failed("Should have succesfully failed to get deleted record")
		}
		tests.Passed("Should have succesfully failed to get deleted record")
	}
}
