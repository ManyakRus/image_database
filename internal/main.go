package main

import (
	"github.com/ManyakRus/image_database/internal/config"
	"github.com/ManyakRus/image_database/internal/constants"
	"github.com/ManyakRus/image_database/internal/logic"
	"github.com/ManyakRus/image_database/pkg/graphml"
	ConfigMain "github.com/ManyakRus/starter/config"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/postgres_gorm"
)

func main() {
	StartApp()
}

func StartApp() {
	ConfigMain.LoadEnv()
	config.FillSettings()
	config.FillFlags()

	postgres_gorm.StartDB()
	postgres_gorm.GetConnection().Logger.LogMode(1)

	graphml.StartReadFile()

	FileName := config.Settings.FILENAME_GRAPHML
	log.Info("file graphml: ", FileName)
	log.Info("postgres host: ", postgres_gorm.Settings.DB_HOST)
	ok := logic.StartFillAll(FileName)
	if ok == false {
		println(constants.TEXT_HELP)
	}

}
