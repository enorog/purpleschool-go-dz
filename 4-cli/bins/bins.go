package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	name      string
	createdAt time.Time
	private   bool
}

type BinList struct {
	bins []Bin
}

func NewBin(id, name string, createdAt time.Time, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("пустой идентификатор")
	}
	if name == "" {
		return nil, errors.New("пустое имя")
	}
	bin := Bin{
		id:        id,
		name:      name,
		createdAt: createdAt,
		private:   private,
	}
	return &bin, nil
}

func NewBinList(bins []Bin) (*BinList, error) {
	binList := BinList{
		bins: bins,
	}
	return &binList, nil
}
