package handlers

import "github.com/gin-gonic/gin"


type Controller interface {
	GetAllOrders(c *gin.Context)
	GetOrder(c *gin.Context)
	CreateOrder(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type OrderService struct {
	
}