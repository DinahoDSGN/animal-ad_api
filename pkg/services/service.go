package services

import (
	models2 "petcard/pkg/models"
	"petcard/pkg/repository"
)

type Authorization interface {
	SignIn(data models2.User) (models2.User, error)
	SignUp(data models2.User) (models2.User, error)
}

type Ad interface {
	Create(data models2.Ad) (models2.Ad, error)
	GetAll() ([]models2.Ad, error)
	GetList(id int) (models2.Ad, error)
	Delete(id int) (models2.Ad, error)
	Update(id int, data models2.Ad) (models2.Ad, error)
}

type User interface {
	GetAll() ([]models2.User, error)
	GetList(id int) (models2.User, error)
	Delete(id int) (models2.User, error)
	Update(id int, data models2.User) (models2.User, error)
}

type Animal interface {
	Create(data models2.Animal) (int, error)
	GetAll() ([]models2.Animal, error)
	GetList(id int) (models2.Animal, error)
	Delete(id int) (models2.Animal, error)
	Update(id int, data models2.Animal) (models2.Animal, error)
}

type Breed interface {
	Create(data models2.Breed) (int, error)
	GetAll() ([]models2.Breed, error)
	GetList(id int) (models2.Breed, error)
	Delete(id int) (models2.Breed, error)
	Update(id int, data models2.Breed) (models2.Breed, error)
}

type Parser interface {
	Push() error
}

type Service struct {
	Authorization
	User
	Ad
	Animal
	Breed
	Parser
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.Authorization),
		User:          NewUserService(repos.User),
		Ad:            NewAdService(repos.Ad),
		Animal:        NewAnimalService(repos.Animal),
		Breed:         NewBreedService(repos.Breed),
		Parser:        NewParserService(repos.Parser),
	}
}
