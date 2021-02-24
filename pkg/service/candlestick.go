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
	Period(period int64) CandleStickServiceInterface
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

// Period will specify a period for candlesticks next request
func (s *CandleStickService) Period(period int64) CandleStickServiceInterface {
	interval, err := adapters.PeriodToInterval(period)
	if err != nil {
		interval = "unknown"
	}

	s.service.Interval(interval)
	return s
}

// Mock section
////////////////////////////////////////////////////////////////////////////////

// MockedCandleSticks are candlesticks that can be used in MockCandleStickService
type MockedCandleSticks struct {
	Symbol       string
	Period       int64
	CandleSticks []models.CandleStick
}

// MockedCandleStickService is the mocked service for candlesticks
type MockedCandleStickService struct {
	MockedCandleSticks []MockedCandleSticks

	// Next request specifications
	symbol string
	period int64
}

// Do will execute a request for candlesticks
func (m *MockedCandleStickService) Do(ctx context.Context) ([]models.CandleStick, error) {
	cs := make([]models.CandleStick, 0)
	for _, t := range m.MockedCandleSticks {
		// Check if symbol is set and correspond
		if m.symbol != "" && t.Symbol != m.symbol {
			continue
		}

		// Check if period is set and correspond
		if m.period != 0 && t.Period != m.period {
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

// Period will specify a period for candlesticks next request
func (m *MockedCandleStickService) Period(period int64) CandleStickServiceInterface {
	m.period = period
	return m
}
