package main

import (
	"latihan-testing/controllers"
	"latihan-testing/models"
	"latihan-testing/models/entity"
	"latihan-testing/repository"
	"latihan-testing/router"
	"latihan-testing/service"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
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
	var (
		db             *gorm.DB              = models.ConnectDB()
		UserRepository repository.Repository = repository.NewRepository(db)
		UserService    service.Service       = service.NewService(UserRepository)

		controller controllers.Controllers = controllers.NewControllers(UserService)
	)
	app := router.CreateRouter(controller)
	app.Run(":8080")
}
