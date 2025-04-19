package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func compareLibrariesAndBuildOutput(user1, user2 APIResponse) Output {
	shared := make(map[int]bool)
	for appID := range user1.Response.Games {
		if _, ok := user2[appID]; ok {
			shared[appID] = true
		}
	}
	return Output{
		User1ID: User{
			Name:  "",
			ID:    0,
			Games: nil,
		},
		User2ID: User{
			Name:  "",
			ID:    0,
			Games: nil,
		},
		SharedGames: nil,
	}
}

func main() {
	if ApiKey == "" {
		log.Fatal("STEAM_API_KEY environment variable not set")
	}

	lib1, err := getOwnedGames(ApiKey, SteamID1)
	if err != nil {
		log.Fatalf("Error fetching games for user 1: %v", err)
	}

	lib2, err := getOwnedGames(ApiKey, SteamID2)
	if err != nil {
		log.Fatalf("Error fetching games for user 2: %v", err)
	}

	output := compareLibrariesAndBuildOutput(lib1, lib2)

	var outData []byte
	outData, err = json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	if err = ioutil.WriteFile("comparison_result.json", outData, 0644); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Println("Comparison saved to comparison_result.json")
}
