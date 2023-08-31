package main

import (
	"github.com/ManyakRus/image_database/internal/config"
	"github.com/ManyakRus/image_database/internal/constants"
	"github.com/ManyakRus/image_database/internal/load_json"
	"github.com/ManyakRus/image_database/internal/logic"
	ConfigMain "github.com/ManyakRus/starter/config"
	"github.com/ManyakRus/starter/log"
)

func main() {
	StartApp()
}

func StartApp() {
	ConfigMain.LoadEnv()
	config.FillSettings()
	config.FillFlags()

	load_json.LoadJSON()

	FileName := config.Settings.FILENAME_GRAPHML
	log.Info("directory: ", config.Settings.DIRECTORY_SOURCE)
	log.Info("file graphml: ", FileName)
	log.Info("service name: ", config.Settings.SERVICE_NAME)
	ok := logic.StartFillAll(FileName)
	if ok == false {
		println(constants.TEXT_HELP)
	}

	////test
	//test1 := postgres_connect.Settings.DB_HOST
	//test2 := whatsapp_connect.Settings.WHATSAPP_PHONE_FROM
	//if test1+test2 == "" {
	//}
}
