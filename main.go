package main

import (
	"fmt"
	"overwatch_discord_bot/jsonUtils"
	"overwatch_discord_bot/jsonUtils/jsonModels"
)

func main() {

	var settings jsonModels.Settings
	_, err := jsonUtils.NewJsonManager("settings.json", &settings)
	if err != nil {
		return
	}





}
