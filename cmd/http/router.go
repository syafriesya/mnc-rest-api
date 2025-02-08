package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func InitializeGin(serviceName string, env string) *gin.Engine {

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(apmgin.Middleware(router))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/app/health/liveness", "/app/health/readiness"},
	}))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:    []string{"Content-Type, Authorization, access-control-allow-origin, access-control-allow-headers"},
	}))

	return router
}
