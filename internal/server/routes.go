package server

import (
	"go-kafka-scope-api/internal/handler/kafka/admin"
	"go-kafka-scope-api/internal/handler/system"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterSystemHandler
func (s *Server) RegisterSystemHandlers() {
	// System operations for process management
	g := s.router.Group("/system")
	g.GET("/health", system.HealthHandler)
	g.GET("/version", system.VersionHandler)
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

// RegisterKafkaAdminHandlers
func (s *Server) RegisterKafkaAdminHandlers() {
	g := s.router.Group("/api/kafka/admin")
	g.GET("/config", admin.ReadConfig)
	g.GET("/cluster", admin.ReadCluster)
	g.GET("/topic", admin.ReadTopic)
	g.GET("/consumer-group", admin.ReadConsumerGroup)
}
