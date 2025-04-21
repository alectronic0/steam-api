package main

import (
	"fmt"
	"log"
	"steam-api/internal/config"
	"steam-api/internal/steamclient"
	"steam-api/internal/steamgamecomparator"
	"steam-api/internal/steamservice"
	"steam-api/pkg/utils"
	"time"
)

const fileNameTemplate = "./output/comparison_result_%v_%v_%v.json"

func main() {
	appConifg, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	steamClient := steamclient.New(appConifg.ApiKey)
	steamService := steamservice.New(steamClient)
	comparatorService := steamgamecomparator.New(steamService)

	response, err := comparatorService.CompareOwnedGames(appConifg.TestUserID1, appConifg.TestUserID2)
	if err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf(fileNameTemplate, appConifg.TestUserID1, appConifg.TestUserID2, time.Now().Unix())
	err = utils.WritePrettyJSONFile(filename, response)
	if err != nil {
		log.Fatal(err)
	}
}
