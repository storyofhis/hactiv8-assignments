package controllers

import (
	"latihan-testing/common"
	"latihan-testing/models/entity"
	"latihan-testing/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controllers interface {
	FindUsers(c *gin.Context)
	FindUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type controller struct {
	service service.Service
}

func NewControllers(svc service.Service) Controllers {
	return &controller{
		service: svc,
	}
}

func (control *controller) FindUsers(c *gin.Context) {
	result, err := control.service.GetData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
	}
	response := common.BuildResponse(true, "OK", result)

	// var users []entity.User
	// entity.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

type UserInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

func (control *controller) CreateUser(c *gin.Context) {
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

func (control *controller) FindUser(c *gin.Context) {
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
