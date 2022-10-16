package router

import (
	"latihan-testing/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRouter(controller controllers.Controllers) *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "get success",
		})
	})
	v1.POST("/create", controller.CreateUser)
	v1.GET("/get-users", controller.FindUsers)
	v1.GET("/get-user/:id", controller.FindUser)
	return router
}
