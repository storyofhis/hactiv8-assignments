package main

import (
	"assignment2/pkg/config"
	"assignment2/pkg/controller"
	"assignment2/pkg/service"
	"assignment2/repo"
	"os"

	"assignment2/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	var (
		db               *gorm.DB                    = config.SetupDatabaseConnection()
		OrderRepository  repo.OrderRepository  = repo.NewOrderRepository(db)
		ItemRepository   repo.ItemRepository   = repo.NewItemRepository(db)
		PersonRepository repo.PersonRepository = repo.NewPersonRepository(os.Getenv("GIDHAN_API"))

		orderService service.OrderService = service.NewOrderService(OrderRepository, ItemRepository, PersonRepository)

		orderController controller.OrderController = controller.NewOrderController(orderService)
	)
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	routes.OrderRoutes(server, orderController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}