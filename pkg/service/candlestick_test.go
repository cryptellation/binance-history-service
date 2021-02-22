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
			{time.Time{}, 10, 10, 10, 10},
			{time.Time{}, 10, 10, 10, 10}},
	},
	{
		"coinbase", "ETH-USDC", models.M5, []models.CandleStick{
			{time.Time{}, 20, 20, 20, 20},
			{time.Time{}, 20, 20, 20, 20}},
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
}
