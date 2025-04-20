package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"steam-api/internal/steamclient"
	"steam-api/internal/steamgamecomparator"
	"steam-api/internal/steamservice"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	apiKey := os.Getenv("STEAM_API_KEY")
	testUser1 := os.Getenv("TEST_USER_1")
	testUser2 := os.Getenv("TEST_USER_2")

	if apiKey == "" {
		log.Fatal("STEAM_API_KEY environment variable not set")
	}

	service := steamgamecomparator.New(
		steamservice.New(
			steamclient.New(apiKey),
		),
	)

	response, err := service.CompareOwnedGames(testUser1, testUser2)
	if err != nil {
		log.Fatalf("Error fetching games for user 1: %v", err)
	}

	var outData []byte
	outData, err = json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	if err = ioutil.WriteFile("comparison_result.json", outData, 0644); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Println("Comparison saved to comparison_result.json")
}
