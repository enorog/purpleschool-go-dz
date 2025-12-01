package storage

import (
	"cli/bins"
	"encoding/json"
	"fmt"
	"os"
)

type Storage struct {
	provider Provider
}

type Provider interface {
	Read() ([]byte, error)
	Write([]byte) error
}

func NewStorage(provider Provider) *Storage {
	return &Storage{
		provider: provider,
	}
}

func (storage *Storage) Save(binList bins.BinList) error {
	content, err := json.Marshal(binList)
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}
	err = storage.provider.Write(content)
	if err != nil {
		return fmt.Errorf("ошибка сохранения: %v", err)
	}
	return nil
}

func (storage *Storage) Load() (binList bins.BinList, err error) {
	content, err := storage.provider.Read()
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
