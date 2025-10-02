package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	// Загружаем .env только один раз
	err := godotenv.Load()
	if err != nil {
		//log.Println("⚠️ Не удалось загрузить .env файл, использую только системные переменные")
	}

	key := os.Getenv("KEY")
	return &Config{Key: key}
}
