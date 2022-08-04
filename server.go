package blog

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}


// Run starts an HTTP server with a given address and handler
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
