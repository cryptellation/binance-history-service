package mock

import "github.com/cryptellation/binance-history-service/pkg/service"

// MockedService represents the Binance service mocked
type MockedService struct {
}

// NewMock will create a mocked service
func NewMock() service.Interface {
	return &MockedService{}
}

// NewCandleStickService will create a new candlestick service
func (m *MockedService) NewCandleStickService() service.CandleStickServiceInterface {
	return &MockedCandleStickService{}
}
