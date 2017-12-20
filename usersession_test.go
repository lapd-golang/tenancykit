package tenancykit_test

import (
	"testing"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
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

	tf := tenancykit.NewUserSessionAPI(m, udb, ufsdb, tfsdb, ttdb, tfdb)

	testUserSessionCreate(t, tf, ufsdb)
	testUserSessionGetAll(t, tf, ufsdb)
	testUserSessionGet(t, tf, ufsdb)
	testUserSessionUpdate(t, tf, ufsdb)
	testUserSessionDelete(t, tf, ufsdb)
}

func testUserSessionCreate(t *testing.T, tf tenancykit.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When creating a two factor record using the UserSessionAPI")
	{
	}
}

func testUserSessionGetAll(t *testing.T, tf tenancykit.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When retrieving all two factor records using the UserSessionAPI")
	{
	}
}

func testUserSessionGet(t *testing.T, tf tenancykit.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When retrieving a two factor record using the UserSessionAPI")
	{
	}
}

func testUserSessionUpdate(t *testing.T, tf tenancykit.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When updating a two factor record using the UserSessionAPI")
	{
	}
}

func testUserSessionDelete(t *testing.T, tf tenancykit.UserSessionAPI, db types.UserSessionDBBackend) {
	tests.Header("When deleting a two factor record using the UserSessionAPI")
	{
	}
}
