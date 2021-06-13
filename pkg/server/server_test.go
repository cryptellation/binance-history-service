package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cryptellation/binance.go/pkg/mock"
)

func newTestServer() (*Server, *mock.MockedService) {
	ms := mock.New()
	ms.AddCandleSticks(mock.TestCandleSticks)

	return New(ms), ms
}

func (s *Server) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)
	return rr
}

func TestNew(t *testing.T) {
	s := New(nil)

	if s == nil {
		t.Fatal("New server is nil")
	}

	if s.Version() != version {
		t.Error("Version does not match")
	}
}

func TestServe(t *testing.T) {
	s := New(nil)
	addr := "127.0.0.1:5757"

	go s.Serve(addr)
	time.Sleep(1 * time.Second)

	httpResponse, err := http.Get(fmt.Sprintf("http://%s/ping", addr))
	if err != nil {
		t.Fatal("There should be no error but there is", err)
	}

	var resp PingResponse
	body, err := ioutil.ReadAll(httpResponse.Body)
	json.Unmarshal(body, &resp)

	if resp.Message != "pong" {
		t.Fatal("The ping response should be 'pong'")
	}
}
