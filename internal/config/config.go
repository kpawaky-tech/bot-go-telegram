package config

import "os"

// Config — публичная структура, чтобы main мог её прочитать
type Config struct {
	Token string
}

// Load читает переменные окружения. Возвращает дефолты, если не найдено.
func Load() Config {
	token := os.Getenv("TG_BOT_TOKEN")
	if token == "" {
		token = "your-default-token" // для локальной разработки
	}
	return Config{Token: token}
}

