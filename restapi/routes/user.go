package routes

import (
	"fmt"
	"net/http"

	"example.com/first-app/models"
	"example.com/first-app/secure"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user has created successfully"})
}
func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}
	err = user.Validateuser()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "please enter valid details"})
		return
	}
	token, err := secure.Generatetoken(user.Email, user.ID)
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "please enter valid details"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successfull", "token": token})

}
