package operations

import (
	"bytes"
	"cli/api"
	"cli/bins"
	"fmt"
	"os"
)

type Operations struct {
	api     *api.JsonBin
	binList *bins.BinList
}

func NewOperation(api *api.JsonBin, binList *bins.BinList) *Operations {
	return &Operations{
		api:     api,
		binList: binList,
	}
}

func (operations *Operations) Create(path string, name string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	bin, err := operations.api.Create(name, bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	operations.binList.Bins = append(operations.binList.Bins, *bin)
	return nil
}

func (operations *Operations) Update(path string, id string) error {
	found := false
	for _, bin := range operations.binList.Bins {
		if bin.Id == id {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("bin с id=%s не найден", id)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = operations.api.Update(id, bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	return nil
}

func (operations *Operations) Delete(id string) error {
	index := -1
	for idx, bin := range operations.binList.Bins {
		if bin.Id == id {
			index = idx
			break
		}
	}
	if index < 0 {
		return fmt.Errorf("bin с id=%s не найден", id)
	}
	err := operations.api.Delete(id)
	if err != nil {
		return err
	}
	if index == len(operations.binList.Bins[:index])-1 {
		operations.binList.Bins = operations.binList.Bins[:index]
	} else {
		operations.binList.Bins = append(operations.binList.Bins[:index], operations.binList.Bins[index+1:]...)
	}
	return nil
}

func (operations *Operations) Get(id string) error {
	content, err := operations.api.Get(id)
	if err != nil {
		return err
	}
	fmt.Print(string(content))
	return nil
}

func (operations *Operations) List() error {
	for _, bin := range operations.binList.Bins {
		fmt.Printf("%s %s\n", bin.Name, bin.Id)
	}
	return nil
}
