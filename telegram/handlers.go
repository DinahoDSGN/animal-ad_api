package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
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
		msgBody, _ := b.adGetAd(message.CommandArguments())
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

	msgTemplate := fmt.Sprintf("Ad id: %d \n Title: %s \n Description: %s \n Location: %s \n",
		record.Id, record.Title, record.Description, record.Location)

	return msgTemplate, nil
}

// TODO: JSON TO STRUCT?
func (b *Bot) adGetAll() (string, error) {
	record, _ := b.service.Ad.GetAll()
	bytes, _ := json.Marshal(record)

	a, _ := ioutil.ReadFile(string(bytes))
	var ad telegramModels.Ad

	json.Unmarshal(a, &ad)

	msgTemplate := fmt.Sprintf("data: %v", string(bytes))

	fmt.Println(ad)

	return msgTemplate, nil
}
