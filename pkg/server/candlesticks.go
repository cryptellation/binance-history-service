package server

import (
	"context"
	"net/http"

	"github.com/cryptellation/models.go"
)

// CandleSticksRequest represents a request for candlesticks
// swagger:parameters CandleSticksRequest
type CandleSticksRequest struct {
	// in:body
}

// CandleSticksResponse represents a response for candlesticks
// swagger:response CandleSticksResponse
type CandleSticksResponse struct {
	// in:body
	CandleSticks []models.CandleStick
}

// CandleSticks will process a request for candlesticks
// swagger:route GET /candlesticks CandleSticks CandleSticksRequest
// This will returns candlesticks based on arguments
//
// Consumes:
//  - application/json
//
// Produces:
//  - application/json
//
// responses:
//   200: CandleSticksResponse
//   500: ErrorResponse
func (s *Server) CandleSticks(w http.ResponseWriter, r *http.Request) {
	var resp CandleSticksResponse
	var err error

	resp.CandleSticks, err = s.binance.NewCandleStickService().Do(context.TODO())
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpResponse(w, http.StatusOK, resp)
}
