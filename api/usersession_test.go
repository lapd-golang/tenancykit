package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gokit/tenancykit/api"

	"github.com/influx6/faux/httputil/httptesting"

	"github.com/gokit/tenancykit/pkg"

	"context"

	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	tenantFixtures "github.com/gokit/tenancykit/pkg/resources/tenantapi/fixtures"
	userFixtures "github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"
	"github.com/gokit/tenancykit/pkg/resources/usersessionapi/fixtures"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestUserSessionAPI(t *testing.T) {
	m := metrics.New()
	tfdb := mock.TFRecordBackend()
	ttdb := mock.TenantDBBackend()
	tfsdb := mock.TFSessionBackend()
	ufsdb := mock.UserSessionBackend()
	udb := mock.UserBackend()

	tf := api.NewUserSessionAPI(m, udb, ufsdb, tfsdb, tfdb)

	createTenant, err := tenantFixtures.LoadCreateJSON(tenantFixtures.TenantCreateJSON)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded tenant record")
	}
	tests.Passed("Should have successfully loaded tenant record")

	tenants := api.TenantBackend{TenantDBBackend: ttdb}
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

	userCreateBody.TwoFactorAuth = false
	userCreateBody.TenantID = tenantRecord.PublicID
	userCreateBody.PasswordConfirm = userCreateBody.Password

	users := api.UserBackend{UserDBBackend: udb}
	userRecord, err := users.Create(context.Background(), userCreateBody)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded user record")
	}
	tests.Passed("Should have successfully loaded user record")

	testUserSessionLoginAndLogout(t, userRecord, userCreateBody, tenantRecord, tf, ufsdb)
	testUserSessionCreate(t, userRecord, userCreateBody, tenantRecord, tf, ufsdb)
}

func testUserSessionLoginAndLogout(t *testing.T, user pkg.User, usercreate pkg.CreateUser, tenant pkg.Tenant, tf api.UserSessionAPI, db types.UserSessionDBBackend) {
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

func testUserSessionCreate(t *testing.T, user pkg.User, usercreate pkg.CreateUser, tenant pkg.Tenant, tf api.UserSessionAPI, db types.UserSessionDBBackend) {
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
