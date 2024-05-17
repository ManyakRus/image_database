package main

import (
	"github.com/ManyakRus/image_database/internal/config"
	"github.com/ManyakRus/image_database/internal/constants"
	"github.com/ManyakRus/image_database/internal/logic"
	"github.com/ManyakRus/image_database/pkg/graphml"
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/postgres_gorm"
)

func main() {
	StartApp()
}

func StartApp() {
	ConfigMain.LoadENV_or_SettingsTXT()
	config.FillSettings()
	config.FillFlags()

	postgres_gorm.StartDB()
	postgres_gorm.GetConnection().Logger.LogMode(1)

	doc := graphml.StartReadFile()
	graphml.ClearElements_from_Document(doc)

	FileName := config.Settings.FILENAME_GRAPHML
	log.Info("file graphml: ", FileName)
	log.Info("postgres host: ", postgres_gorm.Settings.DB_HOST)
	ok := logic.StartFillAll(FileName, doc)
	if ok == false {
		println(constants.TEXT_HELP)
	}

}
