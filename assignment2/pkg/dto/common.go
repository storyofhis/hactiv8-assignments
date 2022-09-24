package dto

import "assignment2/pkg/entity"

type GhidanAPIResponseDTO struct {
	Status struct{ 
		Code int `json:"code"`
		Description string `json:"description"`
	}`json:"status"`
	Result []entity.Person `json:"result"` 
}