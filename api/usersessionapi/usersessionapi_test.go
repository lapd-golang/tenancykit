package usersessionapi_test

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

	httpapi "github.com/gokit/tenancykit/api/usersessionapi"
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

// TestGetAllUserSession validates the retrieval of a UserSession
// record from httpapi.
func TestGetAllUserSession(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	register(tree, api, version, "usersession")

	elem, err := loadJSONFor(usersessionCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: usersessionCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: usersessionCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully saved UserSession record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/usersession", version), nil)
	if err != nil {
		tests.Failed("Should have successfully created request UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request UserSession record.")

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

// TestGetUserSession validates the retrieval of all UserSession
// record from a httpapi.
func TestGetUserSession(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "usersession")

	elem, err := loadJSONFor(usersessionCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: usersessionCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: usersessionCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully saved UserSession record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/usersession/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request UserSession record.")

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

// TestUserSessionCreate validates the creation of a UserSession
// record with a httpapi.
func TestUserSessionCreate(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "usersession")

	var body bytes.Buffer
	body.WriteString(usersessionCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/usersession", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request UserSession record.")

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

// TestUserSessionUpdate validates the update of a UserSession
// record with a httpapi.
func TestUserSessionUpdate(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "usersession")

	var body bytes.Buffer
	body.WriteString(usersessionCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/usersession", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request UserSession record.")

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

	elem.UserID = "Tammy Elliott"

	var bu bytes.Buffer

	if err := json.NewEncoder(&bu).Encode(elem); err != nil {
		tests.Failed("Should have successfully encoded UserSession:  %+q.", err)
	}
	tests.Passed("Should have successfully encoded UserSession.")

	req, err = http.NewRequest("PUT", fmt.Sprintf("/%s/usersession/%s", version, elem.PublicID), &bu)
	if err != nil {
		tests.Failed("Should have successfully created request UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request UserSession record.")

	res = httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}

// TestUserSessionDelete validates the removal of a UserSession
// record from a httpapi.
func TestUserSessionDelete(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "usersession")

	elem, err := loadJSONFor(usersessionCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: usersessionCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: usersessionCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully saved UserSession record.")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/usersession/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request UserSession record : %+q.", err)
	}
	tests.Passed("Should have successfully created request UserSession record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}
