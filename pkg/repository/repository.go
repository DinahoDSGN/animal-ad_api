package repository

import (
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type Authorization interface {
	SignIn(data models.User) (models.User, error)
	SignUp(data models.User) (models.User, error)
	GetUser(username string, password string) (models.User, error)
}

type Ad interface {
	Create(data models.Ad) (models.Ad, error)
	GetAll() ([]models.Ad, error)
	GetMyAds(id int) ([]models.Ad, error)
	GetList(id int) (models.Ad, error)
	Delete(id int) (models.Ad, error)
	Update(id int, data models.Ad) (models.Ad, error)
}

type User interface {
	GetAll() ([]models.User, error)
	GetList(id int) (models.User, error)
	Delete(id int) (models.User, error)
	Update(id int, data models.User) (models.User, error)
}

type Animal interface {
	Create(data models.Animal) (int, error)
	GetAll() ([]models.Animal, error)
	GetList(id int) (models.Animal, error)
	Delete(id int) (models.Animal, error)
	Update(id int, data models.Animal) (models.Animal, error)
}

type Breed interface {
	Create(data models.Breed) (int, error)
	GetAll() ([]models.Breed, error)
	GetList(id int) (models.Breed, error)
	Delete(id int) (models.Breed, error)
	Update(id int, data models.Breed) (models.Breed, error)
}

type Parser interface {
	Push() error
}

type Repository struct {
	Authorization
	User
	Ad
	Animal
	Breed
	Parser
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepo(db),
		Animal:        NewAnimalRepo(db),
		User:          NewUserRepo(db),
		Ad:            NewAdRepo(db),
		Breed:         NewBreedRepo(db),
		Parser:        NewParserRepo(db),
	}
}
