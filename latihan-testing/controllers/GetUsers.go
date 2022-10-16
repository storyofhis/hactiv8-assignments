package controllers

import (
	"latihan-testing/models/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	var users []entity.User
	entity.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
