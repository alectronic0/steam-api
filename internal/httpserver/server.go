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
