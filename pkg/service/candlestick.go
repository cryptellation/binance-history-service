package service

import (
	"context"

	binance "github.com/adshao/go-binance/v2"
	"github.com/cryptellation/binance-messenger-service/pkg/service/adapters"
	"github.com/cryptellation/models.go"
)

// CandleStickServiceInterface is the interface for candle stick services
type CandleStickServiceInterface interface {
	Do(ctx context.Context) ([]models.CandleStick, error)
	Symbol(symbol string) CandleStickServiceInterface
}

// CandleStickService is the real service for candlesticks
type CandleStickService struct {
	service *binance.KlinesService
}

// Do will execute a request for candlesticks
func (s *CandleStickService) Do(ctx context.Context) ([]models.CandleStick, error) {
	// Get KLines
	kl, err := s.service.Do(ctx)
	if err != nil {
		return nil, err
	}

	// Change them to right format
	return adapters.KLinesToCandleSticks(kl)
}

// Symbol will specify a symbol for candlesticks next request
func (s *CandleStickService) Symbol(symbol string) CandleStickServiceInterface {
	s.service.Symbol(symbol)
	return s
}

// Mock section
////////////////////////////////////////////////////////////////////////////////

type testCandleSticks struct {
	Symbol       string
	Period       int64
	CandleSticks []models.CandleStick
}

// MockedCandleStickService is the mocked service for candlesticks
type MockedCandleStickService struct {
	TestCandleSticks []testCandleSticks

	// Next request specifications
	symbol string
}

// Do will execute a request for candlesticks
func (m *MockedCandleStickService) Do(ctx context.Context) ([]models.CandleStick, error) {
	cs := make([]models.CandleStick, 0)
	for _, t := range m.TestCandleSticks {
		if m.symbol != "" && t.Symbol != m.symbol {
			continue
		}
		cs = append(cs, t.CandleSticks...)
	}
	return cs, nil
}

// Symbol will specify a symbol for candlesticks next request
func (m *MockedCandleStickService) Symbol(symbol string) CandleStickServiceInterface {
	m.symbol = symbol
	return m
}
