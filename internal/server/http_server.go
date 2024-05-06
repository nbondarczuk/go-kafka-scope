package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-kafka-scope-api/internal/config"
	"go-kafka-scope-api/internal/handler/system"
	"go-kafka-scope-api/internal/logging"
	"go-kafka-scope-api/internal/middleware"
)

// Server links handlers to paths via routes.
type Server struct {
	router *gin.Engine
}

// New creates server with gin framework.
func New(version string) (*Server, error) {
	system.SetVersion(version)
	gin.SetMode(gin.ReleaseMode)

	s := &Server{
		router: gin.New(),
	}

	// Apply required middleware. The order matters as request log should
	// go before response log.
	s.router.Use(middleware.ResponseLogger())
	s.router.Use(middleware.RequestLogger())

	// Register defined handlers.
	s.RegisterHandlers()

	return s, nil
}

// Run the gin server on routes.
func (s *Server) Run() error {
	port := config.ServerHTTPPort()
	logging.Logger.Info("Starting HTTP server", slog.String("port", port))
	return s.router.Run(":" + port)
}

// RegisterHandlers links handlers to API points.
func (s *Server) RegisterHandlers() {
	// System operations
	s.router.GET("/system/health", system.HealthHandler)
	s.router.GET("/system/version", system.VersionHandler)
	s.router.GET("/system/metrics", gin.WrapH(promhttp.Handler()))
}
