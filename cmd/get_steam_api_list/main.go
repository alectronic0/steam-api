package main

import (
	"log"
	"steam-api/internal/config"
	"steam-api/internal/steamclient"
	"steam-api/internal/steamservice"
	"steam-api/pkg/utils"
)

const filename = "./output/steam_api_list_results.json"

func main() {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	steamClient := steamclient.New(appConfig.ApiKey)
	steamService := steamservice.New(steamClient)

	response, err := steamService.GetSupportApiList(appConfig.ApiKey, appConfig.TestUserID1, appConfig.TestAppID1)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.WritePrettyJSONFile(filename, response)
	if err != nil {
		log.Fatal(err)
	}
}
