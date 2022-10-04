package entity

import "time"

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type DataStatus struct {
	WaterStatus string `json:"waterStatus"`
	WindStatus  string `json:"windStatus"`
}

// response to write your json
type Response struct {
	Data      Data       `json:"data"`
	Status    DataStatus `json:"status"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type Status struct {
	WaterStatus string `json:"waterStatus"`
	WindStatus  string `json:"windStatus"`
}
