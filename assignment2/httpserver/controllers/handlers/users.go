package handlers

import (
	"fmt"

	"gin-clean-architecture/httpserver/controllers/entity"
	"gin-clean-architecture/httpserver/controllers/params"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Data = []entity.Users{
	{ID: "E001", Name: "John", City: "New York", Age: 12},
	{ID: "E002", Name: "Doe", City: "California", Age: 23},
	{ID: "E003", Name: "Harry", City: "London", Age: 22},
	{ID: "E004", Name: "Maguire", City: "Stockholms", Age: 25},
	{ID: "E005", Name: "Gerry", City: "Tokyo", Age: 16},
	{ID: "E006", Name: "Picky", City: "Singapore", Age: 18},
	{ID: "E007", Name: "Davies", City: "New Zealand", Age: 19},
	{ID: "E008", Name: "Alex", City: "Sydney", Age: 17},
}

func GetUsers(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		Data,
	)
}

func GetUserById(c *gin.Context) {
	var objUser entity.Users
	if err := c.ShouldBindQuery(&objUser); err == nil {
		fmt.Printf("user obj - %+v \n", objUser)
	} else {
		fmt.Printf("error = %+v", err)
	}

	// id := c.Param("id")
	c.JSON(
		http.StatusOK,
		gin.H{
			"data":   &objUser,
			"status": "ok",
		},
	)
}

// func GetUserById(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		var id = r.FormValue("id")
// 		var result []byte
// 		var err error

// 		for _, each := range Data {
// 			if each.ID == id {
// 				result, err = json.Marshal(Data)

// 				if err != nil {
// 					http.Error(w, err.Error(), http.StatusInternalServerError)
// 					return
// 				}

// 				w.Write(result)
// 				return
// 			}

// 		}
// 		http.Error(w, "", http.StatusBadRequest)
// 		return
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// }

func PostUserHandler(c *gin.Context) {
	var userInput params.UserInput

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": userInput.ID,
		// "name": userInput.Name,
	})
}
