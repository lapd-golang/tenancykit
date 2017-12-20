package baseapi_test

import (
	"testing"

	"github.com/gokit/tenancykit/baseapi"
	"github.com/gokit/tenancykit/db/types"
	"github.com/gokit/tenancykit/mock"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestTwoFactorRecordAPI(t *testing.T) {
	m := metrics.New()
	tfdb := mock.TFRecordBackend()
	tf := baseapi.NewTwoFactorAPI(m, tfdb)

	testTwoFactorRecordCreate(t, tf, tfdb)
	testTwoFactorRecordGetAll(t, tf, tfdb)
	testTwoFactorRecordGet(t, tf, tfdb)
	testTwoFactorRecordUpdate(t, tf, tfdb)
	testTwoFactorRecordDelete(t, tf, tfdb)
}

func testTwoFactorRecordCreate(t *testing.T, tf baseapi.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When creating a twofactor record using the TwoFactorRecordAPI")
	{
	}
}

func testTwoFactorRecordGetAll(t *testing.T, tf baseapi.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When retrieving all twofactor records using the TwoFactorRecordAPI")
	{
	}
}

func testTwoFactorRecordGet(t *testing.T, tf baseapi.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When retrieving a twofactor record using the TwoFactorRecordAPI")
	{
	}
}

func testTwoFactorRecordUpdate(t *testing.T, tf baseapi.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When updating a twofactor record using the TwoFactorRecordAPI")
	{
	}
}

func testTwoFactorRecordDelete(t *testing.T, tf baseapi.TwoFactorAPI, db types.TFRecordDBBackend) {
	tests.Header("When deleting a twofactor record using the TwoFactorRecordAPI")
	{
	}
}
