package main

import (
	"cli/files"
	"cli/storage"
	"fmt"
	"os"
)

const storagePath = "./storage.json"

func main() {
	fileProvider, err := files.NewJsonFile(storagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка инициализации провайдера: %v", err)
		os.Exit(1)
	}
	storage := storage.NewStorage(fileProvider)
	binList, _ := storage.Load()
	storage.Save(binList)
}
