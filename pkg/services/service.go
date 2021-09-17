package services

import (
	models2 "petcard/pkg/models"
	"petcard/pkg/repository"
)

type Ad interface {
	Create(data models2.Ad) (models2.Ad, error)
	GetAll() ([]models2.Ad, error)
	GetList(id int) (models2.Ad, error)
	Delete(id int) error
	Update(id int, data models2.Ad) error
}

type User interface {
	GetAll() ([]models2.User, error)
	GetList(id int) (models2.User, error)
	Delete(id int) error
	Update(id int, data models2.User) error
}

type Spec interface {
	Create(data models2.Specify) (int, error)
	GetAll() ([]models2.Specify, error)
	GetList(id int) (models2.Specify, error)
	Delete(id int) (models2.Specify, error)
	Update(id int, data models2.Specify) (models2.Specify, error)
}

type Parser interface {
	Push() error
}

type Service struct {
	User
	Ad
	Spec
	Parser
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:   NewUserService(repos.User),
		Ad:     NewAdService(repos.Ad),
		Spec:   NewSpecService(repos.Spec),
		Parser: NewParserService(repos.Parser),
	}
}
