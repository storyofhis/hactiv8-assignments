package httpserver

import (
	"gin-clean-architecture/httpserver/controllers/handlers"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/users", handlers.GetUsers)
	v1.GET("/user", handlers.GetUserById)

	v1.POST("/users", func(ctx *gin.Context) {})

	return router
}
