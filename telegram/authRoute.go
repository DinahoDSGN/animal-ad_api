package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"petcard/pkg/models"
	"petcard/pkg/services"
	"strings"
)

func (b *Bot) login(message string) (string, error) {
	usernameAndPassword := strings.Split(message, " ")

	username := usernameAndPassword[0]
	password := usernameAndPassword[1]

	result, err := b.service.Authorization.GenerateToken(username, password)
	if err != nil {
		return "", err
	}

	if result == "" {
		return "Invalid username of password", nil
	}

	return "You are logged in", nil
}

func (b *Bot) register(message string) (string, error) {
	usernameAndPassword := strings.Split(message, " ")

	username := usernameAndPassword[0]
	password := usernameAndPassword[1]

	data := models.User{Username: username, Password: password}

	result, err := b.service.Authorization.SignUp(data)
	if err != nil {
		return "", err
	}

	if result.Id == 0 {
		return "Something got wrong", nil
	}

	return "You register", nil
}

func (b *Bot) logout() {
	services.UserId = 0
}

func (b *Bot) isAuth(message *tgbotapi.Message) bool {
	if services.UserId == 0 {
		b.sendMessage(message.Chat.ID, "You are not logged in")
		return false
	}
	return true
}
