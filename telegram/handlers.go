package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"petcard/telegram/telegramModels"
	"strconv"
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
		a := message.Text
		msgBody, _ := b.adGetAd(a)
		b.sendMessage(message.Chat.ID, msgBody)
		return
	case "ad_get_all":
		msgBody, _ := b.adGetAll()
		b.sendMessage(message.Chat.ID, msgBody)
		return
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "such command does not exist")
		b.bot.Send(msg)
	}
}

func (b *Bot) sendMessage(message int64, messageBody string) {
	msg := tgbotapi.NewMessage(message, messageBody)
	b.bot.Send(msg)
}

func (b *Bot) adGetAd(id string) (string, error) {
	adId, _ := strconv.Atoi(id)

	record, _ := b.service.Ad.GetList(adId)

	msgTemplate := fmt.Sprintf(""+
		"Id: %v\n"+
		"Description: %s\n"+
		"Location: %s\n"+
		"Animal Name: %v\n"+
		"Animal Type: %s\n"+
		"Animal Breed Name: %s\n"+
		"Animal Price: %v\n"+
		"Animal Profit: %v\n\n",
		record.Id,
		record.Description,
		record.Location,
		record.Animal.Name,
		record.Animal.Type,
		record.Animal.Breed.Name,
		record.Animal.Price,
		record.Animal.Profit)

	return msgTemplate, nil
}

func (b *Bot) adGetAll() (string, error) {
	record, _ := b.service.Ad.GetAll()
	bytes, _ := json.Marshal(record)

	var ad telegramModels.Ad
	json.Unmarshal(bytes, &ad)

	var messageTemplate string
	for i := 0; i < len(ad); i++ {
		messageTemplate = messageTemplate + fmt.Sprintf(
			"Id: %v\n"+
				"Description: %s\n"+
				"Location: %s\n"+
				"Animal Name: %v\n"+
				"Animal Type: %s\n"+
				"Animal Breed Name: %s\n"+
				"Animal Price: %v\n"+
				"Animal Profit: %v\n\n",
			ad[i].AdID,
			ad[i].Description,
			ad[i].Location,
			ad[i].Animal.Name,
			ad[i].Animal.Type,
			ad[i].Animal.Breed.Name,
			ad[i].Animal.Price,
			ad[i].Animal.Profit)
	}

	msgTemplate := fmt.Sprintf("%v", messageTemplate)

	return msgTemplate, nil
}
