package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/usename-Poezd/light-notifier/internal/domain"
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

			_, err := h.Services.Users.Create(&domain.User{
				Id: 0,
				ChatId: update.Message.Chat.ID,
			})
			if err != nil {
				log.Fatalln(err.Error())
			}
			//Отправлем сообщение
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID, 
				"👋Привет, я бот по проверки электричества! \n" +
				"Каждый 30 минут я проверяю электричество путем ping'а домена роутера. \n \n"+

				`Если вы начали общение со мной, то вы уже автоматически подписались на рассылку.` + "\n\n" +

				`Мои команды:` + "\n" +
				"- /check - проверка электричества\n" +
				"- /off - вырубить оповещения для всех пользователей\n" +
				"- /on - врубить оповещения для всех пользователей",
			)
			bot.Send(msg)

		case "/check":
			err := h.Services.Keenetic.Check()
			msg := "✅ OK \n Все круто, электричестов есть! 💡"
			if err != nil {
				msg = "⚠️ ОШИБКА ⚠️\n Роутер не отвечает, скорее всего электричества нету!"
			}

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
		
		case "/off":
			h.Services.Notification.Off()
			

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Уведомления отключены для всех пользователей"))
		case "/on":
			h.Services.Notification.On()
			
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Уведомления включены для всех пользователей"))
		}
	}
}
