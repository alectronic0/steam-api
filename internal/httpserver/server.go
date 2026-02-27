package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"steam-api/internal/cache"
	"steam-api/internal/config"
	"steam-api/internal/httpserver/handlers"
	"steam-api/internal/steamclient"
	"steam-api/internal/steamservice"
)

// CORSMiddleware adds CORS headers to responses
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func New(config config.WebConfig) (*http.Server, error) {
	apiKey := config.AppConfig.ApiKey

	client := steamclient.New(apiKey)
	service := steamservice.New(client)

	// Initialize Redis cache
	redisCache := cache.NewRedisCache()
	ctx := context.Background()
	if err := redisCache.Connect(ctx); err != nil {
		fmt.Printf("Warning: Redis cache connection failed: %v (continuing without cache)\n", err)
	}

	router := gin.Default()

	// Add CORS middleware
	router.Use(CORSMiddleware())

	router.GET("/user", handlers.GetUserInfo(service, redisCache))
	router.GET("/health", handlers.Health())

	return &http.Server{
		Addr:              config.HttpServer.Addr,
		Handler:           router,
		ReadTimeout:       config.HttpServer.ReadTimeout,
		ReadHeaderTimeout: config.HttpServer.ReadHeaderTimeout,
		WriteTimeout:      config.HttpServer.WriteTimeout,
		IdleTimeout:       config.HttpServer.IdleTimeout,
		MaxHeaderBytes:    config.HttpServer.MaxHeaderBytes,
	}, nil
}
