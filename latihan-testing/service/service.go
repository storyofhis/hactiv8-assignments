package service

import (
	"context"
	"latihan-testing/models/entity"
	"latihan-testing/repository"
)

type Service interface {
	GetData(context.Context) ([]entity.User, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetData(ctx context.Context) ([]entity.User, error) {
	result, err := s.repo.GetData(ctx)
	return result, err
}
