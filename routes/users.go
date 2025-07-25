package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhunderdog/go_api/models"
	"github.com/jhunderdog/go_api/utils"
)
func signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func login(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.ValidateCrentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token. Try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token":token})
}
