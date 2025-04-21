package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"steam-api/internal/config"
	"steam-api/internal/httpserver/handlers"
	"steam-api/internal/steamclient"
	"steam-api/internal/steamservice"
)

func New(config config.WebConfig) (*http.Server, error) {
	apiKey := config.AppConfig.ApiKey

	client := steamclient.New(apiKey)
	service := steamservice.New(client)

	router := gin.Default()
	router.GET("/user", handlers.GetUserInfo(service))

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
