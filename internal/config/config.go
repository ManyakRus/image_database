package config

import (
	"os"
)

const FILENAME_GRAPHML = "connections.graphml"

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	FILENAME_GRAPHML string
	INCLUDE_TABLES   string
	EXCLUDE_TABLES   string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.FILENAME_GRAPHML = os.Getenv("FILENAME_GRAPHML")
	Settings.INCLUDE_TABLES = os.Getenv("INCLUDE_TABLES")
	Settings.EXCLUDE_TABLES = os.Getenv("EXCLUDE_TABLES")

	if Settings.FILENAME_GRAPHML == "" {
		Settings.FILENAME_GRAPHML = FILENAME_GRAPHML
	}

	//
}

// CurrentDirectory - возвращает текущую директорию ОС
func CurrentDirectory() string {
	Otvet, err := os.Getwd()
	if err != nil {
		//log.Println(err)
	}

	return Otvet
}

// FillFlags - заполняет параметры из командной строки
func FillFlags() {
	Args := os.Args[1:]
	if len(Args) > 1 {
		return
	}

	if len(Args) > 0 {
		Settings.FILENAME_GRAPHML = Args[0]
	}
}
