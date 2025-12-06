package main

import (
	"cli/api"
	"cli/config"
	"cli/files"
	"cli/storage"
	"fmt"
	"os"
)

const storagePath = "./storage.json"

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения конфига: %v\n", err)
		os.Exit(1)
	}
	api.UseConfig(config)
	fileProvider, err := files.NewJsonFile(storagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка инициализации провайдера: %v", err)
		os.Exit(1)
	}
	storage := storage.NewStorage(fileProvider)
	binList, _ := storage.Load()
	storage.Save(binList)
}
