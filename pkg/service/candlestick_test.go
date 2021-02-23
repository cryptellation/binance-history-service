package service

import (
	"context"
	"testing"
	"time"

	"github.com/cryptellation/models.go"
)

var tCS = []testCandleSticks{
	{
		"binance", "BTC-USDC", models.M1, []models.CandleStick{
			{Time: time.Time{}, Open: 10, High: 10, Low: 10, Close: 10},
			{Time: time.Time{}, Open: 15, High: 15, Low: 15, Close: 15}},
	},
	{
		"coinbase", "ETH-USDC", models.M5, []models.CandleStick{
			{Time: time.Time{}, Open: 20, High: 20, Low: 20, Close: 20},
			{Time: time.Time{}, Open: 25, High: 25, Low: 25, Close: 25}},
	},
}

func TestMockedDo(t *testing.T) {
	// Get new service
	s := MockedCandleStickService{tCS}

	// Do the service
	cs, _ := s.Do(context.Background())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks")
	}

	// Test first case
	for i, c := range tCS[0].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}

	// Test second case
	for i, c := range tCS[1].CandleSticks {
		if c != cs[i+2] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}
