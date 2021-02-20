package service

import (
	binance "github.com/adshao/go-binance/v2"
)

// CandleStickServiceInterface is the interface for candle stick services
type CandleStickServiceInterface interface {
}

// CandleStickService is the real service for candlesticks
type CandleStickService struct {
	service *binance.KlinesService
}

// MockedCandleStickService is the mocked service for candlesticks
type MockedCandleStickService struct {
}
