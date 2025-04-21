package main

import (
	"log"
	"steam-api/internal/config"
	"steam-api/internal/steamclient"
	"steam-api/internal/steamservice"
	"steam-api/pkg/utils"
)

const filename = "./output/list_steam_apps.json"

func main() {
	appConifg, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	steamClient := steamclient.New(appConifg.ApiKey)
	steamService := steamservice.New(steamClient)

	response, err := steamService.GetAppList()
	if err != nil {
		log.Fatal(err)
	}

	err = utils.WritePrettyJSONFile(filename, response)
	if err != nil {
		log.Fatal(err)
	}
}
