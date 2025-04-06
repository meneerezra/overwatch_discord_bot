package main

import (
	"fmt"
	"net/http"
	"overwatch_discord_bot/models"
	"overwatch_discord_bot/overfast"
	"overwatch_discord_bot/utils/jsonUtils"
	"overwatch_discord_bot/utils/jsonUtils/jsonModels"
)

func main() {

	var settings jsonModels.Settings
	_, err := jsonUtils.NewJsonManager("settings.json", &settings)
	if err != nil {
		return
	}


	overfastClient := overfast.OverfastClient{
		Client: models.Client{
			Client: &http.Client{},
			BaseURL: settings.OverfastURL,
		},
	}


	fmt.Println(overfastClient.GetPlayerByName())


}
