package server

import (
	"context"
	"net"
	"net/http"

	"github.com/DevAthhh/auth-service/pkg/config"
)

type Server struct {
	server *http.Server
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func NewServer(cfg *config.Config, routes http.Handler) *Server {
	addr := net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)
	return &Server{
		server: &http.Server{
			Handler:      routes,
			ReadTimeout:  cfg.Server.RTimeout,
			IdleTimeout:  cfg.Server.ITimeout,
			WriteTimeout: cfg.Server.WTimeout,
			Addr:         addr,
		},
	}
}
