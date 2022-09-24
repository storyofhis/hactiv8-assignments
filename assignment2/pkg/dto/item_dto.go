package dto

type ItemInsertDTO struct {
	ItemCode string `json:"itemCode"`
	OrderID uint64 `json:"orderID"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}

type ItemUpdateDTO struct {
	ItemID uint64 `json:"itemID"`
	ItemCode string `json:"itemCode"`
	OrderID uint64 `json:"orderID"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}