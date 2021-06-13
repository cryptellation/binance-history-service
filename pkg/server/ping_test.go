package server

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	s, _ := newTestServer()

	req, _ := http.NewRequest("GET", "/ping", nil)

	response := s.executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var resp PingResponse
	json.Unmarshal(response.Body.Bytes(), &resp)

	if resp.Message != "pong" {
		t.Fatal("The ping response should be 'pong'")
	}
}
