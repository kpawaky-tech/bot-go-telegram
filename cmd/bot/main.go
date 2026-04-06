package main

import (
	"log"

	"github.com/kpawaky-tech/bot-go-telegram/internal/bot"
	"github.com/kpawaky-tech/bot-go-telegram/internal/config"
)

func main() {
	// 1. Читаем конфиг
	cfg := config.Load()

	// 2. Создаём бота
	b, err := bot.New(cfg)
	if err != nil {
		log.Fatal("❌ Не удалось создать бота:", err)
	}

	// 3. Запускаем
	log.Println("✅ Бот запущен. Нажми /start")
	b.Start()
}
