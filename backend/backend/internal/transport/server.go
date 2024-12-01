package transport

import (
	"context"
	"e-commerce/backend/config"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type Logger interface {
	Info(msg string)
}

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, r chi.Router) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.Server.Addr,
			Handler:      r,
			ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
		},
	}
}

func (s *Server) Run() error {
	log.Info().Msg("Server is running")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("Shutting down server...")
	return s.httpServer.Shutdown(ctx)
}
