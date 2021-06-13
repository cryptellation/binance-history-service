package server

import (
	"net/http"
	"time"

	"github.com/cryptellation/binance.go/pkg/binance"
	"github.com/gorilla/mux"
)

var version = "1.0.0"

// Server is the messenger binance server
type Server struct {
	binance binance.ServiceInterface
	router  *mux.Router
}

// New will create a new server
func New(binance binance.ServiceInterface) *Server {
	s := &Server{
		binance: binance,
		router:  mux.NewRouter(),
	}

	s.setRoutes()

	return s
}

// Version returns the server version
func (s *Server) Version() string {
	return version
}

func (s *Server) Serve(addr string) error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
