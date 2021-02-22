package service

import (
	"context"

	binance "github.com/adshao/go-binance/v2"
	"github.com/cryptellation/models.go"
)

// CandleStickServiceInterface is the interface for candle stick services
type CandleStickServiceInterface interface {
	Do(ctx context.Context) ([]models.CandleStick, error)
}

// CandleStickService is the real service for candlesticks
type CandleStickService struct {
	service *binance.KlinesService
}

// Do will execute a request for candlesticks
func (s *CandleStickService) Do(ctx context.Context) ([]models.CandleStick, error) {
	// TODO
	return nil, nil
}

// Mock section
////////////////////////////////////////////////////////////////////////////////

type testCandleSticks struct {
	Exchange     string
	Symbol       string
	Period       int64
	CandleSticks []models.CandleStick
}

// MockedCandleStickService is the mocked service for candlesticks
type MockedCandleStickService struct {
	TestCandleSticks []testCandleSticks
}

// Do will execute a request for candlesticks
func (m *MockedCandleStickService) Do(ctx context.Context) ([]models.CandleStick, error) {
	cs := make([]models.CandleStick, 0)
	for _, t := range m.TestCandleSticks {
		cs = append(cs, t.CandleSticks...)
	}
	return cs, nil
}
