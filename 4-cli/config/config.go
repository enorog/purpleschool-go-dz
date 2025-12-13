package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ошибка чтения .env файла")
	}
	key := os.Getenv("KEY")
	if key == "" {
		return nil, fmt.Errorf("переменная окружения KEY не определена")
	}
	return &Config{
		Key: key,
	}, nil
}
