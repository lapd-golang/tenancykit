package tenantapi_test

import (
	"fmt"
	"os"

	"bytes"

	"testing"

	"encoding/json"

	"net/http"

	"net/http/httptest"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/context"

	"github.com/influx6/faux/metrics/custom"

	httpapi "github.com/gokit/tenancykit/api/tenantapi"
)

var (
	events  = metrics.New(custom.StackDisplay(os.Stdout))
	version = "v1"
)

func registerRoute(router httputil.Router, api *httpapi.HTTPAPI, version string, resource string) {
	router.Handle("GET", fmt.Sprintf("/%s/%s", version, resource), api.GetAll)
	router.Handle("GET", fmt.Sprintf("/%s/%s/:public_id", version, resource), api.Get)

	router.Handle("POST", fmt.Sprintf("/%s/%s", version, resource), api.Create)

	router.Handle("PUT", fmt.Sprintf("/%s/%s/:public_id", version, resource), api.Update)
	router.Handle("DELETE", fmt.Sprintf("/%s/%s/:public_id", version, resource), api.Delete)
}

// TestGetAllTenant validates the retrieval of a Tenant
// record from httpapi.
func TestGetAllTenant(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	register(tree, api, version, "tenant")

	elem, err := loadJSONFor(tenantCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tenantCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tenantCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully saved Tenant record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/tenant", version), nil)
	if err != nil {
		tests.Failed("Should have successfully created request Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Tenant record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusOK, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusOK)

	if res.Body.Len() == 0 {
		tests.Failed("Should have successfully received response body.")
	}
	tests.Passed("Should have successfully received response body.")
}

// TestGetTenant validates the retrieval of all Tenant
// record from a httpapi.
func TestGetTenant(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "tenant")

	elem, err := loadJSONFor(tenantCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tenantCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tenantCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully saved Tenant record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/tenant/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Tenant record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusOK, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusOK)

	if res.Body.Len() == 0 {
		tests.Failed("Should have successfully received response body.")
	}
	tests.Passed("Should have successfully received response body.")
}

// TestTenantCreate validates the creation of a Tenant
// record with a httpapi.
func TestTenantCreate(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "tenant")

	var body bytes.Buffer
	body.WriteString(tenantCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/tenant", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Tenant record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusCreated, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusCreated)

	if res.Body.Len() == 0 {
		tests.Failed("Should have successfully received response body.")
	}
	tests.Passed("Should have successfully received response body.")
}

// TestTenantUpdate validates the update of a Tenant
// record with a httpapi.
func TestTenantUpdate(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "tenant")

	var body bytes.Buffer
	body.WriteString(tenantCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/tenant", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Tenant record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusCreated, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusCreated)

	if res.Body.Len() == 0 {
		tests.Failed("Should have successfully received response body.")
	}
	tests.Passed("Should have successfully received response body.")

	elem, err := loadJSONFor(res.Body.String())
	if err != nil {
		tests.Failed("Should have successfully loaded JSON:  %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON.")

	elem.Name = "John Young"

	var bu bytes.Buffer

	if err := json.NewEncoder(&bu).Encode(elem); err != nil {
		tests.Failed("Should have successfully encoded Tenant:  %+q.", err)
	}
	tests.Passed("Should have successfully encoded Tenant.")

	req, err = http.NewRequest("PUT", fmt.Sprintf("/%s/tenant/%s", version, elem.PublicID), &bu)
	if err != nil {
		tests.Failed("Should have successfully created request Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Tenant record.")

	res = httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}

// TestTenantDelete validates the removal of a Tenant
// record from a httpapi.
func TestTenantDelete(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "tenant")

	elem, err := loadJSONFor(tenantCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tenantCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tenantCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully saved Tenant record.")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/tenant/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request Tenant record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Tenant record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}
