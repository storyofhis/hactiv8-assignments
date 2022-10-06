package controller

import (
	"assignment3/helper"
	"assignment3/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetData(ctx *gin.Context)
}

type controller struct {
	service service.Service
}

func NewController(s service.Service) Controller {
	return &controller{
		service: s,
	}
}

func (c *controller) GetData(ctx *gin.Context) {
	result, err := c.service.GetData()
	if err != nil {
		response := helper.BadRequestError(err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.SuccessFindAllResponse("SUCCESS_GET_DATA", result)
	ctx.JSON(http.StatusOK, response)
}
