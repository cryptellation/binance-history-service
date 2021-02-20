package service

import "testing"

func TestNewService(t *testing.T) {
	s := New("apiKey", "secretKey")
	if s == nil {
		t.Fatal("Service should not be nil")
	}
}

func TestNewMockService(t *testing.T) {
	m := NewMock()
	if m == nil {
		t.Fatal("Mock service should not be nil")
	}
}

func TestNewCandleStickService(t *testing.T) {
	// Create a mock service
	m := NewMock()

	// Get new candlestick service
	s := m.NewCandleStickService()
	if s == nil {
		t.Fatal("New candlestick service should not be nil")
	}
}
