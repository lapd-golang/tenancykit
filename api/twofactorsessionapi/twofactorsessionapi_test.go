package twofactorsessionapi_test

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

	httpapi "github.com/gokit/tenancykit/api/twofactorsessionapi"
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

// TestGetAllTwoFactorSession validates the retrieval of a TwoFactorSession
// record from httpapi.
func TestGetAllTwoFactorSession(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	register(tree, api, version, "twofactorsession")

	elem, err := loadJSONFor(twofactorsessionCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: twofactorsessionCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: twofactorsessionCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully saved TwoFactorSession record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/twofactorsession", version), nil)
	if err != nil {
		tests.Failed("Should have successfully created request TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TwoFactorSession record.")

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

// TestGetTwoFactorSession validates the retrieval of all TwoFactorSession
// record from a httpapi.
func TestGetTwoFactorSession(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "twofactorsession")

	elem, err := loadJSONFor(twofactorsessionCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: twofactorsessionCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: twofactorsessionCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully saved TwoFactorSession record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/twofactorsession/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TwoFactorSession record.")

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

// TestTwoFactorSessionCreate validates the creation of a TwoFactorSession
// record with a httpapi.
func TestTwoFactorSessionCreate(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "twofactorsession")

	var body bytes.Buffer
	body.WriteString(twofactorsessionCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/twofactorsession", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TwoFactorSession record.")

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

// TestTwoFactorSessionUpdate validates the update of a TwoFactorSession
// record with a httpapi.
func TestTwoFactorSessionUpdate(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "twofactorsession")

	var body bytes.Buffer
	body.WriteString(twofactorsessionCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/twofactorsession", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TwoFactorSession record.")

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

	elem.UserID = "Wanda Tucker"

	var bu bytes.Buffer

	if err := json.NewEncoder(&bu).Encode(elem); err != nil {
		tests.Failed("Should have successfully encoded TwoFactorSession:  %+q.", err)
	}
	tests.Passed("Should have successfully encoded TwoFactorSession.")

	req, err = http.NewRequest("PUT", fmt.Sprintf("/%s/twofactorsession/%s", version, elem.PublicID), &bu)
	if err != nil {
		tests.Failed("Should have successfully created request TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TwoFactorSession record.")

	res = httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}

// TestTwoFactorSessionDelete validates the removal of a TwoFactorSession
// record from a httpapi.
func TestTwoFactorSessionDelete(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "twofactorsession")

	elem, err := loadJSONFor(twofactorsessionCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: twofactorsessionCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: twofactorsessionCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully saved TwoFactorSession record.")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/twofactorsession/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request TwoFactorSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TwoFactorSession record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}
