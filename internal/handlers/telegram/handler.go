package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/usename-Poezd/light-notifier/internal/services"
)

type Handler struct {
	Services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) Init(bot *tgbotapi.BotAPI) {
	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":

			//Отправлем сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi!")
			bot.Send(msg)

		case "/check":
			err := h.Services.Keenetic.Check()
			msg := "OK"
			if err != nil {
				msg = "ERROR"
			}

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
		}
	}
}
