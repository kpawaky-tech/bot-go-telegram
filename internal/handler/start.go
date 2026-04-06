package handler

import (
	"github.com/kpawaky-tech/bot-go-telegram/internal/service"
	tele "gopkg.in/telebot.v3"
)

// StartHandler возвращает функцию-обработчик для /start
// Зависит от service, но не от config или bot-клиента
func StartHandler(svc *service.UserService) tele.HandlerFunc {
	return func(c tele.Context) error {
		user := c.Sender()
		reply := svc.GreetUser(user.ID, user.FirstName)
		return c.Send(reply)
	}
}
