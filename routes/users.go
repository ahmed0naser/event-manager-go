package routes

import (
	"net/http"

	"github.com/ahmed0naser/event-manager-go/models"
	"github.com/ahmed0naser/event-manager-go/utils"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed to parse credentials"})
		return
	}
	err=user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't create account"})
		return
	}
	context.JSON(http.StatusCreated,gin.H{"message":"user created successfully"})
}
func Login(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed to parse credentials"})
		return
	}
	err=user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "failed-Invalid Credentials"})
		return
	}
	token,err:=utils.GenerateToken(user.Email,user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "failed to authenticate user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "successful login","token":token})


}