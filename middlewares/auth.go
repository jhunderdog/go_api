package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhunderdog/go_api/utils"
)
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == ""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Autorized."})
		return
	}
	context.Set("userId", userId)
	context.Next()
}