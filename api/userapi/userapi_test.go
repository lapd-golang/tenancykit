package userapi_test

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

	httpapi "github.com/gokit/tenancykit/api/userapi"
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

// TestGetAllUser validates the retrieval of a User
// record from httpapi.
func TestGetAllUser(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	register(tree, api, version, "user")

	elem, err := loadJSONFor(userCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: userCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: userCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved User record : %+q.", err)
	}
	tests.Passed("Should have successfully saved User record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/user", version), nil)
	if err != nil {
		tests.Failed("Should have successfully created request User record : %+q.", err)
	}
	tests.Passed("Should have successfully created request User record.")

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

// TestGetUser validates the retrieval of all User
// record from a httpapi.
func TestGetUser(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "user")

	elem, err := loadJSONFor(userCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: userCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: userCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved User record : %+q.", err)
	}
	tests.Passed("Should have successfully saved User record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/user/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request User record : %+q.", err)
	}
	tests.Passed("Should have successfully created request User record.")

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

// TestUserCreate validates the creation of a User
// record with a httpapi.
func TestUserCreate(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "user")

	var body bytes.Buffer
	body.WriteString(userCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/user", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request User record : %+q.", err)
	}
	tests.Passed("Should have successfully created request User record.")

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

// TestUserUpdate validates the update of a User
// record with a httpapi.
func TestUserUpdate(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "user")

	var body bytes.Buffer
	body.WriteString(userCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/user", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request User record : %+q.", err)
	}
	tests.Passed("Should have successfully created request User record.")

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

	elem.Username = "Judith Dean I II III IV V MD DDS PhD DVM"

	var bu bytes.Buffer

	if err := json.NewEncoder(&bu).Encode(elem); err != nil {
		tests.Failed("Should have successfully encoded User:  %+q.", err)
	}
	tests.Passed("Should have successfully encoded User.")

	req, err = http.NewRequest("PUT", fmt.Sprintf("/%s/user/%s", version, elem.PublicID), &bu)
	if err != nil {
		tests.Failed("Should have successfully created request User record : %+q.", err)
	}
	tests.Passed("Should have successfully created request User record.")

	res = httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}

// TestUserDelete validates the removal of a User
// record from a httpapi.
func TestUserDelete(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "user")

	elem, err := loadJSONFor(userCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: userCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: userCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved User record : %+q.", err)
	}
	tests.Passed("Should have successfully saved User record.")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/user/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request User record : %+q.", err)
	}
	tests.Passed("Should have successfully created request User record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}
