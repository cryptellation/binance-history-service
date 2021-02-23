package service

import "testing"

func TestNewService(t *testing.T) {
	s := New("apiKey", "secretKey")

	// Check nil
	if s == nil {
		t.Fatal("Service should not be nil")
	}

	// Check type
	if _, ok := s.(*Service); !ok {
		t.Fatal("Service is not the good type")
	}
}

func TestNewCandleStickService(t *testing.T) {
	// Create a service
	m := New("apiKey", "secretKey")

	// Check nil on new candlestick service
	s := m.NewCandleStickService()
	if s == nil {
		t.Fatal("New candlestick service should not be nil")
	}

	// Check type
	if _, ok := s.(*CandleStickService); !ok {
		t.Fatal("Service is not the good type")
	}
}

// Mock Test
////////////////////////////////////////////////////////////////////////////////

func TestNewMockService(t *testing.T) {
	m := NewMock()

	// Check nil
	if m == nil {
		t.Fatal("Mock service should not be nil")
	}

	// Check type
	if _, ok := m.(*MockedService); !ok {
		t.Fatal("Mock service is not the good type")
	}
}

func TestNewMockedCandleStickService(t *testing.T) {
	// Create a mock service
	m := NewMock()

	// Check nil on new candlestick service
	s := m.NewCandleStickService()
	if s == nil {
		t.Fatal("New candlestick service should not be nil")
	}

	// Check type
	if _, ok := s.(*MockedCandleStickService); !ok {
		t.Fatal("Service is not the good type")
	}
}
