package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"log"
	"petcard/pkg/services"
)

type Telegram struct {
	database *gorm.DB
	services *services.Service
}

func NewTelegram(database *gorm.DB, services *services.Service) *Telegram {
	return &Telegram{database: database, services: services}
}

func (t *Telegram) InitBot() {
	bot, err := tgbotapi.NewBotAPI("1973585695:AAHinmpTqZZHIWwi983eZM91YezGmRz4BHc")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegram := NewBot(bot, t.services)
	telegram.Start()
}
