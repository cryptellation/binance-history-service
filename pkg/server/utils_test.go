package server

import (
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestHTTPResponse_UncorrectJSON(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		httpResponse(w, http.StatusOK, math.Inf(1))
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	checkResponseCode(t, http.StatusInternalServerError, resp.StatusCode)
}
