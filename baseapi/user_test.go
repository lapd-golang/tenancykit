package baseapi_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gokit/tenancykit/db/types"

	"github.com/influx6/faux/context"
	"github.com/influx6/faux/httputil/httptesting"
	"github.com/influx6/faux/tests"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/api/userapi/fixtures"
	"github.com/gokit/tenancykit/baseapi"
	"github.com/gokit/tenancykit/mock"
	"github.com/influx6/faux/metrics"
)

func TestUserAPI(t *testing.T) {
	m := metrics.New()
	userdb := mock.UserBackend()
	tenantdb := mock.TenantDBBackend()
	tfdb := mock.TFRecordBackend()

	users := baseapi.NewUserAPI(
		tenancykit.GoogleAuthenticatorUserCodeLength,
		m,
		userdb,
		tenantdb,
		tfdb,
	)

	// We are testing the following behaviour
	//Delete(context.Context, string) error
	//Get(context.Context, string) (tenancykit.User, error)
	//Update(context.Context, string, tenancykit.UpdateUser) error
	//GetAll(context.Context, string, string, int, int) ([]tenancykit.User, int, error)
	//Create(context.Context, tenancykit.CreateUser) (tenancykit.User, error)

	testUserCreate(t, users, userdb)
}

func testUserCreate(t *testing.T, users baseapi.UserAPI, db types.UserDBBackend) {
	tests.Header("When creating user with user api")
	{
		createBody := bytes.NewBufferString(fixtures.UserCreateJSON)
		createResponse := httptest.NewRecorder()
		createUser := httptesting.Post("/users", createBody, createResponse)
		if err := users.Create(createUser); err != nil {
			tests.FailedWithError(err, "Should have successfully created user")
		}
		tests.Passed("Should have successfully created user")

		if createResponse.Code != http.StatusCreated {
			tests.Failed("Should have received Status 204")
		}
		tests.Passed("Should have received Status 204")

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

func testUserGet(t *testing.T, users baseapi.UserAPI, db types.UserDBBackend) {
	tests.Header("When creating user with user api")
	{
		records, total, err := db.GetAll(context.New(), "", "", 0, 0)
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
		if err := users.Create(getUser); err != nil {
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
