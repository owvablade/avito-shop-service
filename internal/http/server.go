package http

import (
	"avito-shop-service/internal/config"
	"avito-shop-service/internal/handler"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	router   *gin.Engine
	server   *http.Server
	cfg      *config.Config
	handlers *handler.Handler
}

func NewServer(cfg *config.Config, handlers *handler.Handler) *Server {
	router := gin.Default()
	return &Server{cfg: cfg, router: router, handlers: handlers}
}

func (s *Server) Start() {
	s.initRoutes()

	s.server = &http.Server{
		Addr:    ":" + s.cfg.AppPort,
		Handler: s.router.Handler(),
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\\n", err)
		}
	}()
}

func (s *Server) Close() error {
	if err := s.server.Shutdown(context.Background()); err != nil {
		return err
	}
	return nil
}
