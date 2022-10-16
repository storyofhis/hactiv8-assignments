package controllers

import (
	"latihan-testing/models/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	var user entity.User

	if err := entity.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
