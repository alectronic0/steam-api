package main

import (
	"fmt"
	"log"
	"os"
	"steam-api/internal/steamclient"
	"steam-api/internal/steamservice"
	"steam-api/pkg/utils"

	"github.com/joho/godotenv"
)

const (
	filename = "./output/list_steam_apps.json"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	apiKey := os.Getenv("STEAM_API_KEY")

	if apiKey == "" {
		log.Fatal("STEAM_API_KEY environment variable not set")
	}

	steamClient := steamclient.New(apiKey)
	steamService := steamservice.New(steamClient)

	response, err := steamService.GetAppList()
	if err != nil {
		log.Fatalf("Error fetching games for user 1: %v", err)
	}

	var outData []byte
	outData, err = utils.PrettyJSON(response)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	if err = os.WriteFile(filename, outData, 0644); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Println("Comparison saved to " + filename)
}
