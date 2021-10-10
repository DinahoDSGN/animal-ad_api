package services

import (
	"petcard/pkg/models"
	"petcard/pkg/repository"
)

type AdSortsService struct {
	repo repository.AdSorts
}

func NewAdSortsService(repo repository.AdSorts) *AdSortsService {
	return &AdSortsService{repo: repo}
}

func (s AdSortsService) SortBy(values map[string]interface{}) ([]models.Ad, error) {
	return s.repo.SortBy(values)
}
