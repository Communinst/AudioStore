package server

import (
	"context"
	"net/http"
	"time"
)

// To provide extended control custom server's used
type Server struct {
	httpServer *http.Server
}

func New(addres string,
	handler *http.Handler,
	readTimeout time.Duration,
	writeTimeout time.Duration) *Server {

	return &Server{
		httpServer: &http.Server{
			Addr:         addres,
			Handler:      *handler,
			ReadTimeout:  readTimeout * time.Second,
			WriteTimeout: writeTimeout * time.Second,
		},
	}
}

func (s *Server) Close() {
	s.httpServer.Close()
}

func (s *Server) ShutDown(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}

func (s *Server) Run() {

}
