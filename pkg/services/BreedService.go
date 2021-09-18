package services

import (
	"petcard/pkg/models"
	"petcard/pkg/repository"
)

type BreedService struct {
	repo repository.Breed
}

func NewBreedService(repo repository.Breed) *BreedService {
	return &BreedService{repo: repo}
}

func (s *BreedService) Create(data models.Breed) (int, error) {
	return s.repo.Create(data)
}

func (s *BreedService) GetAll() ([]models.Breed, error) {
	return s.repo.GetAll()
}

func (s *BreedService) GetList(id int) (models.Breed, error) {
	return s.repo.GetList(id)
}

func (s *BreedService) Delete(id int) (models.Breed, error) {
	return s.repo.Delete(id)
}

func (s *BreedService) Update(id int, data models.Breed) (models.Breed, error) {
	return s.repo.Update(id, data)
}
