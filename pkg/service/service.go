package service

import (
	binance "github.com/adshao/go-binance/v2"
)

// Interface is an interface for service
type Interface interface {
}

// Service represents the Binance service
type Service struct {
	client *binance.Client
}

// New will create a new binance service
func New(apiKey, secretKey string) Interface {
	return &Service{
		client: binance.NewClient(apiKey, secretKey),
	}
}

// MockedService represents the Binance service mocked
type MockedService struct {
}

// NewMock will create a mocked service
func NewMock() Interface {
	return &MockedService{}
}
