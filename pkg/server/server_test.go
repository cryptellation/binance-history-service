package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cryptellation/binance.go/mock"
)

func newTestServer() *Server {
	// Create a mock service
	ms := mock.New()
	ms.AddCandleSticks(mock.TestCandleSticks)

	// Return server
	return New(ms)
}

func (s *Server) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)
	return rr
}

func TestNew(t *testing.T) {
	// Create server
	s := New(nil)

	// Check if the server is not null
	if s == nil {
		t.Fatal("New server is nil")
	}

	// Check if version is correct
	if s.Version() != version {
		t.Error("Version does not match")
	}
}
