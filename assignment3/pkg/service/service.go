package service

import (
	"assignment3/helper"
	"assignment3/pkg/entity"
	"assignment3/pkg/repository"
	"time"
)

type Service interface {
	GetData() (entity.Response, error)
	UpdateData() error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetData() (entity.Response, error) {
	return s.repository.GetData()
}

func (s *service) UpdateData() error {
	var response entity.Response

	response.Data.Water = helper.GenerateRandomNumber(1, 100)
	response.Data.Wind = helper.GenerateRandomNumber(1, 100)

	response.Status = entity.DataStatus{
		WaterStatus: helper.CheckWaterStatus(response.Data.Water),
		WindStatus:  helper.CheckWindStatus(response.Data.Wind),
	}

	response.UpdatedAt = time.Now()

	return s.repository.UpdateData(response)
}
