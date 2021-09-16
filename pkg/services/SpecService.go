package services

import (
	"petcard/internal/models"
	"petcard/pkg/repository"
)

type SpecService struct {
	repo repository.Spec
}

func NewSpecService(repo repository.Spec) *SpecService {
	return &SpecService{repo: repo}
}

func (s *SpecService) Create(data models.Specify) error {
	return s.repo.Create(data)
}

func (s *SpecService) GetAll() ([]models.Specify, error) {
	return s.repo.GetAll()
}

func (s *SpecService) GetList(id int) (models.Specify, error) {
	return s.repo.GetList(id)
}

func (s *SpecService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *SpecService) Update(id int, data models.Specify) error {
	return s.repo.Update(id, data)
}
