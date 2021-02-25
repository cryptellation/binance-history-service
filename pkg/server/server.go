package server

import "github.com/cryptellation/binance-history-service/pkg/service"

var version = "1.0.0"

// Server is the messenger binance server
type Server struct {
	binance service.Interface
}

// New will create a new server
func New(binance service.Interface) *Server {
	return &Server{
		binance: binance,
	}
}

// Version returns the server version
func (s *Server) Version() string {
	return version
}
