package app

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"github.com/usename-Poezd/light-notifier/internal/config"
	"github.com/usename-Poezd/light-notifier/internal/handlers/telegram"
	"github.com/usename-Poezd/light-notifier/internal/repositories"
	"github.com/usename-Poezd/light-notifier/internal/services"

	_ "github.com/mattn/go-sqlite3"
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
	db, err := sqlx.Connect("sqlite3", "./db/light.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repos := repositories.NewRepositories(db)
	services := services.NewServices(repos, &services.Deps{
		KeeneticDnsDomain: config.KeeneticDnsDomain,
	})

	telegramHandler := telegram.NewHandler(services)

	s := gocron.NewScheduler(time.UTC)

	s.Every(30).Minutes().Do(func () {
		if !repos.Notification.Status() {
			return
		}

		users, err := repos.Users.GetAll()
		if err != nil {
			log.Fatalf(err.Error())
		}
		err = services.Keenetic.Check()
		if err != nil {
			// Notify all users in database
			for _, u := range users {
				bot.Send(tgbotapi.NewMessage(u.ChatId, "⚠️ ОШИБКА В АВТОМАТИЧЕСКОМ РЕЖИМЕ ⚠️\n Роутер не отвечает, скорее всего электричества нету!"))
			} 
		}		
	})
	s.StartAsync()

	telegramHandler.Init(bot)
}
