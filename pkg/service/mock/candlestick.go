package mock

import (
	"context"
	"time"

	"github.com/cryptellation/binance-messenger-service/pkg/service"
	"github.com/cryptellation/models.go"
)

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
	symbol  string
	period  int64
	endTime time.Time
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

		// Check each candle
		for _, c := range t.CandleSticks {
			// Check if endtime is send and correspond
			if !m.endTime.IsZero() && m.endTime.After(c.Time) {
				continue
			}

			// Add it if it passed tests
			cs = append(cs, c)
		}
	}
	return cs, nil
}

// Symbol will specify a symbol for next candlesticks request
func (m *MockedCandleStickService) Symbol(symbol string) service.CandleStickServiceInterface {
	m.symbol = symbol
	return m
}

// Period will specify a period for next candlesticks request
func (m *MockedCandleStickService) Period(period int64) service.CandleStickServiceInterface {
	m.period = period
	return m
}

// EndTime will specify the time where the list ends (earliest time) for
// next candlesticks request
func (m *MockedCandleStickService) EndTime(endTime time.Time) service.CandleStickServiceInterface {
	m.endTime = endTime
	return m
}
