package server

import (
	"context"
	"net/http"
)

// CandleSticks will process a request for candlesticks
func (s *Server) CandleSticks(w http.ResponseWriter, r *http.Request) {
	cs, err := s.binance.NewCandleStickService().Do(context.Background())
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpResponse(w, http.StatusOK, cs)
}
