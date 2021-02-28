package server

import binance "github.com/cryptellation/binance.go"

var version = "1.0.0"

// Server is the messenger binance server
type Server struct {
	binance binance.Interface
}

// New will create a new server
func New(binance binance.Interface) *Server {
	return &Server{
		binance: binance,
	}
}

// Version returns the server version
func (s *Server) Version() string {
	return version
}
