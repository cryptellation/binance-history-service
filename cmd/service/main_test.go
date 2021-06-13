package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/cryptellation/binance-history-service/pkg/server"
)

func TestServe(t *testing.T) {
	go run("../../configs/binance-history-service.example.toml")
	time.Sleep(1 * time.Second)

	httpResponse, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/ping"))
	if err != nil {
		t.Fatal("There should be no error but there is", err)
	}

	var resp server.PingResponse
	body, err := ioutil.ReadAll(httpResponse.Body)
	json.Unmarshal(body, &resp)

	if resp.Message != "pong" {
		t.Fatal("The ping response should be 'pong'")
	}
}

func TestServe_IncorrectPath(t *testing.T) {
	err := run("/do/not/exist.toml")
	if err == nil {
		t.Fatal("The config should not exist")
	}
}
