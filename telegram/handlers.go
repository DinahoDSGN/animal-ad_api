package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "you typed start command")
		b.bot.Send(msg)
		return
	case "ad_get_ad":
		a := message.CommandArguments()
		msgBody, _ := b.adGetAd(a)
		if msgBody == "" {
			b.sendMessage(message.Chat.ID, "Failed to find ad by id")
		}
		b.sendMessage(message.Chat.ID, msgBody)
		return
	case "ad_get_all":
		msgBody, _ := b.adGetAll()
		b.sendMessage(message.Chat.ID, msgBody)
		return
	case "ad_create":
		a := message.CommandArguments()
		msgBody, _ := b.adCreate(a, message)
		b.sendMessage(message.Chat.ID, msgBody)
		return
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "such command does not exist")
		b.bot.Send(msg)
	}
}

func (b *Bot) sendMessage(message int64, messageBody string) {
	msg := tgbotapi.NewMessage(message, messageBody)
	msg.ParseMode = "markdown"
	b.bot.Send(msg)
}
