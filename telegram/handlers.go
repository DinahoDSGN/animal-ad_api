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
	case "help":
		msg := tgbotapi.NewMessage(message.Chat.ID, ""+
			"/login username password\n"+
			"/register username password\n"+
			"/logout\n\n"+
			"/ad_get_ad [id] --- get single ad information\n"+
			"/ad_get_all --- get all ads\n"+
			"/ad_create --- create an ad\n"+
			"/get_my_ads --- get my ads")
		b.bot.Send(msg)
		return
	case "login":
		params := message.CommandArguments()
		msgBody, _ := b.login(params)
		msg := tgbotapi.NewMessage(message.Chat.ID, msgBody)
		b.bot.Send(msg)
	case "register":
		params := message.CommandArguments()
		msgBody, _ := b.register(params)
		msg := tgbotapi.NewMessage(message.Chat.ID, msgBody)
		b.bot.Send(msg)
	case "logout":
		b.logout()
		msg := tgbotapi.NewMessage(message.Chat.ID, "You are logout")
		b.bot.Send(msg)
		return
	case "ad_get_ad":
		ok := b.isAuth(message)
		if !ok {
			return
		}
		a := message.CommandArguments()
		msgBody, _ := b.adGetAd(a)
		if msgBody == "" {
			b.sendMessage(message.Chat.ID, "Failed to find ad by id")
		}
		b.sendMessage(message.Chat.ID, msgBody)
		return
	case "ad_get_all":
		ok := b.isAuth(message)
		if !ok {
			return
		}
		msgBody, _ := b.adGetAll()
		b.sendMessage(message.Chat.ID, msgBody)
		return
	case "ad_create":
		ok := b.isAuth(message)
		if !ok {
			return
		}

		a := message.CommandArguments()
		msgBody, ok, _ := b.adCreate(a)
		if !ok {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Argument template: "+
				"\nTitle\nLocation\nAnimal Name\nAnimal Type\nAnimal Color\nGender\nVaccinated\nSpayed\nPassport\nBreed Name\nBreed Type\nWool\nGlobal Price\nPrice")
			b.bot.Send(msg)
		}
		b.sendMessage(message.Chat.ID, msgBody)
		return
	case "get_my_ads":
		ok := b.isAuth(message)
		if !ok {
			return
		}
		msgBody, _ := b.getMyAds()
		b.sendMessage(message.Chat.ID, msgBody)
		return
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Such command does not exist, type /help")
		b.bot.Send(msg)
	}
}

func (b *Bot) sendMessage(message int64, messageBody string) {
	msg := tgbotapi.NewMessage(message, messageBody)
	msg.ParseMode = "markdown"
	b.bot.Send(msg)
}
