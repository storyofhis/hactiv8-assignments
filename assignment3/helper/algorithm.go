package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + 1
}

func CheckWaterStatus(water int) string {
	if water <= 5 {
		return "aman"
	} else if water > 5 && water <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func CheckWindStatus(wind int) string {
	if wind <= 6 {
		return "aman"
	} else if wind > 6 && wind <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
