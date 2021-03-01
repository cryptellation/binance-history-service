package server

import (
	binance "github.com/cryptellation/binance.go"
	"github.com/gorilla/mux"
)

var version = "1.0.0"

// Server is the messenger binance server
type Server struct {
	binance binance.Interface
	router  *mux.Router
}

// New will create a new server
func New(binance binance.Interface) *Server {
	// Create server
	s := &Server{
		binance: binance,
		router:  mux.NewRouter(),
	}

	// Set routes
	s.setRoutes()

	return s
}

// Version returns the server version
func (s *Server) Version() string {
	return version
}
