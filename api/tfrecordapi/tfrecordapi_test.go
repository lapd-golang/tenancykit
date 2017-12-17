package tfrecordapi_test

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

	httpapi "github.com/gokit/tenancykit/api/tfrecordapi"
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

// TestGetAllTFRecord validates the retrieval of a TFRecord
// record from httpapi.
func TestGetAllTFRecord(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	register(tree, api, version, "tfrecord")

	elem, err := loadJSONFor(tfrecordCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tfrecordCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tfrecordCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully saved TFRecord record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/tfrecord", version), nil)
	if err != nil {
		tests.Failed("Should have successfully created request TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TFRecord record.")

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

// TestGetTFRecord validates the retrieval of all TFRecord
// record from a httpapi.
func TestGetTFRecord(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "tfrecord")

	elem, err := loadJSONFor(tfrecordCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tfrecordCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tfrecordCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully saved TFRecord record.")

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/tfrecord/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TFRecord record.")

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

// TestTFRecordCreate validates the creation of a TFRecord
// record with a httpapi.
func TestTFRecordCreate(t *testing.T) {

	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "tfrecord")

	var body bytes.Buffer
	body.WriteString(tfrecordCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/tfrecord", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TFRecord record.")

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

// TestTFRecordUpdate validates the update of a TFRecord
// record with a httpapi.
func TestTFRecordUpdate(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "tfrecord")

	var body bytes.Buffer
	body.WriteString(tfrecordCreateJSON)

	req, err := http.NewRequest("POST", fmt.Sprintf("/%s/tfrecord", version), &body)
	if err != nil {
		tests.Failed("Should have successfully created request TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TFRecord record.")

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

	elem.Key = "Mr. Dr. Daniel Morrison"

	var bu bytes.Buffer

	if err := json.NewEncoder(&bu).Encode(elem); err != nil {
		tests.Failed("Should have successfully encoded TFRecord:  %+q.", err)
	}
	tests.Passed("Should have successfully encoded TFRecord.")

	req, err = http.NewRequest("PUT", fmt.Sprintf("/%s/tfrecord/%s", version, elem.PublicID), &bu)
	if err != nil {
		tests.Failed("Should have successfully created request TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TFRecord record.")

	res = httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}

// TestTFRecordDelete validates the removal of a TFRecord
// record from a httpapi.
func TestTFRecordDelete(t *testing.T) {
	db := NewMockAPIOperator()

	api := httpapi.New(events, db)
	tree := httputil.NewMuxRouter(nil)

	// Register routes with router group.
	registerRoute(tree, api, version, "tfrecord")

	elem, err := loadJSONFor(tfrecordCreateJSON)
	if err != nil {
		tests.Failed("Should have successfully loaded JSON: tfrecordCreateJSON : %+q.", err)
	}
	tests.Passed("Should have successfully loaded JSON: tfrecordCreateJSON.")

	ctx := context.New()

	elem, err = db.Create(ctx, elem)
	if err != nil {
		tests.Failed("Should have successfully saved TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully saved TFRecord record.")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/tfrecord/%s", version, elem.PublicID), nil)
	if err != nil {
		tests.Failed("Should have successfully created request TFRecord record : %+q.", err)
	}
	tests.Passed("Should have successfully created request TFRecord record.")

	res := httptest.NewRecorder()

	tree.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		tests.Failed("Should have received status code %d from response but got %d.", http.StatusNoContent, res.Code)
	}
	tests.Passed("Should have received status code %d from response.", http.StatusNoContent)
}
