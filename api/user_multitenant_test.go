package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"

	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg"
	"github.com/influx6/faux/httputil/httptesting"
	"github.com/influx6/faux/tests"

	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	tenantFixtures "github.com/gokit/tenancykit/pkg/resources/tenantapi/fixtures"
	"github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"
	"github.com/influx6/faux/metrics"
)

func TestMultiTenantUserAPI(t *testing.T) {
	m := metrics.New()
	userdb := mock.UserBackend()
	tenantdb := mock.TenantDBBackend()

	users := api.NewMultiTenantUserAPI(
		m,
		userdb,
		tenantdb,
	)

	createTenant, err := tenantFixtures.LoadCreateJSON(tenantFixtures.TenantCreateJSON)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded tenant record")
	}
	tests.Passed("Should have successfully loaded tenant record")

	tenants := api.TenantBackend{TenantDBBackend: tenantdb}
	tenantRecord, err := tenants.Create(context.Background(), createTenant)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created tenant")
	}
	tests.Passed("Should have successfully created tenants")

	testMultiTenantUserCreate(t, tenantRecord, users, userdb)
}

func testMultiTenantUserCreate(t *testing.T, tenant pkg.Tenant, users api.UserAPI, db types.UserDBBackend) {
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
