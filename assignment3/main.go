package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// fmt.Println(os.Getenv("TEST"))
}

func main() {
	// fmt.Println(helper.GenerateRandomNumber(1, 10))
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/reyhan", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "begitu syulit lupakan reyhan",
		})
	})

	router.Run(":8080")
}
