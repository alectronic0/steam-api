package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type HttpServerConfig struct {
	Addr              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

type AppConfig struct {
	ApiKey      string
	TestUserID1 string
	TestUserID2 string
	TestAppID1  string
}

type WebConfig struct {
	AppConfig  AppConfig
	HttpServer HttpServerConfig
}

func LoadAppConfig() (AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return AppConfig{}, fmt.Errorf("error loading .env file: %s", err)
	}

	apiKey := os.Getenv("STEAM_API_KEY")
	if apiKey == "" {
		return AppConfig{}, fmt.Errorf("STEAM_API_KEY environment variable not set")
	}

	testUser1 := os.Getenv("TEST_USER_1")
	testUser2 := os.Getenv("TEST_USER_2")

	return AppConfig{
		ApiKey:      apiKey,
		TestUserID1: testUser1,
		TestUserID2: testUser2,
	}, nil
}

func LoadHttpServerConfig() (WebConfig, error) {
	appConfig, err := LoadAppConfig()
	if err != nil {
		return WebConfig{}, err
	}

	return WebConfig{
		AppConfig: appConfig,
		HttpServer: HttpServerConfig{
			Addr:              ":8080",
			ReadTimeout:       10 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       10 * time.Second,
			MaxHeaderBytes:    1 << 20,
		},
	}, nil
}
