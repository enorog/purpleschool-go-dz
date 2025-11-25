package main

import "cli/storage"

func main() {
	binList, _ := storage.Load()
	storage.Save(binList)
}
