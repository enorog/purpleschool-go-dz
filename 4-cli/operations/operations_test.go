package operations_test

import (
	"cli/api"
	"cli/config"
	"cli/files"
	"cli/operations"
	"cli/storage"
	"encoding/json"
	"os"
	"testing"
)

const storagePath = "./storage_tests.json"
const binPath = "./bin_test.json"

func getOperations(storagePath string) *operations.Operations {
	err := os.Remove(storagePath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	jsonBin := api.NewJsonBin(config)
	fileProvider, err := files.NewJsonFile(storagePath)
	if err != nil {
		panic(err)
	}
	storage := storage.NewStorage(fileProvider)
	binList, err := storage.Load()
	if err != nil {
		panic(err)
	}
	return operations.NewOperation(jsonBin, binList)
}

func createBin(binPath string) string {
	err := os.Remove(binPath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	bin := map[string]string{
		"test": "test",
	}
	binSerialized, err := json.Marshal(bin)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(binPath, binSerialized, os.FileMode(0640))
	if err != nil {
		panic(err)
	}
	return string(binSerialized)
}

func TestCreate(t *testing.T) {
	ops := getOperations(storagePath)
	createBin(binPath)

	id, err := ops.Create(binPath, "test")
	if err != nil {
		t.Errorf("Create(\"%s\",\"%s\"): ожидался успех, получили %v", binPath, "test", err)
	} else if id == "" {
		t.Errorf("Create(\"%s\",\"%s\"): ожидалcя непустой id", binPath, "test")
	}
	os.Remove(binPath)
	os.Remove(storagePath)
}

func TestUpdate(t *testing.T) {
	ops := getOperations(storagePath)
	createBin(binPath)

	id, err := ops.Create(binPath, "test")
	if err != nil {
		panic(err)
	}

	id, err = ops.Update(binPath, id)
	if err != nil {
		t.Errorf("Update(\"%s\",\"%s\"): ожидался успех, получили %v", binPath, id, err)
	} else if id == "" {
		t.Errorf("Update(\"%s\",\"%s\"): ожидалcя непустой id", binPath, id)
	}
	os.Remove(binPath)
	os.Remove(storagePath)
}

func TestGet(t *testing.T) {
	ops := getOperations(storagePath)
	expectedBinContent := createBin(binPath)

	id, err := ops.Create(binPath, "test")
	if err != nil {
		panic(err)
	}

	binContent, err := ops.Get(id)
	if err != nil {
		t.Errorf("Update(\"%s\",\"%s\"): ожидался успех, получили %v", binPath, id, err)
	} else if binContent != expectedBinContent {
		t.Errorf("Get(\"%s\"): ожидалось %s получли %s", id, expectedBinContent, binContent)
	}
	os.Remove(binPath)
	os.Remove(storagePath)
}
