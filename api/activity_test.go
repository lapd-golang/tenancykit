package api_test

import (
	"context"
	"testing"

	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	userFixtures "github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestActivityAPI(t *testing.T) {
	m := metrics.New()
	tfdb := mock.ActivityBackend()
	userdb := mock.UserBackend()
	tf := api.NewActivityAPI(m, tfdb)

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

	testActivityCreate(t, userRecord, tf, tfdb)
	testActivityUserRoles(t, userRecord, tf, tfdb)
	testActivityCount(t, userRecord, tf, tfdb)
	testActivityGetAll(t, userRecord, tf, tfdb)
	testActivityGet(t, userRecord, tf, tfdb)
	testActivityUpdate(t, userRecord, tf, tfdb)
	testActivityDelete(t, userRecord, tf, tfdb)
}

func testActivityUserRoles(t *testing.T, user pkg.User, roles api.ActivityAPI, db types.ActivityDBBackend) {
}
func testActivityCreate(t *testing.T, user pkg.User, roles api.ActivityAPI, db types.ActivityDBBackend) {
}
func testActivityCount(t *testing.T, user pkg.User, roles api.ActivityAPI, db types.ActivityDBBackend) {
}
func testActivityGetAll(t *testing.T, user pkg.User, roles api.ActivityAPI, db types.ActivityDBBackend) {
}
func testActivityGet(t *testing.T, user pkg.User, roles api.ActivityAPI, db types.ActivityDBBackend) {}
func testActivityUpdate(t *testing.T, user pkg.User, roles api.ActivityAPI, db types.ActivityDBBackend) {
}
func testActivityDelete(t *testing.T, user pkg.User, roles api.ActivityAPI, db types.ActivityDBBackend) {
}
