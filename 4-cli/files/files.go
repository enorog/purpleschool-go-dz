package files

import (
	"fmt"
	"os"
	"strings"
)

func ReadJsonFile(path string) (content []byte, err error) {
	if !strings.HasSuffix(path, ".json") {
		return nil, fmt.Errorf("файл должен иметь расширение .json")
	}
	return os.ReadFile(path)
}
