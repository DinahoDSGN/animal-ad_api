package services

import (
	models2 "petcard/pkg/models"
	"petcard/pkg/repository"
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) SignIn(data models2.User) (models2.User, error) {
	return s.repo.SignIn(data)
}

func (s *AuthorizationService) SignUp(data models2.User) (models2.User, error) {
	return s.repo.SignUp(data)
}
