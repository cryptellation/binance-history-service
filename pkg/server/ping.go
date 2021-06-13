package server

import "net/http"

// PingResponse represents a response for ping
// swagger:response PingResponse
type PingResponse struct {
	// in:body
	Message string
}

// Ping will process a request for ping
// swagger:route GET /ping Ping
// This will returns a string with "pong"
//
// Produces:
//  - application/json
//
// responses:
//   200: PingResponse
func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	httpResponse(w, http.StatusOK, PingResponse{Message: "pong"})
}
