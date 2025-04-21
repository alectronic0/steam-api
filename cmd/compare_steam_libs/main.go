package main

import (
	"fmt"
	"log"
	"os"
	"steam-api/internal/steamclient"
	"steam-api/internal/steamgamecomparator"
	"steam-api/internal/steamservice"
	"steam-api/pkg/utils"
	"time"

	"github.com/joho/godotenv"
)

const (
	fileNameTemplate = "./output/comparison_result_%v_%v_%v.json"
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

	testUser1 := os.Getenv("TEST_USER_1")
	testUser2 := os.Getenv("TEST_USER_2")

	steamClient := steamclient.New(apiKey)
	steamService := steamservice.New(steamClient)
	comparatorService := steamgamecomparator.New(steamService)

	response, err := comparatorService.CompareOwnedGames(testUser1, testUser2)
	if err != nil {
		log.Fatalf("Error fetching games for user 1: %v", err)
	}

	var outData []byte
	outData, err = utils.PrettyJSON(response)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	filename := fmt.Sprintf(fileNameTemplate, testUser1, testUser2, time.Now().Unix())
	if err = os.WriteFile(filename, outData, 0644); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Println("Comparison saved to " + filename)
}
