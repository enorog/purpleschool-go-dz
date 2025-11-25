package storage

import (
	"cli/bins"
	"encoding/json"
	"fmt"
	"os"
)

const storage = "./storage.json"

func Save(binList bins.BinList) error {
	content, err := json.Marshal(binList)
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}
	err = os.WriteFile(storage, content, os.FileMode(0660))
	if err != nil {
		return fmt.Errorf("ошибка сохранения: %v", err)
	}
	return nil
}

func Load() (binList bins.BinList, err error) {
	content, err := os.ReadFile(storage)
	if os.IsNotExist(err) {
		binListRef, errbin := bins.NewBinList([]bins.Bin{})
		if errbin != nil {
			err = errbin
			return
		}
		binList = *binListRef
	} else if err != nil {
		err = fmt.Errorf("ошибка чтения из хранилища: %v", err)
		return
	}
	err = json.Unmarshal(content, &binList)
	if err != nil {
		err = fmt.Errorf("ошибка десереиализации: %v", err)
		return
	}
	return
}
