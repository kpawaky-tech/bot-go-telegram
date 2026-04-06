package bot

import (
	"time"

	"github.com/kpawaky-tech/bot-go-telegram/internal/config"
	"github.com/kpawaky-tech/bot-go-telegram/internal/handler"
	"github.com/kpawaky-tech/bot-go-telegram/internal/service"
	tele "gopkg.in/telebot.v3"
)

// New создаёт и настраивает бота. Возвращает готовый к запуску экземпляр.
func New(cfg config.Config) (*tele.Bot, error) {
	// 1. Создаём зависимости
	userSvc := &service.UserService{}

	// 2. Инициализируем Telegram-клиент
	b, err := tele.NewBot(tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	// 3. Регистрируем обработчики
	b.Handle("/start", handler.StartHandler(userSvc))

	return b, nil
}
