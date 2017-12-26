package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gokit/tenancykit/pkg/resources/roleapi"

	"github.com/gokit/tenancykit/pkg/resources/roleapi/fixtures"
	"github.com/influx6/faux/httputil/httptesting"

	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	userFixtures "github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestRoleAPI(t *testing.T) {
	m := metrics.New()
	tfdb := mock.RoleBackend()
	userdb := mock.UserBackend()
	tf := api.NewRoleAPI(m, tfdb)

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

	testRoleCreate(t, userRecord, tf, tfdb)
	testRoleUserRoles(t, userRecord, tf, tfdb)
}

func testRoleUserRoles(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend) {
	tests.Header("When we create roles for a user with RoleAPI")
	{
		recs, total, err := db.GetAll(context.Background(), "", "", 0, 0)
		if err != nil {
			tests.FailedWithError(err, "Should have retrieved all roles in db")
		}
		tests.FailedWithError(err, "Should have retrieved all roles in db")

		if total == 0 {
			tests.Failed("Should have existing records in db")
		}
		tests.Passed("Should have existing records in db")

		roleRecord := recs[0]
		user.Roles = append(user.Roles, roleRecord.PublicID)

		infoResponse := httptest.NewRecorder()
		infoResource := httptesting.NewRequest("INFO", "/roles"+user.PublicID, nil, infoResponse)
		infoResource.Set(pkg.ContextKeyUser, user)
		if err := roles.GetUserRoles(infoResource); err != nil {
			tests.FailedWithError(err, "Should have successfully retrieved user roles")
		}
		tests.Passed("Should have successfully retrieved user roles")

		userRoles, err := pkg.GetUserRoles(infoResource)
		if err != nil {
			tests.FailedWithError(err, "Should have user roles set in context")
		}
		tests.Passed("Should have user roles set in context")

		if len(userRoles) != len(user.Roles) {
			tests.Failed("Should have length match for user roles and provided roles")
		}
		tests.Passed("Should have length match for user roles and provided roles")
	}
}

func testRoleInfo(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend) {
	tests.Header("When we count all Roles with RoleAPI")
	{
		infoResponse := httptest.NewRecorder()
		infoResource := httptesting.NewRequest("INFO", "/roles", nil, infoResponse)
		if err := roles.Info(infoResource); err != nil {
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

		var info roleapi.RoleInfo
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

func testRoleCreate(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend) {
	tests.Header("When we create all Roles with RoleAPI")
	{
		createResponse := httptest.NewRecorder()
		createResource := httptesting.Post("/roles", bytes.NewBufferString(fixtures.RoleCreateJSON), createResponse)
		if err := roles.Create(createResource); err != nil {
			tests.Info("JSON: %+s", fixtures.RoleCreateJSON)
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

		if _, err := fixtures.LoadRoleJSON(createResponse.Body.String()); err != nil {
			tests.FailedWithError(err, "Should have successfully received new record response")
		}
		tests.Passed("Should have successfully received new record response")
	}
}
