package services

import (
	"petcard/pkg/repository"
)

type ParserService struct {
	repo repository.Parser
}

func NewParserService(repo repository.Parser) *ParserService {
	return &ParserService{repo: repo}
}

func (s *ParserService) Push() error {
	return s.repo.Push()
}
