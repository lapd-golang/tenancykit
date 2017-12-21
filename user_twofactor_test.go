package tenancykit_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/influx6/faux/httputil/httptesting"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	tenantFixtures "github.com/gokit/tenancykit/pkg/resources/tenantapi/fixtures"
	userFixtures "github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestTwoFactorAuth(t *testing.T) {
	m := metrics.New()
	tfdb := mock.TFRecordBackend()
	ttdb := mock.TenantDBBackend()
	tfsdb := mock.TFSessionBackend()
	ufsdb := mock.UserSessionBackend()
	udb := mock.UserBackend()

	tf := tenancykit.NewUserSessionAPI(m, udb, ufsdb, tfsdb, ttdb, tfdb)

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

	userCreateBody.TwoFactorAuth = true
	userCreateBody.TenantID = tenantRecord.PublicID
	userCreateBody.PasswordConfirm = userCreateBody.Password

	users := backends.UserBackend{UserDBBackend: udb}
	userRecord, err := users.Create(context.Background(), userCreateBody)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded user record")
	}
	tests.Passed("Should have successfully loaded user record")

	testUserLogin(t, userRecord, userCreateBody, tenantRecord, tf, ufsdb)

	os.RemoveAll("./keys")
}

func testUserLogin(t *testing.T, user pkg.User, usercreate pkg.CreateUser, tenant pkg.Tenant, tf tenancykit.UserSessionAPI, db types.UserSessionDBBackend) {
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

		currentUser, err := pkg.GetCurrentUser(logginUser)
		if err != nil {
			tests.FailedWithError(err, "Should have successfully retrieved current user")
		}
		tests.Passed("Should have successfully retrieved current user")

		if currentUser.TwoFactor == nil {
			tests.Failed("Should have succesfully valiated that current user has twofactor record")
		}
		tests.Passed("Should have succesfully valiated that current user has twofactor record")

	}
}
