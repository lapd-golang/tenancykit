package tenancykit_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/influx6/faux/httputil"

	"github.com/influx6/faux/httputil/httptesting"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	tenantFixtures "github.com/gokit/tenancykit/pkg/resources/tenantapi/fixtures"
	userFixtures "github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"
	tokenmock "github.com/gokit/tokens/db/mocks"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestTwoFactorAuth(t *testing.T) {
	m := metrics.New()
	tfdb := mock.TFRecordBackend()
	ttdb := mock.TenantDBBackend()
	tfsdb := mock.TFSessionBackend()
	tsetdb := tokenmock.TokenSetBackend()
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

	tfrecord := backends.TFBackend{TFRecordDBBackend: tfdb}
	if _, err := tfrecord.Create(context.Background(), pkg.NewTF{
		MaxLength: 6,
		Tenant:    tenantRecord,
		User:      userRecord,
		Domain:    "bob.com",
	}); err != nil {
		tests.FailedWithError(err, "Should have succesfully created user token record")
	}
	tests.Passed("Should have succesfully created user token record")

	tokenSessionAPI := tenancykit.NewTwoFactorSessionAPI(m, tsetdb, tfsdb)
	testUserLogin(t, userRecord, userCreateBody, tokenSessionAPI, tf, ufsdb)

	os.RemoveAll("./keys")
}

func testUserLogin(t *testing.T, user pkg.User, usercreate pkg.CreateUser, tsession tenancykit.TwoFactorSessionAPI, tf tenancykit.UserSessionAPI, db types.UserSessionDBBackend) {
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

		userCode, err := currentUser.TwoFactor.OTP()
		if err != nil {
			tests.FailedWithError(err, "Should have generated new user code")
		}
		tests.Passed("Should have generated new user code")

		tests.Info("Using User Code: %+q", userCode)

		tokenValidateResponse := httptest.NewRecorder()
		tokenValidReq := httptesting.Post("/tokens/login?user_code="+userCode, nil, tokenValidateResponse)
		httputil.SetValueBag(logginUser.Bag())(tokenValidReq)
		tokenValidReq.Set("user_code", userCode)

		if err := tsession.ValidateUserToken(tokenValidReq); err != nil {
			tests.FailedWithError(err, "Should have successfully validated user token")
		}
		tests.Passed("Should have successfully validated user token")

		if tokenValidateResponse.Code != http.StatusNoContent {
			tests.Failed("Should have received Status 204")
		}
		tests.Passed("Should have received Status 204")

		ncurrentUser, err := pkg.GetCurrentUser(tokenValidReq)
		if err != nil {
			tests.FailedWithError(err, "Should have successfully retrieved current user")
		}
		tests.Passed("Should have successfully retrieved current user")

		if ncurrentUser.TFSession != nil {
			tests.Failed("Should not have a twofactor sesion attached")
		}
		tests.Passed("Should not have a twofactor sesion attached")

		time.Sleep(30 * time.Second)

		nuserCode, err := currentUser.TwoFactor.OTP()
		if err != nil {
			tests.FailedWithError(err, "Should have generated new user code")
		}
		tests.Passed("Should have generated new user code")

		if nuserCode == userCode {
			nuserCode, err = currentUser.TwoFactor.OTP()
			fmt.Printf("Regening: %+q -> %+q\n", nuserCode, err)
		}

		tests.Info("Using User Code: %+q", nuserCode)

		ntokenValidateResponse := httptest.NewRecorder()
		ntokenValidReq := httptesting.Post("/tokens/login?user_code="+nuserCode, nil, ntokenValidateResponse)
		httputil.SetValueBag(logginUser.Bag())(ntokenValidReq)
		ntokenValidReq.Set("user_code", nuserCode)

		if err := tsession.NewSession(ntokenValidReq); err != nil {
			tests.FailedWithError(err, "Should have successfully create token session validation")
		}
		tests.Passed("Should have successfully create token session validation")

		if ntokenValidateResponse.Code != http.StatusNoContent {
			tests.Failed("Should have received Status 204")
		}
		tests.Passed("Should have received Status 204")

		nwcurrentUser, err := pkg.GetCurrentUser(ntokenValidReq)
		if err != nil {
			tests.FailedWithError(err, "Should have successfully retrieved current user")
		}
		tests.Passed("Should have successfully retrieved current user")

		if nwcurrentUser.TFSession == nil {
			tests.Failed("Should have a twofactor sesion attached")
		}
		tests.Passed("Should have a twofactor sesion attached")
	}
}
