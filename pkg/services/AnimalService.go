package services

import (
	"petcard/pkg/models"
	"petcard/pkg/repository"
)

type AnimalService struct {
	repo repository.Animal
}

func NewAnimalService(repo repository.Animal) *AnimalService {
	return &AnimalService{repo: repo}
}

func (s *AnimalService) Create(data models.Animal) (int, error) {
	return s.repo.Create(data)
}

func (s *AnimalService) GetAll() ([]models.Animal, error) {
	return s.repo.GetAll()
}

func (s *AnimalService) GetList(id int) (models.Animal, error) {
	return s.repo.GetList(id)
}

func (s *AnimalService) Delete(id int) (models.Animal, error) {
	return s.repo.Delete(id)
}

func (s *AnimalService) Update(id int, data models.Animal) (models.Animal, error) {
	return s.repo.Update(id, data)
}
