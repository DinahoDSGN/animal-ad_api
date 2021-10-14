package services

import (
	"petcard/pkg/models"
	"petcard/pkg/repository"
)

type AdLocationService struct {
	repo repository.AdLocation
}

func (s *AdLocationService) Create(data models.AdLocation) (models.AdLocation, error) {
	return s.repo.Create(data)
}

func (s *AdLocationService) GetAll() ([]models.AdLocation, error) {
	return s.repo.GetAll()
}

func (s *AdLocationService) GetList(id int) (models.AdLocation, error) {
	return s.repo.GetList(id)
}

func NewAdLocationService(repo repository.AdLocation) *AdLocationService {
	return &AdLocationService{repo: repo}
}
