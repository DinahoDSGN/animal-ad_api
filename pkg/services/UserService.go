package services

import (
	"petcard/internal/models"
	"petcard/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetList(id int) (models.User, error) {
	return s.repo.GetList(id)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *UserService) Update(id int, data models.User) error {
	return s.repo.Update(id, data)
}
