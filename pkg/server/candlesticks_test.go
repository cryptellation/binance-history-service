package server

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cryptellation/binance.go/mock"
	"github.com/cryptellation/models.go"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestCandles(t *testing.T) {
	// Create a new server
	s := newTestServer()

	// Build a request
	req, _ := http.NewRequest("GET", "/candlesticks", nil)
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	response := s.executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	// Analyse response
	var cs []models.CandleStick
	json.Unmarshal(response.Body.Bytes(), &cs)

	// Check length
	if len(cs) != mock.TestCandleSticksCount() {
		t.Fatal("There should be", mock.TestCandleSticksCount(), "candlestick, but there is", len(cs))
	}
}
