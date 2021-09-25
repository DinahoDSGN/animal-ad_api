package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"petcard/pkg/services"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	service *services.Service
}

func NewBot(bot *tgbotapi.BotAPI, service *services.Service) *Bot {
	return &Bot{bot: bot, service: service}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := b.bot.GetUpdatesChan(u)

	b.handleUpdates(updates)
	return nil
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Get ad", "/ad_get_ad 6"),
		tgbotapi.NewInlineKeyboardButtonData("All data", "/ad_get_all"),
	),
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		fmt.Println(update.Message.Chat.UserName)

		b.handleMessage(update.Message)
	}
}
