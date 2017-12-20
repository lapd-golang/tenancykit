package tenancykit_test

import (
	"testing"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/pkg/backends"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestTwoFactorSessionAPI(t *testing.T) {
	m := metrics.New()
	tfdb := mock.TFSessionBackend()
	tfset := backends.TokenBackend{TokenDBBackend: mock.TokenSetBackend()}
	tf := tenancykit.NewTwoFactorSessionAPI(m, tfset, tfdb)

	testTwoFactorSessionCreate(t, tf, tfdb)
	testTwoFactorSessionGetAll(t, tf, tfdb)
	testTwoFactorSessionGet(t, tf, tfdb)
	testTwoFactorSessionUpdate(t, tf, tfdb)
	testTwoFactorSessionDelete(t, tf, tfdb)
}

func testTwoFactorSessionCreate(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When creating a twofactor session record using the TwoFactorSessionAPI")
	{
	}
}

func testTwoFactorSessionGetAll(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When getting all twofactor session record using the TwoFactorSessionAPI")
	{
	}
}

func testTwoFactorSessionGet(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When retrieving a twofactor session record using the TwoFactorSessionAPI")
	{
	}
}

func testTwoFactorSessionUpdate(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When updating a twofactor session record using the TwoFactorSessionAPI")
	{
	}
}

func testTwoFactorSessionDelete(t *testing.T, tf tenancykit.TwoFactorSessionAPI, db types.TwoFactorSessionDBBackend) {
	tests.Header("When deleting a twofactor session record using the TwoFactorSessionAPI")
	{
	}
}
