package main

import (
	"assignment3/pkg/controller"
	"assignment3/pkg/repository"
	"assignment3/pkg/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		panic(err)
	}

}

func Always(function interface{}, duration time.Duration) {
	for {
		<-time.After(duration * time.Second)

		update, ok := function.(func() error)
		if ok {
			go update()
		}
	}
}
func main() {
	var (
		repo       repository.Repository = repository.NewDataRepository(os.Getenv("FILEPATH"))
		service    service.Service       = service.NewService(repo)
		controller controller.Controller = controller.NewController(service)
	)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/reyhan", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "begitu syulit lupakan reyhan",
		})
	})

	os.OpenFile(os.Getenv("FILEPATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	go Always(service.UpdateData, 15)

	v1.GET("/", controller.GetData)

	router.Run(":8080")
}
