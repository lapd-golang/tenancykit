package tokenapi_test

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

	httpapi "github.com/gokit/tenancykit/tokenset/tokens/tokenapi"
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

// TestGetAllToken validates the retrieval of a Token
// record from httpapi.
func TestGetAllToken(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	register(tree, api, version, "token")

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tokenCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tokenCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved Token record : %+q.", err)
	}
	tests.Passed("Should have successfully saved Token record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/token", version), nil)
	if err != nil {
		tests.Failed("Should have successfully created request Token record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Token record.")

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

// TestGetToken validates the retrieval of all Token
// record from a httpapi.
func TestGetToken(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "token")

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tokenCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tokenCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved Token record : %+q.", err)
	}
	tests.Passed("Should have successfully saved Token record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/token/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request Token record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Token record.")

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

// TestTokenCreate validates the creation of a Token
// record with a httpapi.
func TestTokenCreate(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "token")

	var body bytes.Buffer
	body.WriteString(tokenCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/token", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request Token record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Token record.")

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

// TestTokenUpdate validates the update of a Token
// record with a httpapi.
func TestTokenUpdate(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "token")

	var body bytes.Buffer
	body.WriteString(tokenCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/token", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request Token record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Token record.")

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

	elem.Value = "Mrs. Ms. Miss Marie Warren"

	var bu bytes.Buffer

	if err := json.NewEncoder(&bu).Encode(elem); err != nil {
		tests.Failed("Should have successfully encoded Token:  %+q.", err)
	}
	tests.Passed("Should have successfully encoded Token.")

	req, err = http.NewRequest("PUT", fmt.Sprintf("/%s/token/%s", version, elem.PublicID), &bu)
	if err != nil {
		tests.Failed("Should have successfully created request Token record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Token record.")

	res = httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}

// TestTokenDelete validates the removal of a Token
// record from a httpapi.
func TestTokenDelete(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "token")

	elem, err := loadJSONFor(tokenCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tokenCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tokenCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved Token record : %+q.", err)
	}
	tests.Passed("Should have successfully saved Token record.")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/token/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request Token record : %+q.", err)
	}
	tests.Passed("Should have successfully created request Token record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}
