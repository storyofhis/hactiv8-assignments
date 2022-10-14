package repository

import "gorm.io/gorm"

type Repository interface {
	GetData() error
}

type connRepository struct {
	connection *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &connRepository{
		connection: db,
	}
}
