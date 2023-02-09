package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"pulley.com/shakesearch/pkg"
	"reflect"
	"testing"
)

// TestVersion tests the Version handler
// this test checks for:
// - the Content-Type header is set to application/json
// - the status code is 200 OK
// - the response body is a JSON object with a key "version" and a value equal to the API version
func TestVersion(t *testing.T) {

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	Version(w, r)

	got := w.Result()

	if got.Header.Get("Content-Type") != pkg.JSONContentType {
		t.Error("Content-Type header is not set to application/json")
	}

	if got.StatusCode != http.StatusOK {
		t.Error("Status code is not 200 OK")
	}

	var expected = map[string]string{}
	expected["version"] = pkg.APIVersion
	body, err := io.ReadAll(got.Body)
	if err != nil {
		t.Error(err)
	}

	jsonContent, err := json.Marshal(expected)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(body, jsonContent) {
		t.Errorf("handler returned unexpected body: got %s want %s", string(body), string(jsonContent))
	}
}
