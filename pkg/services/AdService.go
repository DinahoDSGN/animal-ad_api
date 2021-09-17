package services

import (
	"petcard/pkg/models"
	"petcard/pkg/repository"
)

type AdService struct {
	repo repository.Ad
}

func NewAdService(repo repository.Ad) *AdService {
	return &AdService{repo: repo}
}

func (s *AdService) Create(data models.Ad) (models.Ad, error) {
	return s.repo.Create(data)
}

func (s *AdService) GetAll() ([]models.Ad, error) {
	return s.repo.GetAll()
}

func (s *AdService) GetList(id int) (models.Ad, error) {
	return s.repo.GetList(id)
}

func (s *AdService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *AdService) Update(id int, data models.Ad) error {
	return s.repo.Update(id, data)
}
