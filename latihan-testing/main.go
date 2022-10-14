package main

import (
	"latihan-testing/models"
	"latihan-testing/models/entity"
	"latihan-testing/router"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Cannot Load your env")
		return
	}
	log.Println("Success Load your env")

	models.ConnectDB()
	entity.DB.AutoMigrate(&entity.User{})
}
func main() {
	app := router.CreateRouter()
	app.Run(":8080")
}
