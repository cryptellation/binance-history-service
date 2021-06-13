package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/cryptellation/binance.go/pkg/mock"
)

func TestCandles(t *testing.T) {
	s, _ := newTestServer()

	req, _ := http.NewRequest("GET", "/candlesticks", nil)
	req.Header.Set("Content-Type", "application/json")

	response := s.executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var resp CandleSticksResponse
	json.Unmarshal(response.Body.Bytes(), &resp)

	if len(resp.CandleSticks) != mock.TestCandleSticksCount() {
		t.Fatal("There should be", mock.TestCandleSticksCount(),
			"candlestick, but there is", len(resp.CandleSticks))
	}
}

func TestCandles_Error(t *testing.T) {
	s, ms := newTestServer()
	ms.NextError(errors.New("NextError"))

	req, _ := http.NewRequest("GET", "/candlesticks", nil)
	req.Header.Set("Content-Type", "application/json")

	response := s.executeRequest(req)
	checkResponseCode(t, http.StatusInternalServerError, response.Code)
}
