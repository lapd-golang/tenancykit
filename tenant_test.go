package tenancykit_test

import (
	"testing"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/mock"
	"github.com/influx6/faux/metrics"
	"github.com/influx6/faux/tests"
)

func TestTenantAPI(t *testing.T) {
	m := metrics.New()
	tenantdb := mock.TenantDBBackend()
	tenants := tenancykit.NewTenantAPI(m, tenantdb)

	testTenantCreate(t, tenants, tenantdb)
	testTenantGetAll(t, tenants, tenantdb)
	testTenantGet(t, tenants, tenantdb)
	testTenantUpdate(t, tenants, tenantdb)
	testTenantDelete(t, tenants, tenantdb)
}

func testTenantCreate(t *testing.T, tenants tenancykit.TenantAPI, db types.TenantDBBackend) {
	tests.Header("When creating a tenant using the TenantAPI")
	{
	}
}

func testTenantGetAll(t *testing.T, tenants tenancykit.TenantAPI, db types.TenantDBBackend) {
	tests.Header("When retrieving all tenants using the TenantAPI")
	{
	}
}

func testTenantGet(t *testing.T, tenants tenancykit.TenantAPI, db types.TenantDBBackend) {
	tests.Header("When retrieving a tenant using the TenantAPI")
	{
	}
}

func testTenantUpdate(t *testing.T, tenants tenancykit.TenantAPI, db types.TenantDBBackend) {
	tests.Header("When updating a tenant using the TenantAPI")
	{
	}
}

func testTenantDelete(t *testing.T, tenants tenancykit.TenantAPI, db types.TenantDBBackend) {
	tests.Header("When deleting a tenant using the TenantAPI")
	{
	}
}
