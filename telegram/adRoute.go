package telegram

import (
	"encoding/json"
	"fmt"
	"petcard/pkg/models"
	"petcard/pkg/services"
	"petcard/telegram/telegramModels"
	"strconv"
	"strings"
)

func (b *Bot) adCreate(message string) (string, bool, error) {
	asd := strings.Split(message, "\n")

	if message == "" {
		return "Empty arguments", false, nil
	}

	var data []string

	for i := 0; i < len(asd); i++ {
		data = append(data, asd[i])
	}

	inputGender, _ := strconv.ParseBool(data[5])
	inputVaccinated, _ := strconv.ParseBool(data[6])
	inputSpayed, _ := strconv.ParseBool(data[7])
	inputPassport, _ := strconv.ParseBool(data[8])
	inputGlobalPrice, _ := strconv.Atoi(data[12])
	inputPrice, _ := strconv.Atoi(data[13])
	userId := b.service.Authorization.GetUserId()

	input := models.Ad{
		Title:    data[0],
		Location: data[1],
		//Description: data[14],
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
		AuthorId: uint(userId),
	}

	record, err := b.service.Ad.Create(input)
	if err != nil {
		return "", false, err
	}

	fmt.Println(record.AuthorId)

	msg := "Ad " + record.Title + " successfully created!"

	return msg, true, nil
}

func (b *Bot) getMyAds() (string, error) {
	userId := services.UserId

	record, _ := b.service.Ad.GetMyAds(userId)
	bytes, _ := json.Marshal(record)

	var ad telegramModels.Ad
	err := json.Unmarshal(bytes, &ad)
	if err != nil {
		return "", err
	}

	var messageTemplate string
	for i := 0; i < len(ad); i++ {
		messageTemplate = messageTemplate + fmt.Sprintf(
			"*ID:* _%v_\n"+
				"*AdLocation:* _%v_\n"+
				"*Animal Name:* _%v_\n"+
				"*Animal Type:* _%v_\n"+
				"*Animal Breed:* _%v_\n"+
				"*Animal Price:* _%v_\n"+
				"*Animal Profit:* _%v_\n"+
				"*Description:* _%v_\n"+
				"*Owner:* %v\n"+
				"*Owners email:* _%v_\n\n",
			ad[i].AdID,
			ad[i].Location,
			ad[i].Animal.Name,
			ad[i].Animal.Type,
			ad[i].Animal.Breed.Name,
			ad[i].Animal.Price,
			ad[i].Animal.Profit,
			ad[i].Description,
			ad[i].Author.Username,
			ad[i].Author.Email)
	}

	msgTemplate := fmt.Sprintf("%v", messageTemplate)

	return msgTemplate, nil
}

func (b *Bot) adGetAd(id string) (string, error) {
	adId, _ := strconv.Atoi(id)
	fmt.Println(adId)
	record, _ := b.service.Ad.GetList(adId)

	if record.Id == 0 {
		return "", nil
	}

	msgTemplate := fmt.Sprintf(
		"ID: %v\n"+
			"Username %v\n"+
			"AdLocation: %v\n"+
			"Animal Name: %v\n"+
			"Animal Type: %v\n"+
			"Animal Breed: %v\n"+
			"Animal Price: %v\n"+
			"Animal Profit: %v\n"+
			"Description: %v\n"+
			"Owner: %v %v\n"+
			"Owners email: %v\n\n",
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
	err := json.Unmarshal(bytes, &ad)
	if err != nil {
		return "", err
	}

	var messageTemplate string
	for i := 0; i < len(ad); i++ {
		messageTemplate = messageTemplate + fmt.Sprintf(
			"*ID:* _%v_\n"+
				"*AdLocation:* _%v_\n"+
				"*Animal Name:* _%v_\n"+
				"*Animal Type:* _%v_\n"+
				"*Animal Breed:* _%v_\n"+
				"*Animal Price:* _%v_\n"+
				"*Animal Profit:* _%v_\n"+
				"*Description:* _%v_\n"+
				"*Owner:* %v\n"+
				"*Owners email:* _%v_\n\n",
			ad[i].AdID,
			ad[i].Location,
			ad[i].Animal.Name,
			ad[i].Animal.Type,
			ad[i].Animal.Breed.Name,
			ad[i].Animal.Price,
			ad[i].Animal.Profit,
			ad[i].Description,
			ad[i].Author.Username,
			ad[i].Author.Email)
	}

	msgTemplate := fmt.Sprintf("%v", messageTemplate)

	return msgTemplate, nil
}
