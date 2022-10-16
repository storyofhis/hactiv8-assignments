package repository

import (
	"context"
	"latihan-testing/models/entity"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	GetData(context.Context) ([]entity.User, error)
	// GetDataById() (entity.User, error)
	// CreateData() (entity.User, error)
}

type connRepository struct {
	connection *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &connRepository{
		connection: db,
	}
}

// GET USERS
func (db *connRepository) GetData(ctx context.Context) ([]entity.User, error) {
	var user []entity.User

	if err := db.connection.WithContext(ctx).Find(&user).Error; err != nil {
		log.Println("error on Context to find user")
		return nil, err
	}
	return user, nil
}
