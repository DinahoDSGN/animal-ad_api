package services

import (
	"petcard/internal/models"
	"petcard/pkg/repository"
)

type Ad interface {
	Create(data models.Ad) error
	GetAll() ([]models.Ad, error)
	GetList(id int) (models.Ad, error)
	Delete(id int) error
	Update(id int, data models.Ad) error
}

type User interface {
	GetAll() ([]models.User, error)
	GetList(id int) (models.User, error)
	Delete(id int) error
	Update(id int, data models.User) error
}

type Spec interface {
	Create(data models.Specify) (int, error)
	GetAll() ([]models.Specify, error)
	GetList(id int) (models.Specify, error)
	Delete(id int) (models.Specify, error)
	Update(id int, data models.Specify) (models.Specify, error)
}

type Service struct {
	User
	Ad
	Spec
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Ad:   NewAdService(repos.Ad),
		Spec: NewSpecService(repos.Spec),
	}
}
