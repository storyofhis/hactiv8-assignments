package controllers

import (
	"latihan-testing/models/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

func CreateUser(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := entity.User{
		Email:    input.Email,
		Username: input.Username,
		Name:     input.Name,
		Age:      input.Age,
	}
	entity.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
