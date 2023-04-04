package app

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/usename-Poezd/light-notifier/internal/config"
	"github.com/usename-Poezd/light-notifier/internal/handlers/telegram"
	"github.com/usename-Poezd/light-notifier/internal/services"
)

func Run() {
	config, err := config.Init(".env")
	if err != nil {
		log.Fatalf("Can't init config")
	}

	log.Printf("App is started :)")

	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		panic(err)
	}

	services := services.NewServices(&services.Deps{
		KeeneticDnsDomain: config.KeeneticDnsDomain,
	})
	telegramHandler := telegram.NewHandler(services)

	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(func () {
		err := services.Keenetic.Check()
		if err != nil {
			bot.Send(tgbotapi.NewMessage(606329781, "ERROR"))
		}

		
	})
	s.StartAsync()

	bot.Debug = true
	telegramHandler.Init(bot)
}
