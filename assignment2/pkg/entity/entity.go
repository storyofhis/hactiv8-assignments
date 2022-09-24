package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt  time.Time
	DeletedAt gorm.DeletedAt
}

type Item struct {
	ItemID uint `json:"item_id"`
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	Quantity uint `json:"quantity"`
	OrderID uint `json:"order_id"`
	Order *Order `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL; json:"order"`
	BaseModel
}

type Order struct {
	OrderID uint `gorm:"primaryKey" json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt time.Time `json:"orderedAt"`
	Items []Item `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"items"`
}

type OrderWithPerson struct {
	OrderID uint `gorm:"primaryKey" json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt time.Time `json:"orderedAt"`
	Items []Item `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"items"`
	Person Person
	BaseModel
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	UUID string `json:"uuid"`
}