package files

import (
	"fmt"
	"os"
	"strings"
)

type JsonFile struct {
	path string
}

func NewJsonFile(path string) (*JsonFile, error) {
	if !strings.HasSuffix(path, ".json") {
		return nil, fmt.Errorf("файл должен иметь расширение .json")
	}
	return &JsonFile{
		path: path,
	}, nil
}
func (file *JsonFile) Read() (content []byte, err error) {
	return os.ReadFile(file.path)
}
func (file JsonFile) Write(content []byte) error {
	return os.WriteFile(file.path, content, os.FileMode(0660))
}
