package repository

import (
	"gorm.io/gorm"
	"petcard/internal/models"
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

type Repository struct {
	User
	Ad
	Spec
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Spec: NewSpecRepo(db),
		User: NewUserRepo(db),
		Ad:   NewAdRepo(db),
	}
}
