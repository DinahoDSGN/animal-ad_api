package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"petcard/pkg/models"
	"petcard/telegram/telegramModels"
	"strconv"
	"strings"
)

// TODO: AUTHOR ID...
func (b *Bot) adCreate(message string, contact *tgbotapi.Message) (string, error) {
	asd := strings.Split(message, "\n")

	var data []string

	for i := 0; i < len(asd); i++ {
		data = append(data, asd[i])
	}

	inputFirstname := contact.Chat.FirstName
	inputLastname := contact.Chat.LastName
	inputUsername := contact.Chat.UserName

	inputGender, _ := strconv.ParseBool(data[5])
	inputVaccinated, _ := strconv.ParseBool(data[6])
	inputSpayed, _ := strconv.ParseBool(data[7])
	inputPassport, _ := strconv.ParseBool(data[8])
	inputGlobalPrice, _ := strconv.Atoi(data[12])
	inputPrice, _ := strconv.Atoi(data[13])

	var input = models.Ad{
		Title:       data[0],
		Location:    data[1],
		Description: data[15],
		Animal: &models.Animal{
			Name:       data[2],
			Type:       data[3],
			Color:      data[4],
			Gender:     inputGender,
			Vaccinated: inputVaccinated,
			Spayed:     inputSpayed,
			Passport:   inputPassport,
			Breed: &models.Breed{
				Name:        data[9],
				Type:        data[10],
				Wool:        data[11],
				GlobalPrice: int16(inputGlobalPrice),
			},
			Price: int16(inputPrice),
		},
		Author: &models.User{
			Name:     inputFirstname,
			Lastname: inputLastname,
			Username: inputUsername,
			Email:    data[14],
		},
	}

	record, _ := b.service.Ad.Create(input)

	fmt.Println(record.Animal.Breed.Name)

	return data[0], nil
}

func (b *Bot) adGetAd(id string) (string, error) {
	adId, _ := strconv.Atoi(id)
	fmt.Println(adId)
	record, _ := b.service.Ad.GetList(adId)

	if record.Id == 0 {
		return "", nil
	}

	msgTemplate := fmt.Sprintf(
		"*ID:* _%v_\n"+
			"*Username* _%v_\n"+
			"*Location:* _%v_\n"+
			"*Animal Name:* _%v_\n"+
			"*Animal Type:* _%v_\n"+
			"*Animal Breed:* _%v_\n"+
			"*Animal Price:* _%v_\n"+
			"*Animal Profit:* _%v_\n"+
			"*Description:* _%v_\n"+
			"*Owner:* %v _%v_\n"+
			"*Owners email:* _%v_\n\n",
		record.Id,
		record.Author.Username,
		record.Location,
		record.Animal.Name,
		record.Animal.Type,
		record.Animal.Breed.Name,
		record.Animal.Price,
		record.Animal.Profit,
		record.Description,
		record.Author.Name,
		record.Author.Lastname,
		record.Author.Email)

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
			"*ID:* _%v_\n"+
				"*Username* _%v_\n"+
				"*Location:* _%v_\n"+
				"*Animal Name:* _%v_\n"+
				"*Animal Type:* _%v_\n"+
				"*Animal Breed:* _%v_\n"+
				"*Animal Price:* _%v_\n"+
				"*Animal Profit:* _%v_\n"+
				"*Description:* _%v_\n"+
				"*Owner:* %v _%v_\n"+
				"*Owners email:* _%v_\n\n",
			ad[i].AdID,
			ad[i].Author.Username,
			ad[i].Location,
			ad[i].Animal.Name,
			ad[i].Animal.Type,
			ad[i].Animal.Breed.Name,
			ad[i].Animal.Price,
			ad[i].Animal.Profit,
			ad[i].Description,
			ad[i].Author.Name,
			ad[i].Author.Lastname,
			ad[i].Author.Email)
	}

	msgTemplate := fmt.Sprintf("%v", messageTemplate)

	return msgTemplate, nil
}
