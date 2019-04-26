package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

// shutdownTimer the time that is given to the server to gracefully shut down. After expiration, the server will shut down immediately.
const shutdownTimer = 30 * time.Second

// Config the configuration server stores information about the server settings
type Config struct {
	Logger  log.Logger
	Addr    string
	Storage Storage
}

// A Server defines parameters for running an HTTP server.
type Server struct {
	srv     *http.Server
	cfg     *Config
	storage Storage
}

// New create http Server
func New(cfg *Config) (*Server, error) {
	if cfg.Logger == nil {
		cfg.Logger = log.NewNopLogger()
	}
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	newHandler(&handlerConfig{
		router: subrouter,
		logger: cfg.Logger,
		svc: &basicService{
			logger:  cfg.Logger,
			storage: cfg.Storage,
		},
	})
	server := &Server{
		srv: &http.Server{
			Addr:         cfg.Addr,
			Handler:      router,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
		cfg: cfg,
	}
	return server, nil
}

// Run server by Addr in srv config
func (s *Server) Run() error {
	if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Shutdown gracefully shut down the server
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimer)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
