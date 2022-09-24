package dto

import "time"

type OrderCreateDTO struct {
	CustomerName string `json:"customer_name" binding:"required"`
	Items []ItemInsertDTO `json:"items" binding:"required"`
	OrderedAt time.Time `json:"ordered_at"`
}

type OrderUpdateDTO struct {
	OrderID int64 `json:"order_id" binding:"required"`
	CustomerName string `json:"customer_name" binding:"required"`
	Items []ItemUpdateDTO `json:"items" binding:"required"`
	OrderedAt time.Time `json:"ordered_at"`
}