package routes

import (
	"assignment2/middleware"
	"assignment2/pkg/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, orderController controller.OrderController) {
	orderRoutes := router.Group("/orders", middleware.Auth())
	{
		orderRoutes.GET("", orderController.GetOrders)
		orderRoutes.GET("person/:id", orderController.GetOrderByID)
		orderRoutes.POST("/create", orderController.CreateOrder)
		orderRoutes.PUT("/update", orderController.UpdateOrder)
		orderRoutes.DELETE("/:id", orderController.DeleteOrder)
	}
}