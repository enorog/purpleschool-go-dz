package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	Private   bool      `json:"private"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id, name string, createdAt time.Time, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("пустой идентификатор")
	}
	if name == "" {
		return nil, errors.New("пустое имя")
	}
	bin := Bin{
		Id:        id,
		Name:      name,
		CreatedAt: createdAt,
		Private:   private,
	}
	return &bin, nil
}

func NewBinList(bins []Bin) (*BinList, error) {
	binList := BinList{
		Bins: bins,
	}
	return &binList, nil
}
