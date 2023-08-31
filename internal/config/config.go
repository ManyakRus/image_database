package config

import (
	"os"
)

const FILENAME_GRAPHML = "connections.graphml"
const SERVICE_NAME = "Main"

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	DIRECTORY_SOURCE string
	FILENAME_GRAPHML string
	SERVICE_NAME     string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.DIRECTORY_SOURCE = os.Getenv("DIRECTORY_SOURCE")
	Settings.FILENAME_GRAPHML = os.Getenv("FILENAME_GRAPHML")
	Settings.SERVICE_NAME = os.Getenv("SERVICE_NAME")

	if Settings.DIRECTORY_SOURCE == "" {
		Settings.DIRECTORY_SOURCE = CurrentDirectory()
		//log.Panicln("Need fill DIRECTORY_SOURCE ! in os.ENV ")
	}

	if Settings.FILENAME_GRAPHML == "" {
		Settings.FILENAME_GRAPHML = FILENAME_GRAPHML
	}

	if Settings.SERVICE_NAME == "" {
		Settings.SERVICE_NAME = SERVICE_NAME
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
	if len(Args) > 3 {
		return
	}

	if len(Args) > 0 {
		Settings.DIRECTORY_SOURCE = Args[0]
	}
	if len(Args) > 1 {
		Settings.FILENAME_GRAPHML = Args[1]
	}
	if len(Args) > 2 {
		Settings.SERVICE_NAME = Args[2]
	}
}
