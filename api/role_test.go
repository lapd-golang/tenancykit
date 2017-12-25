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
	testRoleCount(t, userRecord, tf, tfdb)
	testRoleGetAll(t, userRecord, tf, tfdb)
	testRoleGet(t, userRecord, tf, tfdb)
	testRoleUpdate(t, userRecord, tf, tfdb)
	testRoleDelete(t, userRecord, tf, tfdb)
}

func testRoleUserRoles(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend) {}
func testRoleCreate(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend)    {}
func testRoleCount(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend)     {}
func testRoleGetAll(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend)    {}
func testRoleGet(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend)       {}
func testRoleUpdate(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend)    {}
func testRoleDelete(t *testing.T, user pkg.User, roles api.RoleAPI, db types.RoleDBBackend)    {}
