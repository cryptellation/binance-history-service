package server

import "net/http"

// Route is representing a route inside the server
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func (s *Server) setRoutes() {
	routes := []Route{
		{
			Name:        "Ping",
			Method:      "GET",
			Pattern:     "/ping",
			HandlerFunc: s.Ping,
		},
		{
			Name:        "CandleSticks",
			Method:      "GET",
			Pattern:     "/candlesticks",
			HandlerFunc: s.CandleSticks,
		},
	}

	for _, r := range routes {
		s.router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}
}
